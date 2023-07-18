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

func NewUnitTypeRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewUnitTypeRepository(db, domain.CollectionUnitType)
	pc := &controller.UnitTypeController{
		UnitTypeUsecase: usecase.NewUnitTypeUsecase(pr, timeout),
		Env:             env,
	}
	group.POST("/unitType/create", pc.Create)
	group.GET("/unitType/get/:id", pc.GetById)
}
