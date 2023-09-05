package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/shih-chris/grafana_data_eng_demo/pkg/query"
)

func getMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /metrics request\n")
	dsTypeCounts := query.GetData()
	var outputString string
	for _, v := range dsTypeCounts {
		outputString = outputString + fmt.Sprintf("Data Source Type: %s; Count: %d\n", v.DsType, v.DsCount)
	}
	// queryResult = resultString +
	io.WriteString(w, outputString)
}

func main() {
	http.HandleFunc("/metrics", getMetrics)

	http.ListenAndServe(":2112", nil)
}
