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
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// CONFIG ----------------------------------
		api.GET("/config", func(c *gin.Context) {
			bytes, err := os.ReadFile("./frontend-config.json")
			if err != nil {
				fmt.Printf("err: %v\n", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.Data(http.StatusOK, gin.MIMEJSON, bytes)
		})
		api.POST("/config", func(c *gin.Context) {
			// TODO: implement
		})

		// MODULES ----------------------------------
		api.POST("/module/:module", func(c *gin.Context) {
			slug := c.Param("module")
			body, err := c.GetRawData()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			fmt.Printf("slug: %v, body: %v\n", slug, string(body))
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"slug":    slug,
				"body":    string(body),
			})
		})
	}

}
