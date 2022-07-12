package src

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

func RegisterRoutes(r *mux.Router) {
	scrapOfferAvailability(r)
}

func scrapOfferAvailability(r *mux.Router) {
	r.HandleFunc("/api/scrap/novasol/offer/available", scrapNovasolOfferAvailability)
}
