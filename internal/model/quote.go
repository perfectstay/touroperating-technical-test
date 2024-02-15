package model

type QuoteRequest struct {
	DeparatureIATA string `json:"departureIATA"`
	DestinationIATA string `json:"destinationIATA"`
	DeparatureDate string `json:"departureDate"`
	Duration int `json:"duration"`
}

type QuoteItem struct {
	PropertyPrice int `json:"propertyPrice"`
	FlightPrice int `json:"flightPrice"`
	Total int `json:"total"`
}

type QuoteResponse struct {
	DeparatureCity string `json:"departureCity"`
	DestinationCity string `json:"destinationCity"`
	Items []QuoteItem
}