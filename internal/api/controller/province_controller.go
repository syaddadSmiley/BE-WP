package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	ProvinceUsecase domain.ProvinceUsecase
	Env             *bootstrap.Env
}

func (pc *ProvinceController) InsertProvince(c *gin.Context) {
	var province domain.Province
	var provinceResult []domain.ProvinceResult

	err := c.ShouldBind(&province)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	req, err := http.NewRequest("GET", "https://api.rajaongkir.com/starter/province", nil)
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

	err = json.Unmarshal(body, &province)
	if err != nil {
		panic(err)
	}

	for _, v := range province.Rajaongkir.Results {
		provinceResult = append(provinceResult, domain.ProvinceResult{
			ProvinceID: v.ProvinceID,
			Province:   v.Province,
		})
	}

	x, err := pc.ProvinceUsecase.InsertProvince(c, provinceResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, x)

}

func (pc *ProvinceController) GetProvince(c *gin.Context) {
	province, err := pc.ProvinceUsecase.GetProvince(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, province)
}
