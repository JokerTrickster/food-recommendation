package main

import (
	"github.com/labstack/echo/v4"
)

//export PATH=$PATH:~/go/bin
func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
