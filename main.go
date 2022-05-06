package main

import "github.com/klasrak/go-meli-test-dojo/api"

func main() {
	api := api.New()

	if err := api.Run(); err != nil {
		panic(err)
	}
}
