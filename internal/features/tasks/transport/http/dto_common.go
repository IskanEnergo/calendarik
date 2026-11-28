package tasks_transport_http

import (
	"time"

	"github.com/IskanEnergo/calendarik/internal/core/domain"
)

type TaskDTOResponse struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	Created_at   time.Time  `json:"created_at"`
	Completed_at *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func taskDTOFromDOmain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		Created_at:   task.CreatedAt,
		Completed_at: task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOFromDOmains(tasks []domain.Task) []TaskDTOResponse {

	dtos := make([]TaskDTOResponse, len(tasks))
	for i, t := range tasks {
		dtos[i] = taskDTOFromDOmain(t)
	}
	return dtos
}
