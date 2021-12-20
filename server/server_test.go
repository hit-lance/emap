package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETLocationsByPrefix(t *testing.T) {
	server := NewTaxiServer("../data/berkeley.osm.xml")

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/locations/1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		fmt.Println(response.Result().Header.Get("content-type"))
		fmt.Println(response.Body)

		// got := response.Body.String()
		// want := "20"

		// if got != want {
		// 	t.Errorf("got %q, want %q", got, want)
		// }
	})
}
