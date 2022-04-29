package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

//var quotes []quote = make([]quote, 0)

var quotes []quote = []quote{
	{"Learn to Lead", "SVIT"},
	{"People rarely succeed unless they have fun in what they are doing", "person_ABC"},
	{"Even though you're growing up, you should never stop having fun", "Person_xyz"},
	{"A person must never cease striving to enjoy life", "mark"},
}

func main() {
	//rand.Seed(time.Now().Unix())
	api := echo.New()
	api.GET("/home", hello)
	api.POST("/home", hello)
	api.PUT("/home", anotherHandler)
	api.GET("/quotes", getquotes)
	api.GET("/quotes/random", getRandomQuote)

	//getting our own assigned port
	port := os.Getenv("port")
	/*if port == "" {
		port = "32445"
	}*/
	api.Start(":" + port)
	//api.Start(":9090")
}

//hello is a handler
func hello(c echo.Context) error {
	return c.JSON(200, "hello barbie123")
}

func anotherHandler(c echo.Context) error {
	return c.NoContent(200)
}

func getquotes(c echo.Context) error {
	return c.JSON(200, quotes)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}
