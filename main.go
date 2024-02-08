package main

import (
	"fmt"
	"log"
	"net/http"
)
  



func main(){ 

	http.HandleFunc("/quote", quoteHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	//r := gin.Default()

	//r.GET("/", handlerQuote)
  
	//r.Run()
}