package main

import (
	"log"
	"os"
	"time"

	"eParkKtx/config"
	"eParkKtx/controllers"
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"eParkKtx/routes"
	"eParkKtx/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	PayOSClientID    = "07622372-99f6-4a21-9376-52932d63d091"
	PayOSApiKey      = "ac6155d0-0d3e-46c8-8c42-49975f7956d8"
	PayOSChecksumKey = "fdd5d4c9d46d2b984e115a285b08b97243a919ebbbfae812bf2593206d2e324c"
)

// initSampleData khởi tạo dữ liệu mẫu cho các bảng
func initSampleData(db *gorm.DB) error {
	// Tạo user quản lý ký túc xá
	dormManager := entities.User{
		UserID:      "dorm_manager",
		Name:        "Nguyễn Văn A",
		Password:    "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
		PhoneNumber: "0912345678",
		DoB:         time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gender:      "Nam",
	}

	// Tạo user quản lý bãi đỗ xe 1
	parkManager1 := entities.User{
		UserID:      "park_manager1",
		Name:        "Trần Thị B",
		Password:    "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
		PhoneNumber: "0912345679",
		DoB:         time.Date(1991, 2, 2, 0, 0, 0, 0, time.UTC),
		Gender:      "Nữ",
	}

	// Tạo user quản lý bãi đỗ xe 2
	parkManager2 := entities.User{
		UserID:      "park_manager2",
		Name:        "Lê Văn C",
		Password:    "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
		PhoneNumber: "0912345680",
		DoB:         time.Date(1992, 3, 3, 0, 0, 0, 0, time.UTC),
		Gender:      "Nam",
	}

	// Tạo các user nếu chưa tồn tại
	users := []entities.User{dormManager, parkManager1, parkManager2}
	for _, user := range users {
		if err := db.FirstOrCreate(&entities.User{}, user).Error; err != nil {
			return err
		}
	}

	// Tạo dữ liệu mẫu cho bảng DormitoryManagement (1 user quản lý nhiều tòa)
	var dormCount int64
	db.Model(&entities.DormitoryManagement{}).Count(&dormCount)
	if dormCount == 0 {
		dorms := []entities.DormitoryManagement{
			{UserID: "dorm_manager", Building: "Tòa A"},
		}
		if err := db.Create(&dorms).Error; err != nil {
			return err
		}
		log.Println("Đã thêm dữ liệu mẫu cho bảng DormitoryManagement")
	}

	// Tạo dữ liệu mẫu cho bảng ParkManagement (mỗi user quản lý một bãi đỗ)
	var parkCount int64
	db.Model(&entities.ParkManagement{}).Count(&parkCount)
	if parkCount == 0 {
		parks := []entities.ParkManagement{
			{UserID: "park_manager1", ParkName: "Bãi đỗ xe KTX A"},
			{UserID: "park_manager2", ParkName: "Bãi đỗ xe KTX B"},
		}
		if err := db.Create(&parks).Error; err != nil {
			return err
		}
		log.Println("Đã thêm dữ liệu mẫu cho bảng ParkManagement")
	}

	return nil
}

func main() {

	// Load .env file
	_ = godotenv.Load(".env/.env")

	// Kết nối database SQLite
	config.ConnectDatabase()
	db := config.DB

	// Tự động tạo bảng nếu chưa tồn tại
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Student{},
		&entities.Vehicle{},
		&entities.Contract{},
		&entities.DormitoryManagement{},
		&entities.EEHistory{},
		&entities.Garage{},
		&entities.ParkManagement{},
	)
	if err != nil {
		log.Fatalf("Không thể migrate database: %v", err)
	}

	// Khởi tạo dữ liệu mẫu
	if err := initSampleData(db); err != nil {
		log.Fatalf("Không thể khởi tạo dữ liệu mẫu: %v", err)
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

	// Khởi tạo PayOS service
	payOSService := services.NewPayOSService(
		PayOSClientID,
		PayOSApiKey,
		PayOSChecksumKey,
	)

	// Khởi tạo ParkManagement repo, service và controller
	parkManagementRepo := repositories.NewParkManagementRepo(userRepo)
	parkManagementService := services.NewParkManagementService(parkManagementRepo)
	parkManagementController := controllers.NewParkManagementController(parkManagementService)

	// Khởi tạo controllers
	studentController := &controllers.StudentController{
		StudentService: studentService,
	}

	// Khởi tạo payment controller
	paymentController := controllers.NewPaymentController(payOSService)
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
	routes.SetupStudentRoutes(r, studentController,userService)
	routes.SetupParkManagementRoutes(r, parkManagementController,userService)
	routes.SetupPaymentRoutes(r, paymentController)
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
