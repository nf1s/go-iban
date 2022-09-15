package main

func main() {
	app := App{}
	app.Initialize(getDBURL(USER, PASSWORD, DB_NAME))
	app.Run()

}
