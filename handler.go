package main

import(
	"bytes"
	"net/http"
	"encoding/json"
	"fmt"
)




func quoteHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/quote" || r.Method != "POST"{
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	flights,err :=callProgrammadeus()
	if err!=nil{
		fmt.Println("Error")
	}
	for _, f := range flights {
		fmt.Printf("departure: %s, destination: %s, \n", f.DepartureCity, f.DestinationCity)
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(flights)
	if err != nil {
		fmt.Println("encoder error")
	} 
	return 
}



func callProgrammadeus() ([]Flight, error) {
	url := "http://localhost:3001/quoteFlights"

	var jsonStr = []byte(`{
		"duration": 7,
		"departureDate": "2023-11-10",
		"destinationIATA": "DXB",
		"departureIATA": "PAR"
	}`)
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

	
	var flights []Flight
	err = json.NewDecoder(response.Body).Decode(&flights)
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func callHotelcodes() ([]Hotel, error) {
	url := "http://localhost:3001/quoteProperties"



	// marshall data to json (like json_encode)
	var jsonStr = []byte(`{
		"duration": 7,
		"departureDate": "2024-01-17T00:00:00Z",
		"iata": "DXB"
	}`)
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

	
	var hotels []Hotel
	err = json.NewDecoder(response.Body).Decode(&hotels)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}