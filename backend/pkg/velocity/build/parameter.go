package build

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/velocity-ci/velocity/backend/pkg/exec"
	v1 "github.com/velocity-ci/velocity/backend/pkg/velocity/genproto/v1"
)

func getSecrets(params map[string]*Parameter) (r []string) {
	for _, p := range params {
		if p.IsSecret {
			r = append(r, p.Value)
		}
	}

	return r
}

func resolveConfigParameter(
	p v1.Parameter,
	bR BackupResolver,
	projectRoot string,
	writer io.Writer,
) (parameters []*Parameter, err error) {
	// resolve parameter value at build time
	switch x := p.GetImpl().(type) {
	case *v1.Parameter_BasicParameter:
		writer.Write([]byte(fmt.Sprintf("-> resolving parameter %s\n", x.BasicParameter.GetName())))
		return resolveConfigParameterBasic(&p, bR)
	case *v1.Parameter_DerivedParameter:
		writer.Write([]byte(fmt.Sprintf("-> resolving parameter %s\n", x.DerivedParameter.GetUse())))
		return resolveConfigParameterDerived(&p, bR, projectRoot, writer)
	default:
		return parameters, fmt.Errorf("type: %T: %v", x, p)
	}
}

func resolveConfigParameterBasic(p *v1.Parameter, backupResolver BackupResolver) (parameters []*Parameter, err error) {
	val, err := backupResolver.Resolve(p.GetBasicParameter().GetName())
	if err != nil {
		return nil, err
	}
	v := val
	return []*Parameter{{
		Name:     p.GetBasicParameter().GetName(),
		Value:    v,
		IsSecret: p.GetSecret(),
	}}, err
}

func resolveConfigParameterDerived(
	p *v1.Parameter,
	backupResolver BackupResolver,
	projectRoot string,
	writer io.Writer,
) (parameters []*Parameter, err error) {
	// Download binary from use:
	bin, err := getBinary(projectRoot, p.GetDerivedParameter().GetUse(), writer)
	if err != nil {
		return parameters, err
	}
	cmd := []string{bin}

	// Process arguments
	for k, v := range p.GetDerivedParameter().GetArguments() {
		cmd = append(cmd, fmt.Sprintf("-%s=%s", k, v))
	}

	// Run binary
	s := exec.Run(cmd, "", os.Environ(), BlankWriter{})
	if s.Error != nil {
		return parameters, s.Error
	}
	var dOutput derivedOutput
	json.Unmarshal([]byte(s.Stdout[0]), &dOutput)

	if dOutput.State == "warning" {
		for paramName := range dOutput.Exports {
			val, err := backupResolver.Resolve(paramName)
			if err != nil {
				return parameters, err
			}
			parameters = append(parameters, &Parameter{
				Name:     paramName,
				Value:    val,
				IsSecret: dOutput.Secret,
			})
		}
	} else if dOutput.State == "success" {
		for paramName, val := range dOutput.Exports {
			parameters = append(parameters, &Parameter{
				Name:     getExportedParameterName(p.GetDerivedParameter().GetExports(), paramName),
				Value:    val,
				IsSecret: dOutput.Secret,
			})
		}
	} else {
		return parameters, fmt.Errorf("binary %s: %s", dOutput.State, dOutput.Error)
	}

	return parameters, nil
}

type Parameter struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	IsSecret bool   `json:"isSecret"`
}

type BackupResolver interface {
	Resolve(paramName string) (string, error)
}

func getExportedParameterName(pMapping map[string]string, exportedParam string) string {
	if val, ok := pMapping[exportedParam]; ok {
		return val
	}

	return exportedParam
}

type derivedOutput struct {
	Secret  bool              `json:"secret"`
	Exports map[string]string `json:"exports"`
	Expires time.Time         `json:"expires"`
	Error   string            `json:"error"`
	State   string            `json:"state"`
}
