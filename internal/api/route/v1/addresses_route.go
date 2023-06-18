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

func NewAddressesRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewAddressesRepository(db, domain.CollectionAddresses)
	pc := &controller.AddressesController{
		AddressesUsecase: usecase.NewAddressesUsecase(pr, timeout),
		Env:              env,
	}
	group.POST("/addresses/create", pc.Create)
	group.GET("/addresses/get/:id", pc.GetById)
	group.GET("/addresses/getbyiduser", pc.GetByIdUser)
	group.PUT("/addresses/update/:id", pc.UpdateById)
}
