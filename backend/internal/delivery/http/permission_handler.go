package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/middleware"
	"github.com/starter/rbac-kit/internal/usecase"
)

type PermissionHandler struct {
	permUC usecase.PermissionUseCase
}

func NewPermissionHandler(app *fiber.App, authUC usecase.AuthUseCase, permUC usecase.PermissionUseCase, requireAuth fiber.Handler) {
	h := &PermissionHandler{permUC: permUC}

	permsGroup := app.Group("/api/permissions", requireAuth)
	permsGroup.Get("/", middleware.RequirePermission(authUC, "roles", "read"), h.GetAll)
}

func (h *PermissionHandler) GetAll(c *fiber.Ctx) error {
	perms, err := h.permUC.GetAllPermissions(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(perms)
}
