package main

import (
	"fmt"
	"log"
)

func printStockPrices(stockPrices []stockPrice) {
	for _, stockP := range stockPrices {
		fmt.Printf("%v \n", stockP)
	}
}

func main() {

	db, err := CreateConn()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryparam := QueryParam{Ticker: "AAPL"}

	prices, err := readStockPrices(queryparam, db)
	if err != nil {
		log.Fatal(err)
	}
	printStockPrices(prices)

}

// login
