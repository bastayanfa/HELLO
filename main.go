package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/auth", func(c *fiber.Ctx) error {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)

		if err != nil {
			return c.Json(nil)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": ss,
		})
	}

	app.Listen(":9090")
}

/*func xmain() {
	r := gin.New()

	r.GET("/auth", func(c *gin.Context) {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"token": ss,
		})
	})
	api := r.Group("/")
	api.Use(AuthMiddleware)

	api.PUT("/todos", todo.AddTaskfunc)
	api.PUT("/todos/:id", todo.SetDonefunc)
	api.GET("/todos", todo.GetTaskfunc)

	r.Run(":9090") // default 8080
}*/

/*
func AuthMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	mySigningKey := []byte("password")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
