package main

import (
  "net/http"
  "github.com/BurntSushi/toml"
  "github.com/gorilla/mux"
  "io/ioutil"
  "fmt"
  "strconv"
)

type Config struct {
  ListenHost string
  ListenPort int
  RAHost string
  RAPort int
  Username string
  Password string
}

func main() {
  c := getConfig()
  r := mux.NewRouter()
  r.HandleFunc("/", Index).Methods("GET")
  r.HandleFunc("/command", Command).Methods("POST")
  http.Handle("/", r)
  fmt.Println("Listening on " + c.ListenHost + ":" + strconv.Itoa(c.ListenPort))
  http.ListenAndServe(c.ListenHost + ":" + strconv.Itoa(c.ListenPort), nil)
}

func getConfig() (Config) {
  var c Config
  configfile, _ := ioutil.ReadFile("config.toml")
  if _, err := toml.Decode(string(configfile), &c); err != nil {
    fmt.Println("Error reading config file.")
  }
  return c
}
