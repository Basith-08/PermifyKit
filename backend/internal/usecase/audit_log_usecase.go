package usecase

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
)

type AuditLogUseCase interface {
	LogAction(ctx context.Context, userID int, username, action, resource, resourceID, ip string) error
	GetLatestLogs(ctx context.Context, limit int) ([]domain.AuditLog, error)
}

type auditLogUseCase struct {
	auditRepo repository.AuditLogRepository
}

func NewAuditLogUseCase(auditRepo repository.AuditLogRepository) AuditLogUseCase {
	return &auditLogUseCase{auditRepo: auditRepo}
}

func (u *auditLogUseCase) LogAction(ctx context.Context, userID int, username, action, resource, resourceID, ip string) error {
	log := &domain.AuditLog{
		UserID:     userID,
		Username:   username,
		Action:     action,
		Resource:   resource,
		ResourceID: resourceID,
		IPAddress:  ip,
	}
	return u.auditRepo.Create(ctx, log)
}

func (u *auditLogUseCase) GetLatestLogs(ctx context.Context, limit int) ([]domain.AuditLog, error) {
	return u.auditRepo.GetAll(ctx, limit)
}
