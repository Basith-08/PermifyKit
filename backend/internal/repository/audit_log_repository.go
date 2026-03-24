package repository

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
)

type AuditLogRepository interface {
	Create(ctx context.Context, log *domain.AuditLog) error
	GetAll(ctx context.Context, limit int) ([]domain.AuditLog, error)
}
