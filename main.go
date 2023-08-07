package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server running!")

	s := echo.New()

	s.Use(MiddleWare)

	s.GET("/test", handler)

	err := s.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func handler(ctx echo.Context) error {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	s := fmt.Sprintf("Left days: %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
