package src

import (
	"encoding/json"
	"net/http"
)

func scrapNovasolOfferAvailability(w http.ResponseWriter, _ *http.Request) {
	response, err := handle(newCommand(getCallendarOfferUrl(), getDateFromToCheck(), getDateToToCheck()))
	if err != nil {
		panic(err.Error())
	}

	responseJson, err := json.Marshal(&response)
	if nil != err {
		panic(err.Error())
	}

	_, err = w.Write(responseJson)
	if err != nil {
		panic(err.Error())
	}
}
