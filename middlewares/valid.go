package middlewares

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
   

    "eParkKtx/services"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

// List of public endpoints that don't require authentication
var publicEndpoints = []string{
    "/api/auth/login-cccd",
    // Add more public endpoints here
}

// JWTClaims represents the JWT claims structure
type JWTClaims struct {
    UserID string `json:"sub"`
    Name   string `json:"name"`
    jwt.RegisteredClaims
}

// AuthMiddleware handles JWT token validation
type AuthMiddleware struct {
    userService *services.UserService
    secretKey   []byte
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(userService *services.UserService) *AuthMiddleware {
    secretKey := os.Getenv("JWT_SECRET")
    if secretKey == "" {
        secretKey = "default-secret-key" // Must match what's used in AuthService
    }

    return &AuthMiddleware{
        userService: userService,
        secretKey:   []byte(secretKey),
    }
}

// AuthRequired is a middleware that checks for a valid JWT token
func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Printf("🛡️  Auth Middleware - Path: %s", c.FullPath())
        
        // Skip for public endpoints
        currentPath := c.FullPath()
        log.Printf("🛡️  Checking path: %s", currentPath)
        
        for _, endpoint := range publicEndpoints {
            if currentPath == endpoint {
                log.Printf("✅ Path %s is public, skipping auth", currentPath)
                c.Next()
                return
            }
        }

        // Try to get token from Authorization header first
        authHeader := c.GetHeader("Authorization")
        var tokenString string
        
        if authHeader != "" {
            // Extract the token from the header (format: "Bearer <token>")
            parts := strings.Split(authHeader, " ")
            if len(parts) == 2 && parts[0] == "Bearer" {
                tokenString = parts[1]
                log.Printf("🛡️  Using token from Authorization header")
            }
        } else {
            // Fallback to cookie if no Authorization header
            tokenCookie, err := c.Cookie("access_token")
            if err == nil {
                tokenString = tokenCookie
                log.Printf("🛡️  Using token from cookie")
            }
        }

        if tokenString == "" {
            log.Println("❌ No authentication token found")
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "Authentication required",
            })
            c.Abort()
            return
        }

        log.Printf("🛡️  Token string: %s", tokenString)

        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate the alg is what you expect:
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return m.secretKey, nil
        })

        if err != nil {
            log.Printf("❌ JWT Validation Error: %v", err)
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "Invalid or expired token",
                "error":   err.Error(),
            })
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID, _ := claims["sub"].(string)
            name, _ := claims["name"].(string)
            
            // Set user information in context
            c.Set("userID", userID)
            c.Set("name", name)
            log.Printf("✅ Authenticated user - ID: %s, Name: %s", userID, name)
            
            c.Next()
        } else {
            log.Println("❌ Invalid token claims")
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "Invalid token claims",
            })
            c.Abort()
            return
        }
    }
}