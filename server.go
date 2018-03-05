package main

import "github.com/ken-aio/go-echo-sqlboiler/src/route"

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":1313"))
}
