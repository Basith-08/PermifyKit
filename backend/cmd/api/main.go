package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/starter/rbac-kit/internal/config"
	"github.com/starter/rbac-kit/internal/delivery/http"
	"github.com/starter/rbac-kit/internal/middleware"
	"github.com/starter/rbac-kit/internal/repository/postgres"
	"github.com/starter/rbac-kit/internal/usecase"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Setup Database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(gorm_postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate new models
	db.AutoMigrate(&postgres.AuditLog{})

	// 3. Initialize Repositories
	userRepo := postgres.NewUserRepository(db)
	roleRepo := postgres.NewRoleRepository(db)

	auditRepo := postgres.NewAuditLogRepository(db)
	auditUC := usecase.NewAuditLogUseCase(auditRepo)

	authUC := usecase.NewAuthUseCase(userRepo, roleRepo, cfg)
	userUC := usecase.NewUserUseCase(userRepo, auditUC)
	roleUC := usecase.NewRoleUseCase(roleRepo, auditUC)

	// 5. Setup Fiber Application
	app := fiber.New(fiber.Config{
		AppName: "RBAC Starter Kit API",
	})
	
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust for production
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Middlewares
	authMiddleware := middleware.AuthRequired(cfg)

	http.NewAuthHandler(app, authUC, authMiddleware)
	http.NewUserHandler(app, authUC, userUC, authMiddleware)
	http.NewRoleHandler(app, authUC, roleUC, authMiddleware)
	http.NewPermissionHandler(app, authUC, permUC, authMiddleware)
	http.NewAuditLogHandler(app, authUC, auditUC, authMiddleware)

	// Custom Check
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.SendString("OK - Backend is fully operational!")
	})

	// 7. Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
