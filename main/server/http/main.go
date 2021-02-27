package main

func main() {
	iris := newIrisApp()
	iris.Listen(":8080")
}
