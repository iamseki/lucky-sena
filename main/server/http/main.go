package main

import irisapp "lucky-sena/main/server/iris-app"

func main() {
	iris := irisapp.NewIrisApp()
	iris.Listen(":8080")
}
