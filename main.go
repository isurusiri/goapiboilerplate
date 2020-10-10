package main

import "fmt"

var app App

func main() {
	fmt.Println("Welcome to Go API Boilerplate.")
	app.Initialize()
	app.Run(":8080")
}
