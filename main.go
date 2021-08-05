package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
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
}

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

/*func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}



func xmain() {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	r.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})

	api := r.NewRoute().Subrouter()
	api.Use(AuthMiddleware)
	api.HandleFunc("/todos", todo.AddTask).Methods(http.MethodPut)
	api.HandleFunc("/todos", todo.GetTask).Methods(http.MethodGet)
	api.HandleFunc("/todos/{index}", todo.SetDone).Methods(http.MethodPut)

	err := http.ListenAndServe(":9090", r)
	fmt.Print(err)
}
*/
