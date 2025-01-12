package activities

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type Activity struct {
	ActivityEndpoint interfaces.Activity
	TasksEndpoint    interfaces.Task
}

type ActivityService struct {
	credentials.Credentials
}
