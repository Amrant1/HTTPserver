package main

import (
	"fmt"
	"errors"
	"net/http"
)


type Server struct {
	Addr  string// Address to listen on
}


func ListenAndServe(s *Server) error {
	err := errors.New("404 Not Found")
	return err
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "404 Not Found")
	})
	http.ListenAndServe(":8080", mux)
}

