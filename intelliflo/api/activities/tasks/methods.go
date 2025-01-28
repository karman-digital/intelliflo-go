package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	taskmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/tasks"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (s *TaskService) GetTask(taskId int, opts ...sharedmodels.GetOptions) (taskmodels.Task, error) {
	var task taskmodels.Task
	resp, err := s.SendRequest("GET", fmt.Sprintf("/v2/activities/tasks/%d", taskId), nil, opts...)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return task, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &task)
	if err != nil {
		return task, fmt.Errorf("error parsing body: %v", err)
	}
	return task, nil
}

func (s *TaskService) GetTasks(opts ...sharedmodels.GetOptions) (taskmodels.TasksResponse, error) {
	var tasks taskmodels.TasksResponse
	resp, err := s.SendRequest("GET", "activities/tasks", nil, opts...)
	if err != nil {
		return tasks, err
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return tasks, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &tasks)
	if err != nil {
		return tasks, fmt.Errorf("error parsing body: %v", err)
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(task taskmodels.Task) (taskmodels.Task, error) {
	var newTask taskmodels.Task
	reqBody, err := json.Marshal(task)
	if err != nil {
		return newTask, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("POST", "activities/tasks", reqBody)
	if err != nil {
		return newTask, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newTask, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &newTask)
	if err != nil {
		return newTask, fmt.Errorf("error parsing body: %v", err)
	}
	return newTask, nil
}

func (s *TaskService) UpdateTask(taskId int, task taskmodels.Task) (taskmodels.Task, error) {
	var updatedTask taskmodels.Task
	reqBody, err := json.Marshal(task)
	if err != nil {
		return updatedTask, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("PUT", fmt.Sprintf("activities/tasks/%d", taskId), reqBody)
	if err != nil {
		return updatedTask, fmt.Errorf("error making put request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return updatedTask, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &updatedTask)
	if err != nil {
		return updatedTask, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(taskId int) error {
	resp, err := s.SendRequest("DELETE", fmt.Sprintf("activities/tasks/%d", taskId), nil)
	if err != nil {
		return fmt.Errorf("error making delete request: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

func (s *TaskService) GetTaskNotes(taskId int, opts ...sharedmodels.GetOptions) (taskmodels.TaskNotesResponse, error) {
	var notes taskmodels.TaskNotesResponse
	resp, err := s.SendRequest("GET", fmt.Sprintf("activities/tasks/%d/notes", taskId), nil, opts...)
	if err != nil {
		return notes, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return notes, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &notes)
	if err != nil {
		return notes, fmt.Errorf("error parsing body: %v", err)
	}
	return notes, nil
}

func (s *TaskService) CreateTaskNote(taskId int, note taskmodels.TaskNote) (taskmodels.TaskNote, error) {
	var newNote taskmodels.TaskNote
	reqBody, err := json.Marshal(note)
	if err != nil {
		return newNote, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("POST", fmt.Sprintf("activities/tasks/%d/notes", taskId), reqBody)
	if err != nil {
		return newNote, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newNote, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &newNote)
	if err != nil {
		return newNote, fmt.Errorf("error parsing body: %v", err)
	}
	return newNote, nil
}
