package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/bootstrap"
	"waroeng_pgn1/internal/tokenutil"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
	Env          *bootstrap.Env
}

func (oc *OrderController) Create(c *gin.Context) {
	var order domain.Order
	//get token
	auth_header := c.Request.Header.Get("Authorization")
	//split token
	token := strings.Split(auth_header, " ")[1]

	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	//getId user from header authorization
	idUser, err := tokenutil.ExtractSomeFromToken(token, bootstrap.NewEnv().AccessTokenSecret, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	order.IDUser = idUser

	order.ID = gocql.TimeUUID().String()

	order.CurrentStatusOrder = "Menunggu Pembayaran"
	if order.IsRefund {
		order.IsRefund = false
	}

	err = oc.OrderUsecase.Create(c, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = oc.OrderUsecase.CreateOrderStatus(c, &domain.OrderStatus{
		ID:      gocql.TimeUUID().String(),
		IDOrder: order.ID,
		Status:  "Menunggu Pembayaran",
	})

	//count price
	fmt.Println(order)
	var totalPrice int

	for key, orderItem := range order.OrderItems {
		convPricetoInt, err := strconv.Atoi(orderItem.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		qtyInt, err := strconv.Atoi(orderItem.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		totalPrice += convPricetoInt * qtyInt
		orderItem.ID = gocql.TimeUUID().String()
		orderItem.IDOrder = order.ID
		err = oc.OrderUsecase.CreateOrderItem(c, &orderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		order.OrderItems[key] = orderItem
	}
	// order.TaxPrice = strconv.Itoa(totalPrice * 11 / 100)

	// TotalPricewithTax := totalPrice + (totalPrice * 11 / 100)
	// order.TotalPrice = strconv.Itoa(TotalPricewithTax)
	order.TotalPrice = strconv.Itoa(totalPrice)

	//body json
	dt := time.Now()

	var orderItemsString string
	jsonBytes, err := json.Marshal(order.OrderItems)
	if err != nil {
		panic(err)
	}
	orderItemsString = string(jsonBytes)
	fmt.Println(orderItemsString)

	dataBody := []byte(`
		{"transaction_details":
		{
			"order_id":"` + order.ID + `",
			"gross_amount":` + order.TotalPrice + `
		},
		"credit_card":
		{
			"secure":true
		},
		"customer_details":
		{
			"first_name":"` + order.IDUser + `"
		},
		"enabled_payments":
			["credit_card"],
		"expiry":{
			"start_time":"` + dt.Format("2006-01-02 15:04:05 +0700") + `",
			"unit":"minutes",
			"duration":120
		},
		"item_details":` + orderItemsString + `}`)

	//Midtrans Stuff Here
	var responseMidtrans domain.MidtransResponse
	req, err := http.NewRequest("POST", "https://app.sandbox.midtrans.com/snap/v1/transactions", bytes.NewBuffer(dataBody))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBody))

	req.Header.Set("authorization", "Basic"+" "+oc.Env.MidtransKey)
	req.Header.Set("content-type", "application/json")

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
	fmt.Println(string(body))

	err = json.Unmarshal(body, &responseMidtrans)
	if err != nil {
		panic(err)
	}

	if responseMidtrans.Token == "" || responseMidtrans.RedirectURL == "" {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "error while creating order"})
		return
	}

	c.JSON(http.StatusOK, responseMidtrans)
}

func (oc *OrderController) GetById(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.OrderUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetByIdUser(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	token := strings.Split(authorization, " ")[1]

	id, err := tokenutil.ExtractSomeFromToken(token, oc.Env.AccessTokenSecret, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	orders, err := oc.OrderUsecase.GetByIdUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	orderWithItems, err := oc.OrderUsecase.GetOrderItemsByIdOrder(c, orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderWithItems)
}

func (oc *OrderController) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var order domain.Order

	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	order, err = oc.OrderUsecase.UpdateById(c, id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
