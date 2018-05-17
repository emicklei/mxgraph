package mxgraph

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestModel2XML(t *testing.T) {
	m := Model{
		Root: []Cell{
			Cell{ID: "0",
				Style:    TextStyle(),
				Geometry: &Geometry{X: "530"}},
		},
	}
	e := xml.NewEncoder(os.Stdout)
	e.Indent("", "\t")
	e.Encode(m)
}
