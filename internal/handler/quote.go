package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"technical-test/internal/helper"
	"technical-test/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type QuoteHandler struct{}

func NewQuoteHandler() *QuoteHandler {
	return &QuoteHandler{}
}

func getHotels(hotelReq *model.HotelRequest) (*model.HotelResponse, error) {
	res, err := helper.MakeHTTPCall("POST", helper.HOTEL_SERVICE+"/quoteProperties", hotelReq)
	if err != nil {
		return nil, errors.New("unable to make http request")
	}

	fmt.Printf("%+v\n", res.Body)

	defer res.Body.Close()
	var hotelResponse *model.HotelResponse
	json.NewDecoder(res.Body).Decode(&hotelResponse)

	return hotelResponse, nil
}

func getFlights(flightReq *model.AmadeusRequest) (*model.AmadeusResponse, error) {
	res, err := helper.MakeHTTPCall("POST", helper.AMADEUS_SERVICE+"/quote", flightReq)
	if err != nil {
		return nil, errors.New("unable to make http request")
	}

	fmt.Printf("%+v\n", res.Body)

	defer res.Body.Close()
	var flightResp *model.AmadeusResponse
	json.NewDecoder(res.Body).Decode(&flightResp)

	return flightResp, nil
}

func (h *QuoteHandler) HandlePostQuote(c *fiber.Ctx) error {
	var quote model.QuoteRequest

	err := c.BodyParser(&quote)

	if err != nil {
		return c.SendStatus(400)
	}

	customDeparatureFormat, err := time.Parse("2006-01-01", quote.DeparatureDate)

	if err != nil {
		return c.SendStatus(400)
	}

	// replace by fetchData() later
	hotelReq := &model.HotelRequest{
		Duration:      quote.Duration,
		IATA:          quote.DestinationIATA,
		DepartureDate: customDeparatureFormat,
	}
	flightReq := &model.AmadeusRequest{
		Duration:        quote.Duration,
		DeparatureDate:  customDeparatureFormat,
		DestinationIATA: quote.DestinationIATA,
		DeparatureIATA:  quote.DeparatureIATA,
	}

	quoteResponse := &model.QuoteResponse{
		DeparatureCity: quote.DeparatureIATA,
		DestinationCity: quote.DestinationIATA,
	}

	items, err := fetchResults(flightReq, hotelReq)

	if err != nil {
		return c.SendStatus(500)
	}

	quoteResponse.Items = sortByPrice(items)

	return c.SendStatus(200)

}

func fetchResults(flights *model.AmadeusRequest, hotels *model.HotelRequest) ([]model.QuoteItem, error) {
	// use channel here for hotel & flight http call
	hotels, err := getHotels(hotelReq)

	if err != nil {
		return errors.New("")
	}

	flights, err := getFlights(flightReq)

	if err != nil {
		return errors.New("")
	}
	return nil
}

func sortByPrice(items []model.QuoteItem) []model.QuoteItem {
	// implemant bubble sort alg
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j].Total > items[j+1].Total {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	return items
}
