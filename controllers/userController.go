package controllers

import (
	"net/http"
	"os"
	"strconv"

	"eParkKtx/services"

	"github.com/gin-gonic/gin"
)

// Request payload for login
type loginCCCDRequest struct {
	CCCD     string `json:"cccd" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthController handles authentication endpoints.
type AuthController struct {
	AuthSvc *services.AuthService
	// we also rely on existing UserService for validation if needed
	UserSvc *services.UserService
}

// NewAuthController constructor
func NewAuthController(authSvc *services.AuthService, userSvc *services.UserService) *AuthController {
	return &AuthController{AuthSvc: authSvc, UserSvc: userSvc}
}

// LoginCCCD handles POST /api/v1/auth/login-cccd
// Body: { "cccd": "...", "password": "..." }
// Success: sets HttpOnly refresh_token cookie and returns access_token + user basic info.
func (ac *AuthController) LoginCCCD(c *gin.Context) {
	var req loginCCCDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "invalid request", "errors": err.Error()})
		return
	}

	// Validate credentials via existing UserService.GetUserByID (UserID == CCCD)
	user, err := ac.UserSvc.GetUserByID(req.CCCD, req.Password)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "unauthorized", "errors": "invalid credentials"})
		return
	}

	// Generate tokens
	tokens, err := ac.AuthSvc.GenerateTokensForUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "failed to generate tokens", "errors": err.Error()})
		return
	}

	// Set refresh token in HttpOnly cookie
	maxAge := getEnvIntDefault("REFRESH_EXPIRE_MIN", 10080) * 60 // seconds
	secure := os.Getenv("COOKIE_SECURE") == "true"

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})

	// Return access token & basic user info
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "login success",
		"data": gin.H{
			"access_token": tokens.AccessToken,
			"user": gin.H{
				"user_id": user.UserID,
				"name":    user.Name,
				// "role":    user.Role,
			},
		},
	})
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
