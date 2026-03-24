package repository

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
	GetRoles(ctx context.Context, userID int) ([]domain.Role, error)
	AssignRoles(ctx context.Context, userID int, roleIDs []int) error
}
