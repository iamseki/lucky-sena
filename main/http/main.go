package main

import irisapp "lucky-sena/main/http/iris-app"

func main() {
	app := irisapp.NewIrisApp()
	app.Listen(":8080")
}
