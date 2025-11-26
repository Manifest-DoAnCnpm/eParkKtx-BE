package services

import (
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"errors"
	"log"
	"fmt"
	"time"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt" // for encrypte password

	"github.com/golang-jwt/jwt/v4"
)

/*
	WARNING: CHANG PARAMATERS OF FUNCTION TO DTO LATER
*/

/*
	User service:
	- encrypte/decrypte password
	- check username
	- phonenumber
*/

type UserService struct {
	UserRepo *repositories.UserRepo
}

// ------------------------Constructor---------------------------------------
func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

//------------------------Method---------------------------------------

func (us *UserService) CreateUser(user *entities.User) error {
    log.Printf("[UserService] Creating user: %s", user.Name)
    
    // Validate dữ liệu đầu vào
    if user.Name == "" || user.Password == "" || len(user.PhoneNumber) < 10 {
        log.Printf("[UserService] Invalid user data: name=%s, phone=%s", user.Name, user.PhoneNumber)
        return errors.New("invalid user data")
    }

    // Kiểm tra tên đã tồn tại chưa
    existingUser, err := us.UserRepo.GetByName(user.Name)
    if err == nil && existingUser != nil {
        log.Printf("[UserService] Username already exists: %s", user.Name)
        return errors.New("username already exists")
    }

    // Mã hóa mật khẩu
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("[UserService] Failed to hash password: %v", err)
        return errors.New("failed to process password")
    }
    user.Password = string(hashedPassword)

    // Tạo user mới
    if err := us.UserRepo.CreateNewUser(user); err != nil {
        log.Printf("[UserService] Failed to create user in repository: %v", err)
        return fmt.Errorf("failed to create user: %w", err)
    }

    log.Printf("[UserService] Successfully created user: %s", user.Name)
    return nil
}

// Hash password
func HashPassword(password string) (string, error) {

	// default cost có độ phức tạp là 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// copare password string with password-hash in db
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (Userv *UserService) GetUserByName(name string) (*entities.User, error) {

	ExistedUser, err := Userv.UserRepo.GetByName(name)
	if err != nil {
		return nil, err
	}

	// IsMatched := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(password))

	// if IsMatched != nil {
	// 	return nil, IsMatched
	// }

	return ExistedUser, nil
}



func (Userv *UserService) GetUserByID(ID string, Password string) (*entities.User, error) {

	ExistedUser, err := Userv.UserRepo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	IsMatched := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(Password))

	if IsMatched != nil {
		return nil, IsMatched
	}

	return ExistedUser, nil
}

//------------------------ Auth Service ---------------------------------------
// AuthService handles token generation for users.
type AuthService struct {
    UserSvc *UserService
}

// NewAuthService constructor.
func NewAuthService(userSvc *UserService) *AuthService {
    return &AuthService{UserSvc: userSvc}
}

// TokenPair holds access + refresh tokens.
type TokenPair struct {
    AccessToken  string
    RefreshToken string
}

// GenerateTokensForUser creates access and refresh JWTs for a given user.
func (as *AuthService) GenerateTokensForUser(u *entities.User) (*TokenPair, error) {
    accessExp := getEnvIntDefault("ACCESS_EXPIRE_MIN", 15)
    refreshExp := getEnvIntDefault("REFRESH_EXPIRE_MIN", 10080)

    access, err := generateJWT(u, os.Getenv("JWT_SECRET"), accessExp)
    if err != nil {
        return nil, err
    }
    refresh, err := generateJWT(u, os.Getenv("REFRESH_SECRET"), refreshExp)
    if err != nil {
        return nil, err
    }
    return &TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}

func generateJWT(u *entities.User, secret string, expireMinutes int) (string, error) {
    if secret == "" {
        return "", errors.New("jwt secret not configured")
    }
    claims := jwt.MapClaims{
        "sub":  u.UserID,
        "name": u.Name,
        "exp":  time.Now().Add(time.Duration(expireMinutes) * time.Minute).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func getEnvIntDefault(name string, def int) int {
    v := os.Getenv(name)
    if v == "" {
        return def
    }
    if i, err := strconv.Atoi(v); err == nil {
        return i
    }
    return def
}
