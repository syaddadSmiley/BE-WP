package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type CourierController struct {
	CourierUsecase domain.CourierUsecase
	Env            *bootstrap.Env
}

func (cc *CourierController) Create(c *gin.Context) {
	var courier domain.Courier

	err := c.ShouldBind(&courier)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	courier.ID = gocql.TimeUUID().String()
	err = cc.CourierUsecase.Create(c, &courier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Courier created successfully",
	})

}

func (cc *CourierController) GetAll(c *gin.Context) {
	couriers, err := cc.CourierUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, couriers)
}

func (cc *CourierController) GetServiceList(c *gin.Context) {
	var request domain.CostRequest
	var response domain.CostResponse

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	fmt.Println(request)

	idCity, err := cc.CourierUsecase.GetIdCityByName(c, request.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	dataBody := url.Values{}
	dataBody.Set("origin", idCity)
	dataBody.Set("destination", idCity)
	dataBody.Set("weight", strconv.Itoa(request.Weight))
	dataBody.Set("courier", request.Courier)

	req, err := http.NewRequest("POST", "https://api.rajaongkir.com/starter/cost", strings.NewReader(dataBody.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("key", cc.Env.RajaOngkirKey)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.Rajaongkir.Results[0])

}
