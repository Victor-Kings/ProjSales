package main

import (
	"fmt"

	"github.com/Victor-Kings/ProjSales/handlers"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Server Started...")

	echo := echo.New()

	productsHandler := handlers.NewProductsHandler()

	echo.POST("/products", productsHandler.Create)

	echo.Logger.Fatal(echo.Start(":8080"))
}
