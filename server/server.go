package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	servidorOrigemHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[Servidor de Origem] recebida requisicao em: %s\n", time.Now())
		_, _ = fmt.Fprintf(rw, "resposta do servidor de origem")
	})

	log.Fatal(http.ListenAndServe(":8081", servidorOrigemHandler))
}
