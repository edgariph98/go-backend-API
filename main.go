package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service2ResponseType struct {
	Message string
}

func main() {
	fmt.Println("Hello World from Gin")
	r := gin.Default()
	// microservice 1 consumign microservice 2
	r.GET("/service1", func(c *gin.Context) {
		service2_URL := "http://localhost:5000/service2"
		response, request_error := http.Get(service2_URL)
		if request_error != nil {

			c.JSON(400, gin.H{
				"message": request_error,
			})
		} else {
			// decoding json response
			var svc2_resp Service2ResponseType
			json.NewDecoder(response.Body).Decode(&svc2_resp)
			// returning JSON response
			c.JSON(http.StatusAccepted, svc2_resp)
		}

	})
	// microservice 2
	r.GET("/service2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "mesasage from service 2",
		})
	})
	r.Run(":5000")
}
