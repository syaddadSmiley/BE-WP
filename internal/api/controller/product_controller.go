package controller

import (
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
	Env            *bootstrap.Env
}

func (pc *ProductController) Create(c *gin.Context) {
	var product domain.Product

	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	product.ID = gocql.TimeUUID().String()

	err = pc.ProductUsecase.Create(c, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Product created successfully",
	})
}
