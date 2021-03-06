package graph

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

var allowedHighwayTypes = map[string]bool{"motorway": true, "trunk": true, "primary": true, "secondary": true, "tertiary": true, "unclassified": true,
	"residential": true, "living_street": true, "motorway_link": true, "trunk_link": true, "primary_link": true,
	"secondary_link": true, "tertiary_link": true, "service": true}

type states struct {
	active, wayName, wayType string
	Node
	nids []int64
}

func parseXML(g *Graph, r io.Reader) {
	dec := xml.NewDecoder(r)
	var s states

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "node" {
				s.active = "node"

				for _, attr := range tok.Attr {
					switch attr.Name.Local {
					case "id":
						s.id, _ = strconv.ParseInt(attr.Value, 10, 64)
					case "lat":
						s.lat, _ = strconv.ParseFloat(attr.Value, 64)
					case "lon":
						s.lon, _ = strconv.ParseFloat(attr.Value, 64)
					}
				}
			} else if s.active == "node" && tok.Name.Local == "tag" && tok.Attr[0].Value == "name" {
				s.name = tok.Attr[1].Value
			} else if tok.Name.Local == "way" {
				s.active = "way"
			} else if s.active == "way" && tok.Name.Local == "nd" {
				nid, _ := strconv.ParseInt(tok.Attr[0].Value, 10, 64)
				s.nids = append(s.nids, nid)
			} else if s.active == "way" && tok.Name.Local == "tag" {
				if tok.Attr[0].Value == "name" {
					s.wayName = tok.Attr[1].Value
				} else if tok.Attr[0].Value == "highway" {
					s.wayType = tok.Attr[1].Value
				}
			}
		case xml.EndElement:
			if tok.Name.Local == "node" {
				g.AddNode(&Node{s.id, s.lat, s.lon, s.name})
				s = states{}
			} else if tok.Name.Local == "way" {
				if _, ok := allowedHighwayTypes[s.wayType]; ok {
					for i := 0; i < len(s.nids)-1; i++ {
						g.AddEdge(s.nids[i], s.nids[i+1], s.wayName)
						g.AddEdge(s.nids[i+1], s.nids[i], s.wayName)
					}
				}
				s = states{}
			}
		}
	}
}
