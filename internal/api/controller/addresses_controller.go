package controller

import (
	"net/http"
	"strings"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"
	"waroeng_pgn1/internal/tokenutil"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type AddressesController struct {
	AddressesUsecase domain.AddressesUsecase
	Env              *bootstrap.Env
}

func (pc *AddressesController) Create(c *gin.Context) {
	var addresses domain.Addresses
	//get token
	auth_header := c.Request.Header.Get("Authorization")
	//split token
	token := strings.Split(auth_header, " ")[1]

	err := c.ShouldBind(&addresses)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	addresses.ID = gocql.TimeUUID().String()
	addresses.IDUser, err = tokenutil.ExtractIDFromToken(token, bootstrap.NewEnv().AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	listAddresses, err := pc.AddressesUsecase.GetByIdUser(c, addresses.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if len(listAddresses) > 0 {
		addresses.IsDefault = false
	} else {
		addresses.IsDefault = true
	}

	err = pc.AddressesUsecase.Create(c, &addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Addresses created successfully",
	})
}

func (pc *AddressesController) GetById(c *gin.Context) {
	id := c.Param("id")

	addresses, err := pc.AddressesUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func (pc *AddressesController) GetByIdUser(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")
	token := strings.Split(auth_header, " ")[1]
	id, err := tokenutil.ExtractIDFromToken(token, bootstrap.NewEnv().AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	addressess, err := pc.AddressesUsecase.GetByIdUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, addressess)
}

func (pc *AddressesController) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var addresses domain.Addresses

	err := c.ShouldBind(&addresses)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = pc.AddressesUsecase.UpdateById(c, id, addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Addresses updated successfully",
	})
}
