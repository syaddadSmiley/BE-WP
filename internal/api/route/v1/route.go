package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"

	// "waroeng_pgn1/api/middleware"
	"waroeng_pgn1/internal/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gocql.Session, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouterV1)
	NewLoginRouter(env, timeout, db, publicRouterV1)
	// NewRefreshTokenRouter(env, timeout, db, publicRouterV1)

	// protectedRouterV1 := routerV1.Group("")
	// // Middleware to verify AccessToken
	// protectedRouterV1.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// // All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouterV1)
	// NewTaskRouter(env, timeout, db, protectedRouterV1)
}
