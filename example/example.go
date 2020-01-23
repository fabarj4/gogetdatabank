package main

import (
	"fmt"
	"net/http"

	"github.com/fabarj4/gogetdatabank"
)

func main() {
	http.HandleFunc("/", gogetdatabank.DataBankHandler)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
