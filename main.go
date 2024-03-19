package main

import (
	"fmt"
	"net/http"
	"rcTest/requests"
	thirdparty "rcTest/third-party"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("quote", func(c *gin.Context) {
		var req requests.QuoteRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err})
		}

		fmt.Println("req : ", req)
		c.JSON(http.StatusOK, gin.H{
			"message": "/quote",
		})

		flightReqData := requests.CreateFlightReq(req)
		flights, err := thirdparty.RequestFlights(flightReqData)
		if err != nil {
			fmt.Println("ERR1 : ", err.Error())
		}
		fmt.Println("=> flights : ", flights)
		bestFlight := getBestFlight(flightPropertiesReq)

		flightPropertiesReq, _ := requests.CreatePropertiesReq(req)
		fmt.Println("=>req : ", flightPropertiesReq)
		
		properties, err := thirdparty.RequestProperties(flightPropertiesReq)
		if err != nil {
			fmt.Println("ERR2 : ", err.Error())
		}
		fmt.Println("=> properties : ", properties)
		// propertiesReqData := requests.CreatePropertiesReq(req)



	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
