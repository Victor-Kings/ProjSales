package handlers

import (
	"fmt"
	"net/http"

	"github.com/Victor-Kings/ProjSales/repository"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenericHandler interface {
	Create(ctx echo.Context) error
	ListById(ctx echo.Context) error
	// ListById(ctx echo.Context) error
	// Update(ctx echo.Context) error
	// Delete(ctx echo.Context) error
}

type ExternalModel interface { // talvez possa ser somente uma interface
	UpdateQuery(content interface{}) primitive.D
}

type GenericHandlerImp struct {
	repo  repository.GenericRepo
	model ExternalModel
}

func NewGenericHandler(externalModel ExternalModel, collectionName string) GenericHandler {
	return &GenericHandlerImp{
		repo:  repository.NewGenericRepo(externalModel, collectionName),
		model: externalModel,
	}
}

func (g *GenericHandlerImp) ListById(ctx echo.Context) error {

	id := ctx.Param("id")
	content, err := g.repo.ListById(id)
	if err != nil {
		fmt.Println("ERROR - Handler Products [listbyId] --- ", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, content)
}

func (g *GenericHandlerImp) Create(ctx echo.Context) error {
	var content interface{} = g.model

	if err := ctx.Bind(content); err != nil {
		return err
	}

	if err := g.repo.Create(content); err != nil {
		fmt.Println("ERROR - Handler Products [Create] --- ", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, content)
}

func (g *GenericHandlerImp) Update(ctx echo.Context) error {

	content := g.model

	if err := ctx.Bind(content); err != nil {
		return err
	}

	id := ctx.Param("id")
	if err := g.repo.Update(id, content); err != nil {
		fmt.Println("ERROR - Handler Products [Update] --- ", err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, content)

}
