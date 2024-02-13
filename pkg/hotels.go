package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"technical-test/models"
)

type HotelsQuoteResponse struct {
	Properties []models.HotelQuote
}

func CallHotelcodes(hotelFinder models.HotelFinderInfos) ([]models.HotelQuote, error) {
	url := "http://localhost:3000/quoteProperties"

	// marshall data to json (like json_encode)
	var jsonStr ,err = json.Marshal(hotelFinder)
	req, err := http.NewRequest("POST", url,  bytes.NewReader(jsonStr))
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

	
	var hotelsRes HotelsQuoteResponse
	err = json.NewDecoder(response.Body).Decode(&hotelsRes)
	if err != nil {
		return nil, err
	}
	return hotelsRes.Properties, nil
}