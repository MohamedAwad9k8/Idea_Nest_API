package main

import (
	"test/pkg"
)

func main() {
	app := pkg.NewApp()
	app.Run("127.0.0.1:8080")
}
