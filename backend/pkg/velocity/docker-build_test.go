package velocity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/velocity-ci/velocity/backend/pkg/velocity"
	"gopkg.in/yaml.v2"
)

func TestDockerBuildUnmarshal(t *testing.T) {
	taskConfigYaml := `
---
description: Builds a docker image
steps:
  - type: build
    description: Docker build
    dockerfile: test.Dockerfile
    context: ./test
    tags:
      - test/a:333
      - test/b:344
`
	var taskConfig velocity.Task

	err := yaml.Unmarshal([]byte(taskConfigYaml), &taskConfig)
	assert.Nil(t, err)

	expectedTaskConfig := velocity.Task{
		Description: "Builds a docker image",
		Steps: []velocity.Step{
			&velocity.DockerBuild{
				BaseStep: velocity.BaseStep{
					Type:          "build",
					Description:   "Docker build",
					OutputStreams: []string{"build"},
					Params:        map[string]velocity.Parameter{},
				},
				Dockerfile: "test.Dockerfile",
				Context:    "./test",
				Tags: []string{
					"test/a:333",
					"test/b:344",
				},
			},
		},
		Docker: velocity.TaskDocker{
			Registries: []velocity.DockerRegistry{},
		},
		Parameters:         []velocity.ParameterConfig{},
		ValidationErrors:   []string{},
		ValidationWarnings: []string{},
	}

	assert.Equal(t, expectedTaskConfig, taskConfig)
}
