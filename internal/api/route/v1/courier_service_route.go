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

func NewCourierServiceRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	csr := repository.NewCourierServiceRepository(db, domain.CollectionCourier)
	csc := &controller.CourierServiceController{
		CourierServiceUsecase: usecase.NewCourierServiceUsecase(csr, timeout),
		Env:                   env,
	}

	group.POST("/courier_service/create", csc.Create)
	group.GET("/courier_service/getbyid/:id", csc.GetById)
}
