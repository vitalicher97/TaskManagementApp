package main

import (
	"./controller"
	"./handleDb"
)

func main() {
	handleDb.CreateDbObject()

	controller.MainController()

}
