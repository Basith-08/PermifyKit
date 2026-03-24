package usecase

import (
	"context"
	"errors"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
)

type RoleUseCase interface {
	GetAllRoles(ctx context.Context) ([]domain.Role, error)
	GetRoleByID(ctx context.Context, id int) (*domain.Role, error)
	CreateRole(ctx context.Context, name, description string, permissionIDs []int) (*domain.Role, error)
	UpdateRole(ctx context.Context, id int, name, description string, permissionIDs []int) (*domain.Role, error)
	DeleteRole(ctx context.Context, id int) error
}

type roleUseCase struct {
	roleRepo repository.RoleRepository
	auditUC  AuditLogUseCase
}

func NewRoleUseCase(roleRepo repository.RoleRepository, auditUC AuditLogUseCase) RoleUseCase {
	return &roleUseCase{roleRepo: roleRepo, auditUC: auditUC}
}

func (r *roleUseCase) GetAllRoles(ctx context.Context) ([]domain.Role, error) {
	return r.roleRepo.GetAll(ctx)
}

func (r *roleUseCase) GetRoleByID(ctx context.Context, id int) (*domain.Role, error) {
	return r.roleRepo.GetByID(ctx, id)
}

func (r *roleUseCase) CreateRole(ctx context.Context, name, description string, permissionIDs []int) (*domain.Role, error) {
	existingRole, _ := r.roleRepo.GetByName(ctx, name)
	if existingRole != nil {
		return nil, errors.New("role name already exists")
	}

	role := &domain.Role{
		Name:        name,
		Description: description,
	}

	if err := r.roleRepo.Create(ctx, role); err != nil {
		return nil, err
	}

	if len(permissionIDs) > 0 {
		if err := r.roleRepo.AssignPermissions(ctx, role.ID, permissionIDs); err != nil {
			return nil, err
		}
	}

	r.auditUC.LogAction(ctx, 0, "system", "CREATE", "roles", name, "")

	return r.roleRepo.GetByID(ctx, role.ID)
}

func (r *roleUseCase) UpdateRole(ctx context.Context, id int, name, description string, permissionIDs []int) (*domain.Role, error) {
	role, err := r.roleRepo.GetByID(ctx, id)
	if err != nil || role == nil {
		return nil, errors.New("role not found")
	}

	role.Name = name
	role.Description = description

	if err := r.roleRepo.Update(ctx, role); err != nil {
		return nil, err
	}

	if permissionIDs != nil {
		if err := r.roleRepo.AssignPermissions(ctx, role.ID, permissionIDs); err != nil {
			return nil, err
		}
	}

	r.auditUC.LogAction(ctx, 0, "system", "UPDATE", "roles", name, "")

	return r.roleRepo.GetByID(ctx, role.ID)
}

func (r *roleUseCase) DeleteRole(ctx context.Context, id int) error {
	r.auditUC.LogAction(ctx, 0, "system", "DELETE", "roles", string(rune(id)), "")
	return r.roleRepo.Delete(ctx, id)
}
