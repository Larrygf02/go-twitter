package main

import (
	"fmt"
	"log"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/handler"
)

func main() {
	fmt.Println("Hola API Twitter")
	if bd.CheckConnect() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handler.HandlerRouters()
}
