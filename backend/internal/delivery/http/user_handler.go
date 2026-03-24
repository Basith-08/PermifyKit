package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/middleware"
	"github.com/starter/rbac-kit/internal/usecase"
)

type UserHandler struct {
	userUC usecase.UserUseCase
}

func NewUserHandler(app *fiber.App, authUC usecase.AuthUseCase, userUC usecase.UserUseCase, requireAuth fiber.Handler) {
	h := &UserHandler{userUC: userUC}

	usersGroup := app.Group("/api/users", requireAuth)

	usersGroup.Get("/", middleware.RequirePermission(authUC, "users", "read"), h.GetAll)
	usersGroup.Get("/:id", middleware.RequirePermission(authUC, "users", "read"), h.GetByID)
	usersGroup.Post("/", middleware.RequirePermission(authUC, "users", "write"), h.Create)
	usersGroup.Put("/:id", middleware.RequirePermission(authUC, "users", "write"), h.Update)
	usersGroup.Delete("/:id", middleware.RequirePermission(authUC, "users", "write"), h.Delete)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userUC.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	user, err := h.userUC.GetUserByID(c.Context(), id)
	if err != nil || user == nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleIDs  []int  `json:"role_ids"`
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	user, err := h.userUC.CreateUser(c.Context(), req.Username, req.Email, req.Password, req.RoleIDs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	user, err := h.userUC.UpdateUser(c.Context(), id, req.Username, req.Email, req.Password, req.RoleIDs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id format"})
	}

	if err := h.userUC.DeleteUser(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
