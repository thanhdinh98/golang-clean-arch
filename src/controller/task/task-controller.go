package task

import (
	"errors"
	"togo/src"
	"togo/src/domain/task"
	"togo/src/infra/serror"
	"togo/src/schema"
)

type TaskController struct {
	taskWokflow  task.ITaskWorkflow
	errorFactory src.IErrorFactory
}

func (this *TaskController) AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error) {
	yes, err := context.HasPermission([]string{src.CREATE_TASK})
	if err != nil {
		return nil, this.errorFactory.UnauthorizedError(src.TOKEN_INVALID, err)
	}

	if !yes {
		return nil, this.errorFactory.UnauthorizedError(src.NO_PERMISSION, errors.New(src.CREATE_TASK))
	}

	return this.taskWokflow.AddTaskByOwner(context, data)
}

func NewTaskController() ITaskController {
	return &TaskController{
		task.NewTaskWorkflow(),
		serror.NewErrorFactory(),
	}
}
