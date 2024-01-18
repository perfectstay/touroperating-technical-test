package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() {

	router := gin.Default()
	router.POST("/quote", postQuote)
	router.Run("localhost:4242")

}

func getFlights(depDate, DepIATA, DestIATA string) string {
	fmt.Println("Flights...")

	flightURL := "http://localhost:3001/quoteFlights"
	userData := []byte(`{"duration":7, "departureDate":"` + depDate + `","departureIATA":"` + DepIATA + `","destinationIATA":"` + DestIATA + `"}`)

	// create new http request
	request, error := http.NewRequest("POST", flightURL, bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if error != nil {
		fmt.Println(error)
	}

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
	return formattedData
}

func getPropreties(depDate, DestIATA string) string {
	fmt.Println("Propreties...")

	flightURL := "http://localhost:3000/quoteProperties"
	userData := []byte(`{"duration":7, "departureDate":"` + depDate + `","iata":"` + DestIATA + `"}`)

	// create new http request
	request, error := http.NewRequest("POST", flightURL, bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if error != nil {
		fmt.Println(error)
	}

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
	return formattedData
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
