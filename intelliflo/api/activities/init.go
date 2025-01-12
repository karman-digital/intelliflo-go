package activities

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/activities/tasks"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
)

func NewActivityObject(creds credentials.Credentials) *Activity {
	return &Activity{
		ActivityEndpoint: NewActivityService(creds),
		TasksEndpoint:    tasks.NewTaskService(creds),
	}
}

func NewActivityService(creds credentials.Credentials) *ActivityService {
	return &ActivityService{
		creds,
	}
}
