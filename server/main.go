package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewTaxiServer("../data/beijing.osm.xml")
	log.Fatal(http.ListenAndServe(":9000", server))
}
