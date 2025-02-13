package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateApi(router *gin.Engine, channel chan string) {
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
			body, err := c.GetRawData()

			err1 := os.WriteFile("./frontend-config.json", body, 0666)

			if err != nil || err1 != nil {
				fmt.Printf("err: %v\n", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
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

			var bodyMap map[string]interface{}
			json.Unmarshal(body, bodyMap)

			var resultParts []string
			resultParts = append(resultParts, slug)
			for k, v := range bodyMap {
				resultParts = append(resultParts, fmt.Sprintf("%s:%v", k, v))
			}
			var result = strings.Join(resultParts, " ")
			fmt.Printf("%s", result)

			channel <- result

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"slug":    slug,
				"body":    string(body),
			})
		})
	}
}
