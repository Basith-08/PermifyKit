package postgres

import (
	"time"

	"github.com/starter/rbac-kit/internal/domain"
)

type User struct {
	ID           int       `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []Role    `gorm:"many2many:user_roles;"`
}

func (u *User) ToDomain() *domain.User {
	if u == nil {
		return nil
	}
	domainRoles := make([]domain.Role, len(u.Roles))
	for i, r := range u.Roles {
		domainRoles[i] = *r.ToDomain()
	}
	return &domain.User{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Roles:        domainRoles,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func FromDomainUser(du *domain.User) *User {
	if du == nil {
		return nil
	}
	return &User{
		ID:           du.ID,
		Username:     du.Username,
		Email:        du.Email,
		PasswordHash: du.PasswordHash,
		CreatedAt:    du.CreatedAt,
		UpdatedAt:    du.UpdatedAt,
	}
}

type Role struct {
	ID          int          `gorm:"primaryKey"`
	Name        string       `gorm:"unique;not null"`
	Description string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (r *Role) ToDomain() *domain.Role {
	if r == nil {
		return nil
	}
	domainPerms := make([]domain.Permission, len(r.Permissions))
	for i, p := range r.Permissions {
		domainPerms[i] = *p.ToDomain()
	}
	return &domain.Role{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Permissions: domainPerms,
	}
}

type Permission struct {
	ID          int    `gorm:"primaryKey"`
	Resource    string `gorm:"not null"`
	Action      string `gorm:"not null"`
	Description string
}

func (p *Permission) ToDomain() *domain.Permission {
	if p == nil {
		return nil
	}
	return &domain.Permission{
		ID:          p.ID,
		Resource:    p.Resource,
		Action:      p.Action,
		Description: p.Description,
	}
}

type AuditLog struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	Username   string    `gorm:"not null"`
	Action     string    `gorm:"not null"`
	Resource   string    `gorm:"not null"`
	ResourceID string    `gorm:"not null"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
	IPAddress  string
}

func (a *AuditLog) ToDomain() *domain.AuditLog {
	if a == nil {
		return nil
	}
	return &domain.AuditLog{
		ID:         a.ID,
		UserID:     a.UserID,
		Username:   a.Username,
		Action:     a.Action,
		Resource:   a.Resource,
		ResourceID: a.ResourceID,
		Timestamp:  a.Timestamp,
		IPAddress:  a.IPAddress,
	}
}
