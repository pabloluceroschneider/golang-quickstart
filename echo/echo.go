package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	list := []user{ 
		{"Pablo","pablo@gmail.com"}, 
		{"daniel","daniel@gmail.com"}, 
		{"Lucero","Lucero@gmail.com"}, 
	}

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, list)
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(user)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
