package postgres

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetAll(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role
	// Preload permissions to get full context
	if err := r.db.WithContext(ctx).Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetByID(ctx context.Context, id int) (*domain.Role, error) {
	var role domain.Role
	if err := r.db.WithContext(ctx).Preload("Permissions").First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) Update(ctx context.Context, role *domain.Role) error {
	// Full updates, omitting many2many manual sync for simplicity in Starter Kit unless strictly required
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *roleRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Role{}, id).Error
}

func (r *roleRepository) AssignPermissions(ctx context.Context, roleID int, permissionIDs []int) error {
	var role Role
	if err := r.db.WithContext(ctx).First(&role, roleID).Error; err != nil {
		return err
	}
	var perms []Permission
	if len(permissionIDs) > 0 {
		if err := r.db.WithContext(ctx).Where("id IN ?", permissionIDs).Find(&perms).Error; err != nil {
			return err
		}
	}
	// Replace associated permissions
	return r.db.WithContext(ctx).Model(&role).Association("Permissions").Replace(perms)
}


func (r *roleRepository) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	var role Role
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return role.ToDomain(), nil
}

func (r *roleRepository) GetPermissions(ctx context.Context, roleID int) ([]domain.Permission, error) {
	var role Role
	err := r.db.WithContext(ctx).Preload("Permissions").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	domainPerms := make([]domain.Permission, len(role.Permissions))
	for i, p := range role.Permissions {
		domainPerms[i] = *p.ToDomain()
	}
	return domainPerms, nil
}
