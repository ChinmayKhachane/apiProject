package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
)

type stockPrice struct {
	Ticker       string
	Date         string
	Volume       int64
	Open         float64
	Close        float64
	High         float64
	Low          float64
	Transactions int64
}

type QueryParam struct {
	Ticker     string
	End_date   string
	Start_date string
}

var server = "sec-filings-server.database.windows.net"
var port = 1433
var user = "SEC_admin"
var password = "FremontFreaks!"
var database = "SEC_filings"

func CreateConn() (*sql.DB, error) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	newerr := db.Ping()
	if newerr != nil {
		return nil, newerr

	} else {
		return db, nil

	}

}

func readStockPrices(queryparam QueryParam, db *sql.DB) ([]stockPrice, error) {

	ctx := context.Background()
	err := db.PingContext(ctx)
	var tsql string
	if err != nil {
		return nil, err

	}
	if queryparam.Ticker != "" && queryparam.End_date != "" && queryparam.Start_date != "" {
		tsql = fmt.Sprintf(`SELECT TOP 10 ticker, date, volume, 
    [open], [close], high, low, transactions FROM stock_prices WHERE ticker = '%s' 
    AND date BETWEEN '%s' AND '%s' ORDER BY date DESC`, queryparam.Ticker, queryparam.Start_date, queryparam.End_date)
	} else if queryparam.Ticker != "" {
		tsql = fmt.Sprintf(`SELECT TOP 10 ticker, date, volume, 
    [open], [close], high, low, transactions FROM stock_prices WHERE ticker= '%s' ORDER BY date DESC`, queryparam.Ticker)
	} else {
		nodata := errors.New("No ticker or start date or end date")
		return nil, nodata
	}

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stockPrices := []stockPrice{}

	for rows.Next() {
		var candidate stockPrice

		err := rows.Scan(&candidate.Ticker, &candidate.Date, &candidate.Volume,
			&candidate.Open, &candidate.Close, &candidate.High, &candidate.Low, &candidate.Transactions)
		if err != nil {
			return nil, err
		}
		stockPrices = append(stockPrices, candidate)
	}
	return stockPrices, nil

}
