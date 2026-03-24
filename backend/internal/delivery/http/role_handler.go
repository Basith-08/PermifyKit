package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/middleware"
	"github.com/starter/rbac-kit/internal/usecase"
)

type RoleHandler struct {
	roleUC usecase.RoleUseCase
}

func NewRoleHandler(app *fiber.App, authUC usecase.AuthUseCase, roleUC usecase.RoleUseCase, requireAuth fiber.Handler) {
	h := &RoleHandler{roleUC: roleUC}

	rolesGroup := app.Group("/api/roles", requireAuth)

	rolesGroup.Get("/", middleware.RequirePermission(authUC, "roles", "read"), h.GetAll)
	rolesGroup.Get("/:id", middleware.RequirePermission(authUC, "roles", "read"), h.GetByID)
	rolesGroup.Post("/", middleware.RequirePermission(authUC, "roles", "write"), h.Create)
	rolesGroup.Put("/:id", middleware.RequirePermission(authUC, "roles", "write"), h.Update)
	rolesGroup.Delete("/:id", middleware.RequirePermission(authUC, "roles", "write"), h.Delete)
}

func (h *RoleHandler) GetAll(c *fiber.Ctx) error {
	roles, err := h.roleUC.GetAllRoles(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(roles)
}

func (h *RoleHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	role, err := h.roleUC.GetRoleByID(c.Context(), id)
	if err != nil || role == nil {
		return c.Status(404).JSON(fiber.Map{"error": "role not found"})
	}
	return c.JSON(role)
}

type CreateRoleRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	PermissionIDs []int  `json:"permission_ids"`
}

func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var req CreateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	role, err := h.roleUC.CreateRole(c.Context(), req.Name, req.Description, req.PermissionIDs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(role)
}

func (h *RoleHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	var req CreateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	role, err := h.roleUC.UpdateRole(c.Context(), id, req.Name, req.Description, req.PermissionIDs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(role)
}

func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	if err := h.roleUC.DeleteRole(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
