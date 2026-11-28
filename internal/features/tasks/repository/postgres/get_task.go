package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/IskanEnergo/calendarik/internal/core/domain"
	core_errors "github.com/IskanEnergo/calendarik/internal/core/errors"
	core_postgres_pool "github.com/IskanEnergo/calendarik/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(ctx context.Context, id int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM calendarik.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.Created_at,
		&taskModel.Completed_at,
		&taskModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with id ='%d': %w", id, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	taskDomain := domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.Created_at,
		taskModel.Completed_at,
		taskModel.AuthorUserID,
	)
	return taskDomain, nil
}
