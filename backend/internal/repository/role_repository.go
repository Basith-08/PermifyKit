package repository

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
)

type RoleRepository interface {
	GetAll(ctx context.Context) ([]domain.Role, error)
	GetByID(ctx context.Context, id int) (*domain.Role, error)
	GetByName(ctx context.Context, name string) (*domain.Role, error)
	GetPermissions(ctx context.Context, roleID int) ([]domain.Permission, error)
	Create(ctx context.Context, role *domain.Role) error
	Update(ctx context.Context, role *domain.Role) error
	Delete(ctx context.Context, id int) error
	AssignPermissions(ctx context.Context, roleID int, permissionIDs []int) error
}
