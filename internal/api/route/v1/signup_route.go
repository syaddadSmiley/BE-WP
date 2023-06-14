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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
