package main

import (
	"fmt"
	"log"
)

func printStockPrices(stockPrices []stockPrice) {
	for _, stockP := range stockPrices {
		fmt.Printf("Ticker: %v, Date: %v, Volume: %v, Open: %v, High: %v, Low: %v, Transactions: %v\n",
			stockP.Ticker, stockP.Date, stockP.Volume, stockP.Open, stockP.High, stockP.Low, stockP.Transactions)
	}
}

func printCompany(comp company) {
	fmt.Printf("Ticker: %s, CIK: %d, Title: %s", comp.Ticker, comp.cik, comp.Title)
}

func main() {

	db, err := CreateConn()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryparam := QueryParamDate{}
	candidateCom := company{Ticker: "AAPL"}

	prices, err := readStockPrices(queryparam, db)
	if err == nil {
		printStockPrices(prices)
	} else {
		fmt.Printf(err.Error(), "\n")
	}

	com, err := readCompany(candidateCom, db)
	if err != nil {
		log.Fatal(err)
	}
	printCompany(com)

}

// login
