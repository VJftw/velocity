package rest

import (
	"net/http"

	"github.com/velocity-ci/velocity/backend/pkg/domain/githistory"
	"github.com/velocity-ci/velocity/backend/velocity"

	"github.com/velocity-ci/velocity/backend/pkg/domain/project"

	"github.com/labstack/echo"
	"github.com/velocity-ci/velocity/backend/pkg/domain/task"
)

type taskResponse struct {
	UUID string `json:"id"`
	Slug string `json:"slug"`
	*velocity.Task
}

type taskList struct {
	Total int             `json:"total"`
	Data  []*taskResponse `json:"data"`
}

func newTaskResponse(t *task.Task) *taskResponse {
	return &taskResponse{
		UUID: t.UUID,
		Slug: t.Slug,
		Task: t.Task,
	}
}

type taskHandler struct {
	projectManager *project.Manager
	commitManager  *githistory.CommitManager
	taskManager    *task.Manager
}

func newTaskHandler(
	projectManager *project.Manager,
	commitManager *githistory.CommitManager,
	taskManager *task.Manager,
) *taskHandler {
	return &taskHandler{
		projectManager: projectManager,
		commitManager:  commitManager,
		taskManager:    taskManager,
	}
}

func (h *taskHandler) getAllForCommit(c echo.Context) error {

	commit := getCommitByProjectAndHash(c, h.projectManager, h.commitManager)
	if commit == nil {
		return nil
	}

	pQ := getPagingQueryParams(c)
	if pQ == nil {
		return nil
	}

	cs, total := h.taskManager.GetAllForCommit(commit, pQ)
	rTasks := []*taskResponse{}
	for _, c := range cs {
		rTasks = append(rTasks, newTaskResponse(c))
	}

	c.JSON(http.StatusOK, taskList{
		Total: total,
		Data:  rTasks,
	})

	return nil
}

func (h *taskHandler) getByProjectCommitAndSlug(c echo.Context) error {

	if task := getTaskByCommitAndSlug(c, h.projectManager, h.commitManager, h.taskManager); task != nil {
		c.JSON(http.StatusOK, newTaskResponse(task))
	}

	return nil
}

func getTaskByCommitAndSlug(
	c echo.Context,
	pM *project.Manager,
	cM *githistory.CommitManager,
	tM *task.Manager,
) *task.Task {
	commit := getCommitByProjectAndHash(c, pM, cM)
	if c == nil {
		return nil
	}

	slug := c.Param("taskSlug")
	t, err := tM.GetByCommitAndSlug(commit, slug)
	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return nil
	}

	return t
}

func (h *taskHandler) sync(c echo.Context) error {
	p := getProjectBySlug(c, h.projectManager)
	if p == nil {
		return nil
	}

	p, _ = h.taskManager.Sync(p)

	c.JSON(http.StatusOK, newProjectResponse(p))
	return nil
}
