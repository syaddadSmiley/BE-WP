package route

// import (
// 	"database/sql"
// 	"time"

// 	"waroeng_pgn1/api/controller"
// 	"waroeng_pgn1/bootstrap"
// 	"waroeng_pgn1/domain"
// 	"waroeng_pgn1/repository"
// 	"waroeng_pgn1/usecase"

// 	"github.com/gin-gonic/gin"
// )

// func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
// 	ur := repository.NewUserRepository(db, domain.CollectionUser)
// 	rtc := &controller.RefreshTokenController{
// 		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
// 		Env:                 env,
// 	}
// 	group.POST("/refresh", rtc.RefreshToken)
// }
