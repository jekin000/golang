package main

import (
	"fmt"
	"net/http"
)

func SearchSingle() {
	resp, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}

	fmt.Println("hello world!")
	resp.Body.Close()
}

func main() {
	SearchSingle()
}
