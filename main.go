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
	echo.GET("/products", productsHandler.ListAll)

	echo.Logger.Fatal(echo.Start(":8080"))
}
