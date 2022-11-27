package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	servidorOrigemURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("URL invalida do servidor de origem")
	}

	proxyReverso := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[proxy Reverso] recebido a requisição a: %s\n", time.Now())

		req.Host = servidorOrigemURL.Host
		req.URL.Host = servidorOrigemURL.Host
		req.URL.Scheme = servidorOrigemURL.Scheme
		req.RequestURI = ""

		respostaServidorOrigem, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, respostaServidorOrigem.Body)

	})

	log.Fatal(http.ListenAndServe(":8080", proxyReverso))
}
