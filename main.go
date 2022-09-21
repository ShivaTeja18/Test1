package main

import (
	"EMP_API/Route"
	_ "golang.org/x/crypto/openpgp/errors"
)

func main() {
	Route.Connect()
	Route.Initialize()

}
