package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(addr string, handler http.Handler)
}
