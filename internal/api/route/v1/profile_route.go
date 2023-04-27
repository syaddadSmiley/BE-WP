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

// func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
// 	ur := repository.NewUserRepository(db, domain.CollectionUser)
// 	pc := &controller.ProfileController{
// 		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
// 	}
// 	group.GET("/profile", pc.Fetch)
// }
