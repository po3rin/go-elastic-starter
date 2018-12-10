package main // import "api_server"

import (
	"fmt"
	"net/http"

	"api_server/logger"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))
	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Elastic Stack"
	logger.Info(msg)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, msg)
}
