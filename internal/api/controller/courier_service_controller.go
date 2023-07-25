package controller

import (
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type CourierServiceController struct {
	CourierServiceUsecase domain.CourierServiceUsecase
	Env                   *bootstrap.Env
}

func (csc *CourierServiceController) Create(c *gin.Context) {
	var courierService domain.CourierService
	err := c.ShouldBindJSON(&courierService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courierService.ID = gocql.TimeUUID().String()

	err = csc.CourierServiceUsecase.Create(c, &courierService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": courierService.ID})
}

func (csc *CourierServiceController) GetById(c *gin.Context) {
	id := c.Param("id")
	courierService, err := csc.CourierServiceUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courierService)
}
