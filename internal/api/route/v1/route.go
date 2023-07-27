package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"

	"waroeng_pgn1/internal/api/middleware"
	"waroeng_pgn1/internal/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouterV1)
	NewLoginRouter(env, timeout, db, publicRouterV1)
	NewProductRouter(env, timeout, db, publicRouterV1)
	NewProvinceRouter(env, timeout, db, publicRouterV1)
	NewCityRouter(env, timeout, db, publicRouterV1)
	NewCourierRouter(env, timeout, db, publicRouterV1)
	// NewRefreshTokenRouter(env, timeout, db, publicRouterV1)

	protectedAdminRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedAdminRouterV1.Use(middleware.JwtAuthGudangMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProductRouterGudang(env, timeout, db, protectedAdminRouterV1)
	NewCourierRouterGudang(env, timeout, db, protectedAdminRouterV1)
	NewUnitTypeRouter(env, timeout, db, protectedAdminRouterV1)
	NewCategoryProductRouter(env, timeout, db, protectedAdminRouterV1)
	NewOrderRouterGudang(env, timeout, db, protectedAdminRouterV1)

	protectedUserRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedUserRouterV1.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewAddressesRouter(env, timeout, db, protectedUserRouterV1)
	NewOrderRouter(env, timeout, db, protectedUserRouterV1)
	NewCourierServiceRouter(env, timeout, db, protectedUserRouterV1)

}
