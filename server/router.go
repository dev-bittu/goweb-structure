package server

import (
  "github.com/gin-gonic/gin"
  "github.com/dev-bittu/goweb-structure/controllers"
)

func NewRouter() *gin.Engine {
  router := gin.New()

  router.Use(gin.Logger())
  router.Use(gin.Recovery())
  
  router.GET("/", controllers.Index)

  return router
}
