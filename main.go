package main

import (
	"EMP_API/Route"
	"EMP_API/dbc"
)

func main() {
	dbc.Connect()
	Route.Initialize()

}
