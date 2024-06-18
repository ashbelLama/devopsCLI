package service

import (
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func GetJoke(baseAPI string) ([]byte, error) {
	if baseAPI == "" {
		return nil, errors.New("invalid or broken baseAPI")
	}
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		return nil, errors.New("Could not make a request. Get " + baseAPI + ": http: no Host in request URL")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/example/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}
	return responseBytes, nil
}
