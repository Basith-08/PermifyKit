package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starter/rbac-kit/internal/usecase"
)

// RequirePermission checks if the authenticated user has the required permission 'resource.action'
func RequirePermission(authUC usecase.AuthUseCase, resource, action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		perms, err := authUC.GetUserPermissions(context.Background(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check permissions"})
		}

		hasPermission := false
		for _, p := range perms {
			if p.Resource == resource && p.Action == action {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf("forbidden: requires %s.%s", resource, action),
			})
		}

		return c.Next()
	}
}
