package route

// import (
// 	"time"

// 	"waroeng_pgn1/domain"
// 	"waroeng_pgn1/internal/api/controller"
// 	"waroeng_pgn1/internal/bootstrap"
// 	"waroeng_pgn1/iternal/repository"
// 	"waroeng_pgn1/mongo"
// 	"waroeng_pgn1/usecase"

// 	"github.com/gin-gonic/gin"
// )

// func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
// 	tr := repository.NewTaskRepository(db, domain.CollectionTask)
// 	tc := &controller.TaskController{
// 		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
// 	}
// 	group.GET("/task", tc.Fetch)
// 	group.POST("/task", tc.Create)
// }
