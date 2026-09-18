package main

import (
	"kautsarhasby/ewallet-transaction/cmd"
	"kautsarhasby/ewallet-transaction/helpers"
)

func main() {

	helpers.SetupEnv()
	helpers.SetupLogger()
	helpers.SetupDatabase()

	// go cmd.ServeGRPC()
	cmd.ServeHTTP()

}
