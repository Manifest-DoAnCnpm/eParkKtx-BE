package main

import (
	"log"
	"os"

	"eParkKtx/config"
	"eParkKtx/controllers"
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"eParkKtx/routes"
	"eParkKtx/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	_ = godotenv.Load(".env/.env")

	// Kết nối database SQLite
	config.ConnectDatabase()
	db := config.DB

	// Tự động tạo bảng nếu chưa tồn tại
	err := db.AutoMigrate(&entities.User{}, &entities.Student{})
	if err != nil {
		log.Fatalf("Không thể migrate database: %v", err)
	}

	// Khởi tạo repositories
	userRepo := repositories.NewUserRepository()
	userRepo.DB = db

	// Khởi tạo student repo với userRepo
	studentRepo := &repositories.StudentRepo{
		UserRepo: userRepo,
	}

	// Khởi tạo services
	userService := services.NewUserService(userRepo)
	// Khởi tạo student service với userService và studentRepo
	studentService := &services.StudentService{
		UserService: userService,
		StudentRepo: studentRepo,
	}

	// Khởi tạo controllers
	studentController := &controllers.StudentController{
		StudentService: studentService,
	}

	// Auth services + controller
	authService := services.NewAuthService(userService)
	authController := controllers.NewAuthController(authService, userService)

	// Khởi tạo Gin router
	r := gin.Default()

	// Cấu hình CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Thiết lập routes
	routes.SetupStudentRoutes(r, studentController)
	routes.AuthRoutes(r, authController)

	// Chạy server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server đang chạy tại http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Không thể khởi động server: %v", err)
	}
}
