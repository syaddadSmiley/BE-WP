package controller

import (
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type CategoryProductController struct {
	CategoryProductUsecase domain.CategoryProductUsecase
	Env                    *bootstrap.Env
}

func (cpc *CategoryProductController) Create(c *gin.Context) {
	var category domain.CategoryProduct

	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	category.ID = gocql.TimeUUID().String()

	err = cpc.CategoryProductUsecase.Create(c, &category)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "success create category"})
}

func (cpc *CategoryProductController) GetById(c *gin.Context) {
	id := c.Param("id")

	category, err := cpc.CategoryProductUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cpc *CategoryProductController) GetAll(c *gin.Context) {
	category, err := cpc.CategoryProductUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}
