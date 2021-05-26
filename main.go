package main

import (
	"net/http"

	"br.com.industrial/loja/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8080", nil)
}
