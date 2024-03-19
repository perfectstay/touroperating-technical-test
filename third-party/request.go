package thirdparty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rcTest/requests"
	"rcTest/responses"
)

var FLIGHTS_URL = "http://localhost:3001/quoteFlights"

var PROPERTIES_URL = "http://localhost:3000/quoteProperties"

func RequestFlights(reqData requests.FlightRequest) ([]responses.FlightData, error) {
	bjson, err := json.Marshal(reqData)
	if err != nil {
		fmt.Printf("=> json marshall erro : ", err.Error())
		return []responses.FlightData{}, err
	}
	// fmt.Println("json : >", string(bjson))

	client := &http.Client{}
	req, err := http.NewRequest("POST", FLIGHTS_URL, bytes.NewBuffer(bjson))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return []responses.FlightData{}, err
	}
	// req.Header.Add("Authorization", "Bearer <token>")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return []responses.FlightData{}, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return []responses.FlightData{}, err
	}
	// fmt.Println(string(body))
	flightData, err := responses.ExtractFlightData(body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return []responses.FlightData{}, err
	}

	return flightData, nil
}

func RequestProperties(reqData requests.PropertiesRequest) ([]responses.Property, error){
	bjson, err := json.Marshal(reqData)
	if err != nil {
		fmt.Printf("=> json marshall erro : ", err.Error())
		return []responses.Property{}, err
	}
	// fmt.Println("json : >", string(bjson))

	client := &http.Client{}
	req, err := http.NewRequest("POST", PROPERTIES_URL, bytes.NewBuffer(bjson))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return []responses.Property{}, err
	}
	// req.Header.Add("Authorization", "Bearer <token>")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return []responses.Property{}, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return []responses.Property{}, err
	}
	// fmt.Println(string(body))
	propertyData, err := responses.ExtractPropertiesData(body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return []responses.Property{}, err
	}

	return propertyData, nil
}
