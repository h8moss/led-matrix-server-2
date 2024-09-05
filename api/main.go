package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateApi(router *gin.Engine) {
  api := router.Group("/api")
  {
    // PING ------------------------------------
    api.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })

    // CONFIG ----------------------------------
    api.GET("/config", func(c *gin.Context) {
      bytes, err := os.ReadFile("./config.json")
      if err != nil {
        fmt.Printf("err: %v\n", err)
        c.JSON(http.StatusOK, gin.H {
          "error": err.Error(),
        })
        return
      }
      c.Data(http.StatusOK, gin.MIMEJSON, bytes)
    })
    api.POST("/config", func(c *gin.Context) {
  
    })

    // MODULES ----------------------------------

  }


}
