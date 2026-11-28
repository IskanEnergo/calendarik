package tasks_postgres_repository

import (
	"time"

	"github.com/IskanEnergo/calendarik/internal/core/domain"
)

type TaskModel struct {
	ID      int
	Version int

	Title        string
	Description  *string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
	AuthorUserID int
}

func taskDomainsFromModels(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))

	for i, m := range models {
		domains[i] = taskDomainFromModel(m)
	}
	return domains
}

func taskDomainFromModel(model TaskModel) domain.Task {

	taskDomain := domain.NewTask(
		model.ID,
		model.Version,
		model.Title,
		model.Description,
		model.Completed,
		model.Created_at,
		model.Completed_at,
		model.AuthorUserID,
	)
	return taskDomain
}
