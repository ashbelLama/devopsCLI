package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ashbelLama/cli-golang/entity"
	"github.com/ashbelLama/cli-golang/service"
)

func JokeController() {
	const jokeURL string = "https:/icanhazdadjoke.com/"
	responseBytes, err := service.GetJoke(jokeURL)
	if err != nil {
		log.Fatal(err)
	}
	joke := entity.Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(string(joke.Joke))
}
