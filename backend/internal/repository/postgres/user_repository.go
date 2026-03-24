package postgres

import (
	"context"
	"errors"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetByID(ctx context.Context, id int) (*domain.User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found, handled gracefully
		}
		return nil, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("username = ?", username).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) Create(ctx context.Context, du *domain.User) error {
	userModel := FromDomainUser(du)
	err := r.db.WithContext(ctx).Create(userModel).Error
	if err != nil {
		return err
	}
	du.ID = userModel.ID
	return nil
}

func (r *userRepo) GetRoles(ctx context.Context, userID int) ([]domain.Role, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", userID).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return nil, err
	}
	domainRoles := make([]domain.Role, len(user.Roles))
	for i, role := range user.Roles {
		domainRoles[i] = *role.ToDomain()
	}
	return domainRoles, nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []User
	if err := r.db.WithContext(ctx).Preload("Roles").Preload("Roles.Permissions").Find(&users).Error; err != nil {
		return nil, err
	}
	domainUsers := make([]domain.User, len(users))
	for i, u := range users {
		domainUsers[i] = *u.ToDomain()
	}
	return domainUsers, nil
}

func (r *userRepo) Update(ctx context.Context, du *domain.User) error {
	userModel := FromDomainUser(du)
	// Clear associations to prevent unintended overwrites, just update base fields
	return r.db.WithContext(ctx).Model(&User{}).Where("id = ?", userModel.ID).Updates(userModel).Error
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&User{}, id).Error
}

func (r *userRepo) AssignRoles(ctx context.Context, userID int, roleIDs []int) error {
	var user User
	if err := r.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return err
	}
	var roles []Role
	if len(roleIDs) > 0 {
		if err := r.db.WithContext(ctx).Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
			return err
		}
	}
	// Replace associated roles
	return r.db.WithContext(ctx).Model(&user).Association("Roles").Replace(roles)
}

