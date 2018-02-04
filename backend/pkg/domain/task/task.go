package task

import (
	"github.com/velocity-ci/velocity/backend/pkg/domain/githistory"
	"github.com/velocity-ci/velocity/backend/velocity"
)

type Task struct {
	ID     string             `json:"id"`
	Commit *githistory.Commit `json:"commit"`
	Slug   string             `json:"slug"`
	*velocity.Task
}
