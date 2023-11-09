// authMiddleware.go
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

// authMiddleware.go
// package middleware

// import (
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gofiber/fiber/v2"
// )

// func JWTMiddleware(c *fiber.Ctx) error {

// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}

// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("your-secret-key"), nil
// 	})

// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}

// 	c.Locals("user", claims)
// 	return c.Next()

// }
