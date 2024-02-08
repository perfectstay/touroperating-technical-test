package main

import ()


type flightJson struct {
	Duration int `json:"duration"`
	DepartureDate string`json:"departureDate"`
	DestinationIATA string `json:"destinationIATA"`
	DepartureIATA string `json:"departureIATA"`

}

type Item struct{
	PropertyPrice int
	FlightPrice int
	Total int
} 

type Room struct {
	ID string
	BoardCode string
	Refundable bool
	Price int
	Pax string
}

type Hotel struct {
	ID string
	Rooms []Room
}

type Flight struct {
	DepartureCity string
	DestinationCity string
	Items []Item

}  