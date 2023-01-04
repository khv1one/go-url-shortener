package main

func main() {
	app := InitializeApp()

	app.Logger.Fatal(app.Start(":8888"))
}
