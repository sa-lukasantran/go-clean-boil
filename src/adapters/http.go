package adapters

import "github.com/gin-gonic/gin"

var Router *gin.Engine

type HttpServer struct {
	*gin.Engine
}

// AppContextHandler is a handler function that uses AppContext instead of gin.Context
type AppContextHandler func(*AppContext)

func GetHttpServer() *HttpServer {
	if Router != nil {
		return &HttpServer{Router}
	}
	Router = gin.Default()
	return &HttpServer{Router}
}

func StartHttpServer() {
	server := GetHttpServer()
	server.Run(":8080")
}
