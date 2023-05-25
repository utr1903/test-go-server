package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Exposes key-value pairs which are separeted by a semicolon
func kvp(w http.ResponseWriter, req *http.Request) {

	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"

	l := ""
	for k, v := range m {

		l += k + ":" + v + "\n"
	}
	fmt.Fprintf(w, "%v", l)
}

// Returns the provided HTTP headers
func headers(w http.ResponseWriter, req *http.Request) {

	bytes, err := json.Marshal(req.Header)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/kvp", kvp)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}
