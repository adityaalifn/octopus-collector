package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

type Content struct {
	Status    int        `json:"status"`
	Response  []Response `json:"response"`
	Timestamp time.Time  `json:"timestamp"`
}

func main() {
	ReadJSONFromFile("data.json")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", index)
	e.GET("/pods", getAllPods)
	e.GET("/pods/:id", getPodById)

	e.Logger.Fatal(e.Start(":8000"))
}

func index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func getAllPods(ctx echo.Context) error {
	response := &Content{}
	response.Status = http.StatusOK
	response.Timestamp = time.Now()
	response.Response = pods
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, response)
}

func getPodById(ctx echo.Context) error {
	id := ctx.Param("id")
	var pod []Response
	for _, v := range pods {
		if v.Id == id {
			pod = append(pod, v)
			break
		}
	}
	response := &Content{}
	response.Status = http.StatusOK
	response.Timestamp = time.Now()
	response.Response = pod
	return ctx.JSON(http.StatusOK, response)
}
