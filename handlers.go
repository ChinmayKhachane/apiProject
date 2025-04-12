package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type AppHandler struct {
	db *sql.DB
}

func (h *AppHandler) BasePageResponse(w http.ResponseWriter, r *http.Request) {
	var factdata []condensedFacts
	var err error
	var pricedata []stockPrice
	var articledata []articleInfo
	var wg sync.WaitGroup
	wg.Add(3)
	var pageBody basePageBody

	err = Decode(w, r, &pageBody)
	if err != nil {
		log.Fatal(err)
	}
	newsParam := structToMap(pageBody.NewsParam)
	factParam := pageBody.FilingParam
	fullUrl := formatUrl(newsUrl, newsParam)
	go func() {
		defer wg.Done()
		factdata, err = getFactData(factParam, h.db)
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		pricedata, err = readStockPrices(pageBody.StockParam, h.db)
	}()

	go func() {
		defer wg.Done()
		articledata = newsRequest(fullUrl)
		fmt.Println(fullUrl)
	}()

	wg.Wait()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var baseResponse = basePageResponse{
		Facts:       factdata,
		Articles:    articledata,
		StockPrices: pricedata,
	}
	newerr := Encode(w, r, baseResponse)
	if newerr != nil {
		http.Error(w, newerr.Error(), http.StatusInternalServerError)
	}

}
