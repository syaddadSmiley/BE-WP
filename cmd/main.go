package main

import (
	"fmt"
	"time"

	routeV1 "waroeng_pgn1/internal/api/route/v1"
	"waroeng_pgn1/internal/bootstrap"

	"github.com/gin-gonic/gin"

	"waroeng_pgn1/internal/database"
)

func main() {

	// app := bootstrap.App()

	env := bootstrap.NewEnv()
	fmt.Println(env)
	db, _ := database.ConnectToDB()
	// defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routerV1 := gin.Group("v1")

	routeV1.Setup(env, timeout, db, routerV1)

	gin.Run(env.ServerAddress)
}
