package repository

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
)

type PermissionRepository interface {
	GetAll(ctx context.Context) ([]domain.Permission, error)
}
