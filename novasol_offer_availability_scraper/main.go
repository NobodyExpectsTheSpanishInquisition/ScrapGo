package main

import (
	"net/http"
	"novasol_offer_availability_scraper/src"
)

func main() {
	var err error
	err = src.LoadEnv()
	if err != nil {
		panic(err.Error())
	}

	r := src.CreateRouter()

	src.RegisterRoutes(r)

	err = http.ListenAndServe(":8000", r)
	if nil != err {
		panic(err.Error())
	}
}
