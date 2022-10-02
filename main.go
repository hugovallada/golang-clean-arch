package main

import (
	"log"
	"net/http"

	"github.com/hugovallada/go-clean-arch/client"
)

func main() {
	err := http.ListenAndServe(":9000", client.Routes())

	if err != nil {
		log.Fatal("Não foi possível iniciar o servidor")
	}
}
