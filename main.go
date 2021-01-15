// main.go
package main

import "demo/route"

func main () {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}