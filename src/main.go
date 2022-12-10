package main

import (
    "github.com/labstack/echo"
    "src/routing"
)

func main() {
    e := echo.New()

    e.GET("/trees", routing.GetTree)

    e.Logger.Fatal(e.Start(":8080"))
}
