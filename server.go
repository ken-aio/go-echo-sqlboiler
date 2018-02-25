package main

import "github.com/ken-aio/go-echo-sqlboiler/route"

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":1313"))
}
