package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"eParkKtx/config"
	"eParkKtx/entities"
	"eParkKtx/repositories"

	"gorm.io/gorm"
)

type seedUser struct {
	CCCD     string `json:"cccd"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func main() {
	// Load JSON file from data/accounts.json
	dataPath := "data/accounts.json"
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		log.Fatalf("Seed file not found: %s", dataPath)
	}

	b, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Failed to read seed file: %v", err)
	}

	var users []seedUser
	if err := json.Unmarshal(b, &users); err != nil {
		log.Fatalf("Invalid JSON: %v", err)
	}

	// Connect to database (using existing config)
	config.ConnectDatabase()
	db := config.DB

	// AutoMigrate User model to ensure table exists
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Create repository
	repo := repositories.NewUserRepository()

	fmt.Println("=== Starting user seeding ===")

	// Iterate through seed users
	for _, su := range users {
		// Check if user already exists
		existingUser, err := repo.GetByID(su.CCCD)
		if err == nil && existingUser != nil {
			fmt.Printf("✓ User exists, skipping: %s (%s)\n", su.CCCD, su.Name)
			continue
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Printf("✗ Error checking user %s: %v\n", su.CCCD, err)
			continue
		}

		// Hash password using bcrypt
		hashed, err := bcrypt.GenerateFromPassword([]byte(su.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("✗ Failed to hash password for %s: %v\n", su.CCCD, err)
			continue
		}

		// Create user object
		user := &entities.User{
			UserID:      su.CCCD,
			Name:        su.Name,
			Password:    string(hashed),
			PhoneNumber: su.Phone,
			Role:        su.Role,
		}

		// Insert into database
		if err := repo.CreateNewUser(user); err != nil {
			fmt.Printf("✗ Failed to create user %s: %v\n", su.CCCD, err)
			continue
		}

		fmt.Printf("✓ Created user: %s (%s) - Role: %s\n", su.CCCD, su.Name, su.Role)
	}

	fmt.Println("=== Seeding finished successfully ===")
}
