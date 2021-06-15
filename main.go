package main

import (
	"fmt"

	"github.com/Victor-Kings/ProjSales/routers/products"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Server Started...")

	echo := echo.New()

	products.ProductsRouters(echo)

	echo.Logger.Fatal(echo.Start(":8080"))
}
