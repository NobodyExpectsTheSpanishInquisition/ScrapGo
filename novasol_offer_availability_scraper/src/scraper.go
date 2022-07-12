package src

import (
	"encoding/json"
	"io"
	"net/http"
)

func scrap(url string) (offerAvailabilities, error) {
	var err error
	var data offerAvailabilities
	response, err := http.Get(url)
	bodyBytes, err := io.ReadAll(response.Body)

	err = json.Unmarshal(bodyBytes, &data)

	return data, err
}
