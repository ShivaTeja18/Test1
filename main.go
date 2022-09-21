package main

import (
	"EMP_API/Route"

	_ "EMP_API/Details"
	_ "golang.org/x/crypto/openpgp/errors"
)

func main() {
	Route.Connect()
	Route.Initialize()

}
