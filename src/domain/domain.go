package domain

import (
	"togo/src/domain/task"
	"togo/src/domain/user"
)

var (
	NewUserWorkflow func() user.IUserWorkflow = user.NewUserWorkflow
	NewTaskWorkflow func() task.ITaskWorkflow = task.NewTaskWorkflow
)
