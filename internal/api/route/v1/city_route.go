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

func NewCityRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewCityRepository(db, domain.CollectionCity)
	pc := &controller.CityController{
		CityUsecase: usecase.NewCityUsecase(pr, timeout),
		Env:         env,
	}
	group.GET("/city/insert", pc.InsertCity)
	group.GET("/city/get/:province", pc.GetCityByProvince)
}
