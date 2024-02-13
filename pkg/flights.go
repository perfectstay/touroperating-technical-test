package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"technical-test/models"
)

type FlightQuotesResponse struct {
	Flights []models.FlightQuote
}

func CallProgrammadeus(quoteInfos models.FightJson) ([]models.FlightQuote, error) {
	url := "http://localhost:3001/quoteFlights"

	var jsonStr, err = json.Marshal(quoteInfos)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	var flightsQuotes FlightQuotesResponse
	err = json.NewDecoder(response.Body).Decode(&flightsQuotes)
	if err != nil {
		return nil, err
	}
	return flightsQuotes.Flights, nil
}
