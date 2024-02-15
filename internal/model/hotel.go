package model

import "time"

type HotelRequest struct {
	Duration int `json:"duration"`
	DepartureDate time.Time `json:"departureDate"`
	IATA string `json:"iata"`
}

type HotelResponse struct {
	Properties []HotelProperty `json:"properties"`
}

type HotelRoom struct {
	ID string `json:"id"`
	BoardCode string `json:"boardCode"`
	Refundable bool `json:"refundable"`
	Price float32 `json:"price"`
	Pax int `json:"pax"`
}

type HotelProperty struct {
	ID string `json:"id"`
	Rooms []HotelRoom `json:"rooms"`
}