package requests

type QuoteRequest struct {
	DepartureIATA   string `json:"departureIATA"`
	DestinationIATA string `json:"destinationIATA"`
	DepartureDate   string `json:"departureDate"`
}

type FlightRequest struct {
	DepartureIATA   string `json:"departureIATA"`
	DestinationIATA string `json:"destinationIATA"`
	DepartureDate   string `json:"departureDate"`
	Duration        int    `json:"duration"`
}

type PropertiesRequest struct {
	Duration      int    `json:"duration"`
	DepartureDate string `json:"departureDate"`
	Iata          string `json:"iata"`
}

func CreateFlightReq(qt QuoteRequest) FlightRequest {
	return FlightRequest{
		Duration:        7,
		DepartureDate:   qt.DepartureDate,
		DestinationIATA: qt.DestinationIATA,
		DepartureIATA:   qt.DepartureIATA,
	}
}

func CreatePropertiesReq(qt QuoteRequest) (PropertiesRequest, error) {
	// fmt.Println("= > req : ", qt)
	// dateString := qt.DepartureDate
	// date, err := time.Parse("2024-01-17T00:00:00Z", dateString)
	// fmt.Println("date error : ", err)
	// if err != nil {
	// 	fmt.Errorf("DATE err: ", err.Error())
	// 	return PropertiesRequest{}, err
	// }
	// fmt.Println("created Date : ", date)
	return PropertiesRequest{
		Duration:      7,
		Iata:          qt.DestinationIATA,
		DepartureDate: qt.DepartureDate + "T00:00:00Z",
	}, nil
}
