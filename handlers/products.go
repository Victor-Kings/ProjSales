package handlers

import (
	"net/http"

	"github.com/Victor-Kings/ProjSales/models"
	"github.com/labstack/echo"
)

type ProductsHandler interface {
	Create(ctx echo.Context) error
}
type ProductsHandlerImp struct {
	db models.ProductsDB
}

func NewProductsHandler() ProductsHandler {
	return &ProductsHandlerImp{
		db: models.NewProductsDB(),
	}
}

func (p *ProductsHandlerImp) Create(ctx echo.Context) error {
	var product models.Products

	if err := ctx.Bind(&product); err != nil {
		return err
	}

	if err := p.db.CreateProducts(product); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}
