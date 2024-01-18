package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func postQuote(ctx *gin.Context) {
	var newdataQuote dataQuote
	//var allFlights []flight

	// Call BindJSON to bind the received JSON to
	// newdataQuote.
	if err := ctx.BindJSON(&newdataQuote); err != nil {
		return
	}

	// 1. get quote flights
	flights := getFlights(newdataQuote.DepartureDate, newdataQuote.DepartureIATA, newdataQuote.DestinationIATA)
	fmt.Println("Flights..." + flights)
	// 2. get quote propreties
	propreties := getPropreties(newdataQuote.DepartureDate, newdataQuote.DestinationIATA)
	fmt.Println("propreties..." + propreties)
	// make the response

}
