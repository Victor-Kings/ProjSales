package handlers

import (
	"net/http"

	"github.com/Victor-Kings/ProjSales/models"
	"github.com/labstack/echo"
)

type ProductsHandler interface {
	Create(ctx echo.Context) error
	ListAll(ctx echo.Context) error
	ListById(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
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
	product := models.Products{}

	if err := ctx.Bind(&product); err != nil {
		return err
	}

	if err := p.db.CreateProducts(product); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (p *ProductsHandlerImp) ListAll(ctx echo.Context) error {

	products, err := p.db.ListAllProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, products)
}

func (p *ProductsHandlerImp) ListById(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := p.db.ListByIdProducts(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, product)
}

func (p *ProductsHandlerImp) Update(ctx echo.Context) error {

	product := models.Products{}

	if err := ctx.Bind(&product); err != nil {
		return err
	}

	id := ctx.Param("id")
	if err := p.db.UpdateProduct(id, product); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, product)

}

func (p *ProductsHandlerImp) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := p.db.DeleteProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, id)
}
