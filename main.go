package main

import "os"

func main() {
	var USER = os.Getenv("DB_USER")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")

	app := App{}
	app.Initialize(getDBURL(USER, PASSWORD, DB_NAME))
	app.Run()

}
