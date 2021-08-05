package fizzbuzz

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Say(n int) string {
	if n%15 == 0 {
		return ("FizzBuzz")
	}
	if n%3 == 0 {
		return ("Fizz")
	}
	if n%5 == 0 {
		return ("Buzz")
	}

	return strconv.Itoa(n)
}

type Intner interface {
	Intn(n int) int
}

func RandomFizzBuzz(rd Intner) string {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	n := rd.Intn(100) + 1

	return Say(n)
}

type RandomFizzBuzzHandler struct {
	rd Intner
}

func (r RandomFizzBuzzHandler) Handler(c *gin.Context) string {
	n := r.rd.Intn(100) + 1

	return Say(n)
}
