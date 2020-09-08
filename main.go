package main

import (
	"net/http"

	"./routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8001", nil)

}
