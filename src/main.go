// 1. git clone to the package, in this example case echo
// 2. Edit/personalice the package
// 3. Run this command to replace the package we were using for the edited version: go mod edit -replace github.com/labstack/echo=../echo
// 4. If you want to delete the edited package you can use: go mod edit -dropreplace github.com/labstack/echo
// 5. To package all that we are using: go mod vendor
// 6. To clean the packages that we're not using: go mod tidy

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Instanciar echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))

	fmt.Println("Hello")
}
