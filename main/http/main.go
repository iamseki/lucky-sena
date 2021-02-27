package main

import irisapp "lucky-sena/main/http/iris-app"

func main() {
	iris := irisapp.NewIrisApp()
	iris.Listen(":8080")
}
