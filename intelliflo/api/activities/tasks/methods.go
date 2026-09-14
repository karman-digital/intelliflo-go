package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	taskmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/tasks"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (s *TaskService) GetTask(taskId int, opts ...sharedmodels.GetOptions) (taskmodels.Task, error) {
	var task taskmodels.Task
	resp, err := s.SendRequest("GET", fmt.Sprintf("activities/tasks/%d", taskId), nil, opts...)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return task, fmt.Errorf("get task returned status %d: %w", resp.StatusCode, err)
	}
	err = json.Unmarshal(respBody, &task)
	if err != nil {
		return task, fmt.Errorf("error parsing body: %v", err)
	}
	return task, nil
}

func (s *TaskService) GetTasksByReference(reference string) (taskmodels.TasksResponse, error) {
	if strings.TrimSpace(reference) == "" || strings.Contains(reference, "'") {
		return taskmodels.TasksResponse{}, fmt.Errorf("invalid task reference")
	}
	all, err := s.GetAllTasks()
	if err != nil {
		return taskmodels.TasksResponse{}, err
	}
	matches := make([]taskmodels.Task, 0, 1)
	for _, task := range all.Items {
		if strings.HasSuffix(strings.TrimSpace(task.Description), "["+reference+"]") {
			matches = append(matches, task)
		}
	}
	all.Items = matches
	all.Count = len(matches)
	all.NextHref = ""
	return all, nil
}

func (s *TaskService) GetAllTasks() (taskmodels.TasksResponse, error) {
	const pageSize = 500
	first, err := s.GetTasks(sharedmodels.GetOptions{Top: pageSize})
	if err != nil {
		return taskmodels.TasksResponse{}, err
	}
	pageCount := (first.Count + pageSize - 1) / pageSize
	if pageCount <= 1 {
		first.NextHref = ""
		return first, nil
	}

	pages := make([]taskmodels.TasksResponse, pageCount-1)
	jobs := make(chan int)
	errCh := make(chan error, pageCount-1)
	var workers sync.WaitGroup
	workerCount := 6
	if workerCount > len(pages) {
		workerCount = len(pages)
	}
	for worker := 0; worker < workerCount; worker++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for pageIndex := range jobs {
				page, pageErr := s.GetTasks(sharedmodels.GetOptions{Top: pageSize, Skip: (pageIndex + 1) * pageSize})
				if pageErr != nil {
					errCh <- pageErr
					continue
				}
				pages[pageIndex] = page
			}
		}()
	}
	for pageIndex := range pages {
		jobs <- pageIndex
	}
	close(jobs)
	workers.Wait()
	close(errCh)
	if err := <-errCh; err != nil {
		return taskmodels.TasksResponse{}, err
	}
	for _, page := range pages {
		first.Items = append(first.Items, page.Items...)
	}
	first.NextHref = ""
	return first, nil
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

func (s *TaskService) CreateTask(task taskmodels.TaskCreateRequest) (taskmodels.Task, error) {
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
		return newTask, fmt.Errorf("create task returned status %d: %w", resp.StatusCode, err)
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
	if _, err := shared.HandleCustomResponseCode(resp, http.StatusNoContent); err != nil {
		return fmt.Errorf("delete task returned status %d: %w", resp.StatusCode, err)
	}
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

func (s *TaskService) CreateTaskNote(taskId int, note taskmodels.TaskNoteCreateRequest) (taskmodels.TaskNote, error) {
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
		return newNote, fmt.Errorf("create task note returned status %d: %w", resp.StatusCode, err)
	}
	err = json.Unmarshal(respBody, &newNote)
	if err != nil {
		return newNote, fmt.Errorf("error parsing body: %v", err)
	}
	return newNote, nil
}
