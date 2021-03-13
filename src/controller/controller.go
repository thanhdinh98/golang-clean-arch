package controller

import (
	"togo/src/controller/task"
	"togo/src/controller/user"
)

var (
	NewUserController func() user.IUserController = user.NewUserController
	NewTaskController func() task.ITaskController = task.NewTaskController
)
