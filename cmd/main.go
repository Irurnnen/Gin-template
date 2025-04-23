package main

import (
	"github.com/Irurnnen/gin-template/internal/application"
)

func main() {
	app := application.New()
	app.Run()
}
