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

func NewProvinceRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewProvinceRepository(db, domain.CollectionProvince)
	pc := &controller.ProvinceController{
		ProvinceUsecase: usecase.NewProvinceUsecase(pr, timeout),
		Env:             env,
	}
	group.GET("/province/insert", pc.InsertProvince)
	group.GET("/province/get", pc.GetProvince)
}
