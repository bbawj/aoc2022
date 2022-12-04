package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetInput(day int) {
	if _, err := os.Stat(fmt.Sprintf("/home/bawj/projects/aoc2022/day%d/input.txt", day)); errors.Is(err, os.ErrNotExist) {

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
		f, err := os.Create(fmt.Sprintf("/home/bawj/projects/aoc2022/day%d/input.txt", day))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(body); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Input already downloaded")
	}
}
