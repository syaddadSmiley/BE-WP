package route

import (
	"database/sql"
	"time"

	"waroeng_pgn1/internal/api/controller"
	"waroeng_pgn1/internal/bootstrap"
	"waroeng_pgn1/internal/repository"
	"waroeng_pgn1/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewCategoryProductRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	categoryProductRepository := repository.NewCategoryProductRepository(db, "category_product")
	categoryProductUsecase := usecase.NewCategoryUsecase(categoryProductRepository, timeout)
	categoryProductController := controller.CategoryProductController{
		CategoryProductUsecase: categoryProductUsecase,
		Env:                    env,
	}

	group.POST("/category", categoryProductController.Create)
	group.GET("/category/:id", categoryProductController.GetById)
	group.GET("/category", categoryProductController.GetAll)
}
