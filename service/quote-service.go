package service

import (
	"fmt"
	"net/http"
)

func GetQuote() {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}
