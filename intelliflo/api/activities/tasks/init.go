package tasks

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewTaskService(creds credentials.Credentials) *TaskService {
	return &TaskService{
		creds,
	}
}
