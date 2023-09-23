package server

import (
  "github.com/gin-gonic/gin"
)

func Init(){
  router := NewRouter()
  
  router.GET("/ping", func (c *gin.Context)  {
    c.JSON(200, gin.H{
      "msg": "pong",
    })
  })

  router.Run()
}
