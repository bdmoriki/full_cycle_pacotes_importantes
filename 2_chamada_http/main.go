package main

import (
	"io"
	"net/http"
)

func main() {

	client, err := http.Get("https://www.google.com/")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(client.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	client.Body.Close()

}
