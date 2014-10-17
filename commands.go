package main

import "net/http"
import "fmt"

func Command(o http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(o, "Command handler reached")
}
