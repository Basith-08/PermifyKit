package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/middleware"
	"github.com/starter/rbac-kit/internal/usecase"
)

type AuditLogHandler struct {
	auditUC usecase.AuditLogUseCase
}

func NewAuditLogHandler(app *fiber.App, authUC usecase.AuthUseCase, auditUC usecase.AuditLogUseCase, requireAuth fiber.Handler) {
	h := &AuditLogHandler{auditUC: auditUC}

	auditGroup := app.Group("/api/audit", requireAuth)
	auditGroup.Get("/", middleware.RequirePermission(authUC, "users", "read"), h.GetLogs)
}

func (h *AuditLogHandler) GetLogs(c *fiber.Ctx) error {
	logs, err := h.auditUC.GetLatestLogs(c.Context(), 50)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(logs)
}
