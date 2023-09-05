package main

import (
	"fmt"
	"io"
	"net/http"
)

func getMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /metrics request\n")
	outputString := query.getData()
	io.WriteString(w, outputString)
}

func main() {
	http.HandleFunc("/metrics", getMetrics)

	http.ListenAndServe(":2112", nil)
}
