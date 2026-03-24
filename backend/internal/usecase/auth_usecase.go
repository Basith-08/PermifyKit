package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/starter/rbac-kit/internal/config"
	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(ctx context.Context, username, password string) (string, error)
	GetUserPermissions(ctx context.Context, userID int) ([]domain.Permission, error)
	GetUserProfile(ctx context.Context, userID int) (*domain.User, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
	cfg      config.Config
}

func NewAuthUseCase(userRepo repository.UserRepository, roleRepo repository.RoleRepository, cfg config.Config) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
		roleRepo: roleRepo,
		cfg:      cfg,
	}
}

func (u *authUseCase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": float64(user.ID),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func (u *authUseCase) GetUserPermissions(ctx context.Context, userID int) ([]domain.Permission, error) {
	roles, err := u.userRepo.GetRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	permMap := make(map[int]domain.Permission)
	for _, role := range roles {
		perms, err := u.roleRepo.GetPermissions(ctx, role.ID)
		if err != nil {
			return nil, err
		}
		for _, p := range perms {
			permMap[p.ID] = p
		}
	}

	var result []domain.Permission
	for _, p := range permMap {
		result = append(result, p)
	}
	return result, nil
}

func (u *authUseCase) GetUserProfile(ctx context.Context, userID int) (*domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
