package velocity

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerRun struct {
	BaseStep       `yaml:",inline"`
	Image          string            `json:"image" yaml:"image"`
	Command        []string          `json:"command" yaml:"command"`
	Environment    map[string]string `json:"environment" yaml:"environment"`
	WorkingDir     string            `json:"workingDir" yaml:"working_dir"`
	MountPoint     string            `json:"mountPoint" yaml:"mount_point"`
	IgnoreExitCode bool              `json:"ignoreExitCode" yaml:"ignore_exit"`
}

func NewDockerRun() *DockerRun {
	return &DockerRun{
		Image:          "",
		Command:        []string{},
		Environment:    map[string]string{},
		WorkingDir:     "",
		MountPoint:     "",
		IgnoreExitCode: false,
		BaseStep: BaseStep{
			Type:          "run",
			OutputStreams: []string{"run"},
		},
	}
}

func (dR DockerRun) GetDetails() string {
	return fmt.Sprintf("image: %s command: %s", dR.Image, dR.Command)
}

func (dR *DockerRun) Execute(emitter Emitter, params map[string]Parameter) error {
	writer := emitter.NewStreamWriter("run")
	writer.SetStatus(StateRunning)
	writer.Write([]byte(fmt.Sprintf("%s\n## %s\n\x1b[0m", infoANSI, dR.Description)))

	if dR.MountPoint == "" {
		dR.MountPoint = "/velocity_ci"
	}
	env := []string{}
	for k, v := range dR.Environment {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	cwd, _ := os.Getwd()
	if os.Getenv("SIB_CWD") != "" {
		cwd = os.Getenv("SIB_CWD")
	}

	config := &container.Config{
		Image: dR.Image,
		Cmd:   dR.Command,
		Volumes: map[string]struct{}{
			dR.MountPoint: struct{}{},
		},
		WorkingDir: fmt.Sprintf("%s/%s", dR.MountPoint, dR.WorkingDir),
		Env:        env,
	}

	hostConfig := &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:%s", cwd, dR.MountPoint),
		},
	}

	var wg sync.WaitGroup
	cli, _ := client.NewEnvClient()
	ctx := context.Background()

	networkResp, err := cli.NetworkCreate(ctx, fmt.Sprintf("vci-%s", dR.GetRunID()), types.NetworkCreate{
		Labels: map[string]string{"owner": "velocity-ci"},
	})
	if err != nil {
		log.Println(err)
	}

	sR := newServiceRunner(
		cli,
		ctx,
		writer,
		&wg,
		params,
		fmt.Sprintf("%s-%s", dR.GetRunID(), "run"),
		dR.Image,
		nil,
		config,
		hostConfig,
		nil,
		networkResp.ID,
	)

	sR.PullOrBuild()
	sR.Create()
	stopServicesChannel := make(chan string, 32)
	wg.Add(1)
	go sR.Run(stopServicesChannel)
	_ = <-stopServicesChannel
	sR.Stop()
	wg.Wait()
	err = cli.NetworkRemove(ctx, networkResp.ID)
	if err != nil {
		log.Printf("network %s remove err: %s", networkResp.ID, err)
	}

	exitCode := sR.exitCode

	if err != nil {
		return err
	}

	if exitCode != 0 && !dR.IgnoreExitCode {
		writer.SetStatus(StateFailed)
		writer.Write([]byte(fmt.Sprintf("%s\n### FAILED (exited: %d)\x1b[0m", errorANSI, exitCode)))
		return fmt.Errorf("Non-zero exit code: %d", exitCode)
	}

	writer.SetStatus(StateSuccess)
	writer.Write([]byte(fmt.Sprintf("%s\n### SUCCESS (exited: %d)\x1b[0m", successANSI, exitCode)))
	return nil

}

func (dR DockerRun) Validate(params map[string]Parameter) error {
	re := regexp.MustCompile("\\$\\{(.+)\\}")

	requiredParams := re.FindAllStringSubmatch(dR.Image, -1)
	if !isAllInParams(requiredParams, params) {
		return fmt.Errorf("Parameter %v missing", requiredParams)
	}
	requiredParams = re.FindAllStringSubmatch(dR.WorkingDir, -1)
	if !isAllInParams(requiredParams, params) {
		return fmt.Errorf("Parameter %v missing", requiredParams)
	}
	for _, c := range dR.Command {
		requiredParams = re.FindAllStringSubmatch(c, -1)
		if !isAllInParams(requiredParams, params) {
			return fmt.Errorf("Parameter %v missing", requiredParams)
		}
	}

	for key, val := range dR.Environment {
		requiredParams = re.FindAllStringSubmatch(key, -1)
		if !isAllInParams(requiredParams, params) {
			return fmt.Errorf("Parameter %v missing", requiredParams)
		}
		requiredParams = re.FindAllStringSubmatch(val, -1)
		if !isAllInParams(requiredParams, params) {
			return fmt.Errorf("Parameter %v missing", requiredParams)
		}
	}
	return nil
}

func (dR *DockerRun) SetParams(params map[string]Parameter) error {
	for paramName, param := range params {
		dR.Image = strings.Replace(dR.Image, fmt.Sprintf("${%s}", paramName), param.Value, -1)
		dR.WorkingDir = strings.Replace(dR.WorkingDir, fmt.Sprintf("${%s}", paramName), param.Value, -1)

		cmd := []string{}
		for _, c := range dR.Command {
			correctedCmd := strings.Replace(c, fmt.Sprintf("${%s}", paramName), param.Value, -1)
			cmd = append(cmd, correctedCmd)
		}
		dR.Command = cmd

		env := map[string]string{}
		for key, val := range dR.Environment {
			correctedKey := strings.Replace(key, fmt.Sprintf("${%s}", paramName), param.Value, -1)
			correctedVal := strings.Replace(val, fmt.Sprintf("${%s}", paramName), param.Value, -1)
			env[correctedKey] = correctedVal
		}

		dR.Environment = env
	}
	return nil
}

func (dR *DockerRun) String() string {
	j, _ := json.Marshal(dR)
	return string(j)
}

func isAllInParams(matches [][]string, params map[string]Parameter) bool {
	for _, match := range matches {
		found := false
		for paramName, _ := range params {
			if paramName == match[1] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
