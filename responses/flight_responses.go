package responses

import (
	"encoding/json"
	"fmt"
)

type FlightResponse struct {
	Flights []FlightData `json:"flights"`
}

type FlightData struct {
	DepartureAirport string `json:"departureAirport"`
	ArrivalAirport   string `json:"arrivalAiport"`
	// int /float ????
	Amount        int    `json:"amount"`
	Pax           int    `json:"pax"`
	DepartureDate string `json:"departureDate"`
}

func ExtractFlightData(bFlightData []byte) ([]FlightData, error) {
	var flightReponse FlightResponse
	err := json.Unmarshal(bFlightData, &flightReponse)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR : ", err)
		return []FlightData{}, err
	}
	return flightReponse.Flights, nil
}

// ======Properties========

type PropertiesResponse struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Id    string     `json:"id"`
	Rooms []RoomData `json:"rooms"`
}

type RoomData struct {
	Id         string `json:"id"`
	BoardCode  string `json:"boardCode"`
	Refundable bool   `json:"refundable"`
}

func ExtractPropertiesData(bPropertiesData []byte) ([]Property, error) {
	var propertyResponse PropertiesResponse
	err := json.Unmarshal(bPropertiesData, &propertyResponse)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR : ", err)
		return []Property{}, err
	}
	return propertyResponse.Properties, nil
}
