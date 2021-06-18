package products

import (
	"github.com/Victor-Kings/ProjSales/handlers"
	"github.com/Victor-Kings/ProjSales/models"
	"github.com/labstack/echo"
)

func ProductsRouters(echo *echo.Echo) {
	var m models.Products
	GenericHandler := handlers.NewGenericHandler(&m, "products")
	//echo.POST("/products", productsHandler.Create)
	//echo.GET("/products", productsHandler.ListAll)
	echo.GET("/products/:id", GenericHandler.ListById)
	//echo.POST("/products/:id", productsHandler.Update)
	//echo.DELETE("/products/:id", productsHandler.Delete)
}
