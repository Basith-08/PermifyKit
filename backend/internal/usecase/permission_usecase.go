package usecase

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
)

type PermissionUseCase interface {
	GetAllPermissions(ctx context.Context) ([]domain.Permission, error)
}

type permissionUseCase struct {
	permRepo repository.PermissionRepository
}

func NewPermissionUseCase(permRepo repository.PermissionRepository) PermissionUseCase {
	return &permissionUseCase{permRepo: permRepo}
}

func (p *permissionUseCase) GetAllPermissions(ctx context.Context) ([]domain.Permission, error) {
	return p.permRepo.GetAll(ctx)
}
