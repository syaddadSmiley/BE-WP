package route

import (
	"database/sql"
	"time"

	"waroeng_pgn1/internal/api/controller"
	"waroeng_pgn1/internal/bootstrap"
	"waroeng_pgn1/internal/repository"
	"waroeng_pgn1/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewOrderRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	orderRepository := repository.NewOrderRepository(db, "order")
	orderUsecase := usecase.NewOrderUsecase(orderRepository, timeout)
	orderController := controller.OrderController{OrderUsecase: orderUsecase, Env: env}

	group.POST("/order", orderController.Create)
	group.GET("/order/user", orderController.GetByIdUser)
	group.PUT("/order/:id", orderController.UpdateById)
}

func NewOrderRouterGudang(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	orderRepository := repository.NewOrderRepository(db, "order")
	orderUsecase := usecase.NewOrderUsecase(orderRepository, timeout)
	orderController := controller.OrderController{OrderUsecase: orderUsecase, Env: env}

	group.PUT("/order/gudang/:id", orderController.UpdateById)
	group.GET("/order/gudang/:id", orderController.GetById)
}
