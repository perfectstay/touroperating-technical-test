# technical test

## Getting started

This repo is our technical test that the future candidates need to pass.

It is currently composed of two fake providers:

- Amadeus which becomes Programmadeus.
- hotelbeds which becomes HotelCodes

## Objective 

The candidate must implement a `POST /quote` endpoint.
The request body will be the following:

```json
{
  "departureIATA": string, // departure city in IATA format (ex. Paris=PAR (check format on web))
  "destinationIATA": string, // destination city in IATA format
  "departureDate": string, // departure date in YYYY-MM-DD format
  "duration": int //  number of nights at the destination
}
```

The endpoint must answer with a list of properties and flight legs:

```json
{
  "departureCity": string,
  "destinationCity": string,
  "items": [
    {
      "propertyPrice": int,
      "flightPrice": int,
      "total": int
    },
    ...
  ]
}
```

:warning: The items array must be sorted by price.

## Models

You can use the thunder client intregrated to visual studio code to see and test the provider's api.

### Programmadeus

Request

```json
{
  "duration": int,
  "departureDate": date, // in YYYY-MM-DD format. ex. "2023-11-12"
  "destinationIATA": string, // destination city in IATA format
  "departureIATA": string // destination city in IATA format
}
```

Response

```json
{
  "flights": [
    {
      "departureAirport": string,
      "arrivalAirport": string,
      "amount": int,
      "pax": int,
      "departureDate": string
    },
    ...
  ]
}
```

or an error

```json
{
  "message": string,
  "code": int
}
```

### HotelCodes

Request

```json
{
  "duration": int,
  "departureDate": date, //  in YYYY-MM-DD format. ex. "2023-11-12"
  "iata": string // destination city in IATA format
}
```

Response

```json
{
  "properties": [
    {
      "id": string,
      "rooms": [
        {
          "id": string,
          "boardCode": string,
          "refundable": boolean
        },
        {
          "id": string,
          "boardCode": string,
          "refundable": boolean,
          "price": float,
          "pax": number
        }
      ]
    }
  ]
}
```

:warning: if a room's price is empty then the room is unavailable.
