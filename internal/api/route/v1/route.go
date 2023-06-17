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
	// NewRefreshTokenRouter(env, timeout, db, publicRouterV1)

	protectedAdminRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedAdminRouterV1.Use(middleware.JwtAuthAdminMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProductRouter(env, timeout, db, protectedAdminRouterV1)
	// NewProfileRouter(env, timeout, db, protectedRouterV1)
	// NewTaskRouter(env, timeout, db, protectedRouterV1)
}
