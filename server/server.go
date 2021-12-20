package main

import (
	"encoding/json"
	"etaxi/router"
	"etaxi/streetmap"
	"etaxi/streetmap/graph"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TaxiServer struct {
	*streetmap.StreetMap
	http.Handler
}

type Location struct {
	ID   int64   `json:"id"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}

type Direction struct {
	Nodes [][2]float64 `json:"nodes"`
	Text  string       `json:"text"`
}

func NewLocationFromNode(n *graph.Node) *Location {
	return NewLocation(n.ID(), n.Lat(), n.Lon(), n.Name())
}

func NewLocation(id int64, lat, lon float64, name string) *Location {
	return &Location{ID: id, Lat: lat, Lon: lon, Name: name}
}

func NewTaxiServer(fn string) *TaxiServer {
	t := new(TaxiServer)
	t.StreetMap = streetmap.NewStreetMap(fn)

	router := http.NewServeMux()
	router.Handle("/locations/", http.HandlerFunc(t.locationsHandler))
	router.Handle("/direction/", http.HandlerFunc(t.directionHandler))

	t.Handler = router
	return t
}

// GET /locations?name=xxx return every location ids of which name has prefix xxx
// GET /locations/id return node info of specific id
func (t *TaxiServer) locationsHandler(w http.ResponseWriter, r *http.Request) {
	s := strings.TrimPrefix(r.URL.Path, "/locations/")
	if s == "" {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		if v, ok := m["name"]; ok {
			var res []struct {
				ID   int64  `json:"id"`
				Name string `json:"name"`
			}
			locNames := t.GetLocationsByPrefix(v[0])
			for _, ln := range locNames {
				ids := t.Get(ln)
				for _, id := range ids {
					res = append(res, struct {
						ID   int64  `json:"id"`
						Name string `json:"name"`
					}{ID: id, Name: ln})
				}
			}
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		n := t.GetNode(id)
		if n != nil {
			loc := NewLocationFromNode(n)
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(&loc)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// GET /dir?from=xxx&to=yyy return navigation from xxx to yyy
func (t *TaxiServer) directionHandler(w http.ResponseWriter, r *http.Request) {
	s := strings.TrimPrefix(r.URL.Path, "/direction/")
	if s == "" {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		s1, ok1 := m["slat"]
		s2, ok2 := m["slon"]
		s3, ok3 := m["dlat"]
		s4, ok4 := m["slon"]

		if ok1 && ok2 && ok3 && ok4 {
			slat, _ := strconv.ParseFloat(s1[0], 64)
			slon, _ := strconv.ParseFloat(s2[0], 64)
			dlat, _ := strconv.ParseFloat(s3[0], 64)
			dlon, _ := strconv.ParseFloat(s4[0], 64)

			dir := Direction{}
			shortestPath := router.Navigate(t.StreetMap, slat, slon, dlat, dlon)

			for e := shortestPath.Front(); e != nil; e = e.Next() {
				node := t.GetNode(e.Value.(int64))
				dir.Nodes = append(dir.Nodes, [2]float64{node.Lat(), node.Lon()})
			}

			dir.Text = router.GetDirectionsText(t.StreetMap, shortestPath)

			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(&dir)
			w.WriteHeader(http.StatusOK)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
