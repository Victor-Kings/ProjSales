package custumer

import (
	"github.com/Victor-Kings/ProjSales/handlers"
	"github.com/Victor-Kings/ProjSales/models"
	"github.com/labstack/echo"
)

func CustumerRouters(echo *echo.Echo) {
	var m models.Custumer
	GenericHandler := handlers.NewGenericHandler(&m, "custumer")
	echo.POST("/custumer", GenericHandler.Create)
	//echo.GET("/products", productsHandler.ListAll)
	echo.GET("/custumer/:id", GenericHandler.ListById)
	//echo.POST("/products/:id", productsHandler.Update)
	//echo.DELETE("/products/:id", productsHandler.Delete)
}
