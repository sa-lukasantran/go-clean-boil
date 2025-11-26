package controller

import (
	_ "go-basesource/docs"
	"go-basesource/src/adapters"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize() {
	router := adapters.GetHttpServer()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	NewExampleController(router)
}
