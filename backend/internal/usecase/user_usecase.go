package usecase

import (
	"context"
	"errors"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, username, email, password string, roleIDs []int) (*domain.User, error)
	UpdateUser(ctx context.Context, id int, username, email, password string, roleIDs []int) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
	auditUC  AuditLogUseCase
}

func NewUserUseCase(userRepo repository.UserRepository, auditUC AuditLogUseCase) UserUseCase {
	return &userUseCase{userRepo: userRepo, auditUC: auditUC}
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUseCase) CreateUser(ctx context.Context, username, email, password string, roleIDs []int) (*domain.User, error) {
	existingUser, _ := u.userRepo.GetByUsername(ctx, username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	if len(roleIDs) > 0 {
		if err := u.userRepo.AssignRoles(ctx, user.ID, roleIDs); err != nil {
			return nil, err
		}
	}

	u.auditUC.LogAction(ctx, 0, "system", "CREATE", "users", username, "")

	return u.userRepo.GetByID(ctx, user.ID)
}

func (u *userUseCase) UpdateUser(ctx context.Context, id int, username, email, password string, roleIDs []int) (*domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	user.Username = username
	user.Email = email

	if password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hash)
	}

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	// Overwrite roles if a new set is provided.
	if roleIDs != nil {
		if err := u.userRepo.AssignRoles(ctx, user.ID, roleIDs); err != nil {
			return nil, err
		}
	}

	u.auditUC.LogAction(ctx, 0, "system", "UPDATE", "users", username, "")

	return u.userRepo.GetByID(ctx, user.ID)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int) error {
	u.auditUC.LogAction(ctx, 0, "system", "DELETE", "users", string(rune(id)), "")
	return u.userRepo.Delete(ctx, id)
}
