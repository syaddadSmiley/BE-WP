package controller

import (
	"fmt"
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
	if product.Discount == "" {
		product.Discount = "0"
	}
	if product.SoldAmount == "" {
		product.SoldAmount = "0"
	}

	err = pc.ProductUsecase.Create(c, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Product created successfully",
	})
}

func (pc *ProductController) GetById(c *gin.Context) {
	id := c.Param("id")

	product, err := pc.ProductUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(c *gin.Context) {
	products, err := pc.ProductUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetAllByCity(c *gin.Context) {
	city := c.Query("city")

	products, err := pc.ProductUsecase.GetAllByCity(c, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetAllByCategory(c *gin.Context) {
	category := c.Query("category")
	city := c.Query("city")
	fmt.Println("ANJG", category, city)

	products, err := pc.ProductUsecase.GetAllByCategory(c, category, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
