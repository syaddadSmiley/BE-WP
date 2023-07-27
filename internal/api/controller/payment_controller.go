package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"waroeng_pgn1/domain"
// 	"waroeng_pgn1/internal/bootstrap"
// 	"waroeng_pgn1/internal/tokenutil"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gocql/gocql"
// )

// type PaymentController struct {
// 	PaymentUsecase domain.PaymentUsecase
// 	Env          *bootstrap.Env
// }

// func (pc *PaymentController) CreatePayment(c *gin.Context) {
// 	var payment domain.Payment
// 	err := c.ShouldBind(&payment)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	payment.ID = gocql.TimeUUID().String()

// 	err = pc.PaymentUsecase.Create(c, &payment)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, payment)
// }

// func (pc *PaymentController) ValidatePayment(c *gin.Context) {
// 	var payment domain.Payment
// 	err := c.ShouldBind(&payment)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	//Midtrans Stuff Here
// 	var responseMidtrans domain.MidtransTransactionStatusRespone
// 	req, err := http.NewRequest("GET", "https://api.sandbox.midtrans.com/v2/6f01094f-2b8f-11ee-973d-0a002700000f/status", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req.Header.Set("authorization", "Basic"+" "+pc.Env.MidtransKey)
// 	req.Header.Set("content-type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(body))

// 	err = json.Unmarshal(body, &responseMidtrans)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if(responseMidtrans.TransactionStatus == "capture"){
// 		payment.Status = "success"
// 	} else {
// 		payment.Status = "failed"
// 	}

// 	paymentsRes, err := pc.PaymentUsecase.UpdateById(c, payment.ID, payment)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, paymentsRes)

// }
