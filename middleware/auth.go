package middleware

import (
	"course-api/models"
	"course-api/responses"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint        `json:"user_id"`
	Role   models.Role `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a given user ID and role
func GenerateToken(userID uint, role models.Role) (string, error) {
	claims := TokenClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return responses.SendError(c, fiber.StatusUnauthorized, "Missing authorization header")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return responses.SendError(c, fiber.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(*TokenClaims)
		if !ok || !token.Valid {
			return responses.SendError(c, fiber.StatusUnauthorized, "Invalid token claims")
		}

		c.Locals("user_id", claims.UserID)
		fmt.Println(claims.Role)
		c.Locals("user_role", string(claims.Role))
		return c.Next()
	}
}

// RequireRole middleware to check if user has required role
func RequireRole(roles ...models.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleStr := c.Locals("user_role")
		if roleStr == nil {
			return responses.SendError(c, fiber.StatusUnauthorized, "User role not found")
		}

		userRole := models.Role(roleStr.(string))
		if userRole == "" {
			return responses.SendError(c, fiber.StatusUnauthorized, "Invalid user role")
		}

		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}

		return responses.SendError(c, fiber.StatusForbidden, "Insufficient permissions")
	}
}
