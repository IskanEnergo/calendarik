package users_service

import (
	"context"
	"fmt"

	"github.com/IskanEnergo/calendarik/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	//ctx

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("user validation: %w", err)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil

	//1.  user.Validate()
	//2. repo.Save(user)
	//3. return user
}
