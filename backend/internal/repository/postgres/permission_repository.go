package postgres

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) GetAll(ctx context.Context) ([]domain.Permission, error) {
	var perms []Permission
	if err := r.db.WithContext(ctx).Find(&perms).Error; err != nil {
		return nil, err
	}
	domainPerms := make([]domain.Permission, len(perms))
	for i, p := range perms {
		domainPerms[i] = *p.ToDomain()
	}
	return domainPerms, nil
}
