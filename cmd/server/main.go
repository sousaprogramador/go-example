package main

import (
	"log"

	"github.com/sousaprogramador/session-finance-go-lang/adapter/http"
)

func main() {
	log.Fatal(http.Init())
}
