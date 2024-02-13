package models

import "time"

type FightJson struct {
	Duration        int    `json:"duration"`
	DepartureDate   time.Time `json:"departureDate"`
	DestinationIATA string `json:"destinationIATA"`
	DepartureIATA   string `json:"departureIATA"`
}

type FullQuote struct {
	DepartureCity   string
	DestinationCity string
	Items           []Price
}

type Price struct {
	PropertyPrice int
	FlightPrice   int
	Total         int
}

type HotelFinderInfos struct {
	Duration      int
	DepartureDate time.Time
	Iata          string
}

type HotelQuote struct {
	Id    string
	Rooms []Room
}

type Room struct {
	Id         string
	BoardCode  string
	Refundable bool
	Price      int
	Pax        string
}

type FlightQuote struct {
	DepartureAirport string
	ArrivalAirport   string
	Amount           int
	Pax              int
	DepartureDate    string
}
