package route

import (
	"database/sql"
	"time"

	"waroeng_pgn1/domain"
	"waroeng_pgn1/internal/api/controller"
	"waroeng_pgn1/internal/bootstrap"
	"waroeng_pgn1/internal/repository"
	"waroeng_pgn1/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewCourierRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewCourierRepository(db, domain.CollectionCourier)
	pc := &controller.CourierController{
		CourierUsecase: usecase.NewCourierUsecase(pr, timeout),
		Env:            env,
	}

	group.GET("/courier/getall", pc.GetAll)
	group.POST("/courier/getservice", pc.GetServiceList)

}

func NewCourierRouterGudang(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewCourierRepository(db, domain.CollectionCourier)
	pc := &controller.CourierController{
		CourierUsecase: usecase.NewCourierUsecase(pr, timeout),
		Env:            env,
	}
	group.POST("/courier/create", pc.Create)
}
