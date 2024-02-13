package handlers

import (
	"fmt"
	"net/http"
	"technical-test/models"
	"technical-test/pkg"

	"github.com/gin-gonic/gin"
)

func RetrieveStayQuoteHandler(c *gin.Context) {
	var reqBody models.FightJson
	err := c.Bind(&reqBody)

	if err != nil {
		http.Error(c.Writer, "500 "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Check for must have fields

	//get flight prices
	flights, err := pkg.CallProgrammadeus(reqBody)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(flights)
	//get hotel prices
	hotelFinder := models.HotelFinderInfos{
		Duration:      reqBody.Duration,
		DepartureDate: reqBody.DepartureDate,
		Iata:          reqBody.DestinationIATA,
	}
	hotels, err := pkg.CallHotelcodes(hotelFinder)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(hotels)
	//calculates and parse data

	var resp models.FullQuote

	

	c.JSON(http.StatusOK, hotels)
}
