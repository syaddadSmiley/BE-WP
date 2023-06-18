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

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewProductRepository(db, domain.CollectionProduct)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
		Env:            env,
	}
	group.POST("/product/create", pc.Create)
	group.GET("/product/get/:id", pc.GetById)
	group.GET("/product/getall", pc.GetAll)
}
