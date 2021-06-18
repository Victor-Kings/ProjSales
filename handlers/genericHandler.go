package handlers

import (
	"fmt"
	"net/http"

	"github.com/Victor-Kings/ProjSales/repository"
	"github.com/labstack/echo"
)

type GenericHandler interface {
	// /	Create(ctx echo.Context) error
	ListById(ctx echo.Context) error
	// ListById(ctx echo.Context) error
	// Update(ctx echo.Context) error
	// Delete(ctx echo.Context) error
}

type GenericHandlerImp struct {
	repo           repository.GenericRepo
	model          interface{}
	collectionName string
}

func NewGenericHandler(t interface{}, collectionName string) GenericHandler {
	return &GenericHandlerImp{
		repo:           repository.NewGenericRepo(),
		model:          t,
		collectionName: collectionName,
	}
}

func (g GenericHandlerImp) ListById(ctx echo.Context) error {

	id := ctx.Param("id")
	product, err := g.repo.ListById(g.model, id, g.collectionName)
	if err != nil {
		fmt.Println("ERROR - Handler Products [listbyId] --- ", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, product)
}
