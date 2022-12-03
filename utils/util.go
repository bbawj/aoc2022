package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetInput(day int) string {
	client := http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	if err != nil {
		log.Fatal(err)
	}
	cookie := http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5f3d518a628771ef8b579c04ce9a60374dab99a23467c9d91d5772ba581426b423b0821976c760009d467b31ebe344217d410b25c5df36192d",
	}
	request.AddCookie(&cookie)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}
