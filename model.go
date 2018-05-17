package mxgraph

import "encoding/xml"

// Model is a mxgraph element.
type Model struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Dx      string   `xml:"dx,attr"`
	Dy      string   `xml:"dy,attr"`
	Arrows  string   `xml:"arrows,attr,omitempty"`
	Root    []Cell   `xml:"root"`
}

// Cell is a mxgraph element.
type Cell struct {
	XMLName  xml.Name `xml:"mxCell"`
	ID       string   `xml:"id,attr"`
	Style    Style    `xml:"style,attr"`
	ParentID string   `xml:"parent,attr"`
	Vertex   string   `xml:"vertex,attr"`
	Geometry *Geometry
}

// Geometry is a mxgraph element.
type Geometry struct {
	XMLName xml.Name `xml:"mxGeometry"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	As      string   `xml:"as,attr"`
}

// Style models the attribute value of a Cell.
// text;html=1;strokeColor=none;fillColor=none;align=center;verticalAlign=middle;whiteSpace=wrap;rounded=0;dashed=1;
type Style struct {
	Attributes map[string]string
}

// TextStyle return a new Style with text.
func TextStyle() Style {
	return Style{Attributes: map[string]string{"text": ""}}
}

// MarshalXMLAttr implements xml.MarshalXMLAttr
func (a Style) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: xml.Name{Local: "style"}, Value: "text"}, nil
}
