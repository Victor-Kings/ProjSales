package products

import (
	"github.com/Victor-Kings/ProjSales/handlers"
	"github.com/labstack/echo"
)

func ProductsRouters(echo *echo.Echo) {
	productsHandler := handlers.NewProductsHandler()
	echo.POST("/products", productsHandler.Create)
	echo.GET("/products", productsHandler.ListAll)
	echo.GET("/products/:id", productsHandler.ListById)
	echo.POST("/products/:id", productsHandler.Update)
	echo.DELETE("/products/:id", productsHandler.Delete)
}
