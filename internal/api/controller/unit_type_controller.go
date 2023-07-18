package controller

import (
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type UnitTypeController struct {
	UnitTypeUsecase domain.UnitTypeUsecase
	Env             *bootstrap.Env
}

func (utc *UnitTypeController) Create(c *gin.Context) {
	var unitType domain.UnitType

	err := c.ShouldBind(&unitType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	unitType.ID = gocql.TimeUUID().String()

	err = utc.UnitTypeUsecase.Create(c, &unitType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "UnitType created successfully",
	})
}

func (utc *UnitTypeController) GetById(c *gin.Context) {
	id := c.Param("id")

	unitType, err := utc.UnitTypeUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, unitType)
}
