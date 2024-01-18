package server

type dataQuote struct {
	DepartureIATA   string `json:"departureIATA"`   // departure city in IATA format (ex. Paris=PAR (check format on web))
	DestinationIATA string `json:"destinationIATA"` // destination city in IATA format
	DepartureDate   string `json:"departureDate"`   // departure date in YYYY-MM-DD format
}
type flight struct {
	DepartureAirport string `json:"departureAirport"`
	ArrivalAiport    string `json:"arrivalAiport"`
	Amount           string `json:"amount"`
	Pax              string `json:"pax"`
	DepartureDate    string `json:"departureDate"`
}

type properties struct {
	Duration      string `json:"duration"`
	DepartureDate string `json:"departureDate"`
	Iata          string `json:"iata"`
}
