package model

import "time"

type AmadeusRequest struct {
	Duration int `json:"duration"`
	DeparatureDate time.Time `json:"departureDate"`
	DestinationIATA string `json:"destinationIATA"`
	DeparatureIATA string `json:"departureIATA"`
}

type AmadeusFlight struct {
	DeparatureAirport string `json:"departureAirport"`
	ArrivalAirport string `json:"arrivalAirport"`
	Amont int `json:"amont"`
	Pax int `json:"pax"`
	DeparatureDate string `json:"deparatureDate"`
}

type AmadeusResponse struct {
	Flights []AmadeusFlight `json:"flights"`
}

type AmadeuseErrorResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
}