package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

type CityController struct {
	CityUsecase domain.CityUsecase
	Env         *bootstrap.Env
}

func (pc *CityController) InsertCity(c *gin.Context) {
	var city domain.City
	var cityResult []domain.CityResult

	err := c.ShouldBind(&city)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	req, err := http.NewRequest("GET", "https://api.rajaongkir.com/starter/city", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("key", pc.Env.RajaOngkirKey)

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

	err = json.Unmarshal(body, &city)
	if err != nil {
		panic(err)
	}

	for _, v := range city.Rajaongkir.Results {
		cityResult = append(cityResult, domain.CityResult{
			CityID:     v.CityID,
			ProvinceID: v.ProvinceID,
			Province:   v.Province,
			Type:       v.Type,
			CityName:   v.CityName,
			PostalCode: v.PostalCode,
		})
	}

	x, err := pc.CityUsecase.InsertCity(c, cityResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, x)

}

func (pc *CityController) GetCityByProvince(c *gin.Context) {
	province := c.Param("province")
	x, err := pc.CityUsecase.GetCityByProvince(c, province)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, x)
}
