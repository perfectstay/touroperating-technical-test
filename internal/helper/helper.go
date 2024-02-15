package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var (
	HOTEL_SERVICE = "http://localhost:3000"
	AMADEUS_SERVICE = "http://localhost:3001"
)

func MakeHTTPCall(method string, url string, body interface{}) (*http.Response, error){
	var reqData *bytes.Buffer
	
	client := &http.Client{
		Timeout: time.Millisecond * 500,
	}

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		reqData = bytes.NewBuffer(jsonData)
	}
	
	req, err := http.NewRequest(method, url, reqData)

	if err != nil {
		panic(err)
	}

	return client.Do(req)
}
