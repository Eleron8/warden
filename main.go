package main

import "github.com/Eleron8/warden/api"

func main() {
	app := api.NewApp("3000")
	app.Run()
}
