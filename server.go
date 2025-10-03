package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	ser := &http.Server{
		Addr:":8080",
		Handler: mux,
	}
	if err := ser.ListenAndServe(); err != nil {
		ser.Close()
	}
	
}

