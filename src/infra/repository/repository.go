package repository

import (
	"togo/src/entity/task"
	"togo/src/entity/user"
	taskRepo "togo/src/infra/repository/task"
	userRepo "togo/src/infra/repository/user"
)

var (
	NewUserRepository func() user.IUserRepository = userRepo.NewUserRepository
	NewTaskRepository func() task.ITaskRepository = taskRepo.NewTaskRepository
)
