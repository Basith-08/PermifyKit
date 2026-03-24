package postgres

import (
	"context"

	"github.com/starter/rbac-kit/internal/domain"
	"github.com/starter/rbac-kit/internal/repository"
	"gorm.io/gorm"
)

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) repository.AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(ctx context.Context, log *domain.AuditLog) error {
	m := &AuditLog{
		UserID:     log.UserID,
		Username:   log.Username,
		Action:     log.Action,
		Resource:   log.Resource,
		ResourceID: log.ResourceID,
		IPAddress:  log.IPAddress,
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *auditLogRepository) GetAll(ctx context.Context, limit int) ([]domain.AuditLog, error) {
	var logs []AuditLog
	query := r.db.WithContext(ctx).Order("timestamp desc")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	
	domainLogs := make([]domain.AuditLog, len(logs))
	for i, l := range logs {
		domainLogs[i] = *l.ToDomain()
	}
	return domainLogs, nil
}
