package step

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/velocity-ci/velocity/backend/pkg/velocity"
	"github.com/velocity-ci/velocity/backend/pkg/velocity/io"
)

type DockerPush struct {
	BaseStep
	Tags []string `json:"tags"`
}

func NewDockerPush() *DockerPush {
	return &DockerPush{
		Tags:     []string{},
		BaseStep: newBaseStep("push", []string{"push"}),
	}
}

func (dP DockerPush) GetDetails() string {
	return fmt.Sprintf("tags: %s", dP.Tags)
}

func (dP *DockerPush) Execute(emitter velocity.Emitter, tsk *velocity.Task) error {
	writer := emitter.GetStreamWriter("push")
	defer writer.Close()
	writer.SetStatus(StateRunning)
	fmt.Fprintf(writer, io.ColorFmt(io.ANSIInfo, "-> %s"), dP.Description)

	cli, _ := client.NewEnvClient()
	ctx := context.Background()

	for _, t := range dP.Tags {
		imageIDProgress = map[string]string{}
		// Determine correct authToken
		authToken := getAuthToken(t, tsk.Docker.Registries)
		reader, err := cli.ImagePush(ctx, t, types.ImagePushOptions{
			All:          true,
			RegistryAuth: authToken,
		})
		if err != nil {
			velocity.GetLogger().Error("could not push docker image", zap.String("image", t), zap.Error(err))
			writer.SetStatus(StateFailed)
			fmt.Fprintf(writer, io.ColorFmt(io.ANSIError, "-> push failed: %s"), err)
			return err
		}
		handleOutput(reader, tsk.ResolvedParameters, writer)
		fmt.Fprintf(writer, io.ColorFmt(io.ANSIInfo, "-> pushed: %s"), t)

	}

	writer.SetStatus(StateSuccess)
	fmt.Fprintf(writer, io.ColorFmt(io.ANSISuccess, "-> success"))
	return nil

}

func (dP DockerPush) Validate(params map[string]Parameter) error {
	return nil
}

func (dP *DockerPush) SetParams(params map[string]Parameter) error {
	for paramName, param := range params {
		tags := []string{}
		for _, c := range dP.Tags {
			correctedTag := strings.Replace(c, fmt.Sprintf("${%s}", paramName), param.Value, -1)
			tags = append(tags, correctedTag)
		}
		dP.Tags = tags
	}
	return nil
}
