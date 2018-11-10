package main

import "github.com/Asprilla24/vermouth/api"

func main() {
	api := api.New()

	api.Run("8080")
}
