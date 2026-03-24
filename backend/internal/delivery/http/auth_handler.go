package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/usecase"
)

type AuthHandler struct {
	authUC usecase.AuthUseCase
}

func NewAuthHandler(app *fiber.App, authUC usecase.AuthUseCase, authMiddleware fiber.Handler) {
	handler := &AuthHandler{
		authUC: authUC,
	}

	app.Post("/api/auth/login", handler.Login)
	
	protected := app.Group("/api/user", authMiddleware)
	protected.Get("/me", handler.GetMe)
	protected.Get("/permissions", handler.GetPermissions)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	token, err := h.authUC.Login(c.Context(), req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	user, err := h.authUC.GetUserProfile(c.Context(), userID)
	if err != nil || user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

func (h *AuthHandler) GetPermissions(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	perms, err := h.authUC.GetUserPermissions(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	// Convert array of structs to flat array of strings "resource.action" for easy parsing on frontend
	permStrings := make([]string, len(perms))
	for i, p := range perms {
		permStrings[i] = p.Resource + "." + p.Action
	}
	
	return c.JSON(fiber.Map{
		"permissions": permStrings,
	})
}
