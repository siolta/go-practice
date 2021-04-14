// go has built in xml and xml-like format support

package main
import (
	"encoding/xml"
	"fmt"
)

// plant will map to XML, and uses field tags to direct the en/decode
type Plant struct {
	XMLName xml.Name 	`xml:"plant"`
	Id 		int 	 	`xml:"id,attr"`
	Name	string		`xml:"name"`
	Origin	[]string	`xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v,",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethopia", "Brazil"}

	// emit XML of our plant, using MarshalIndent for more readable output
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// add a header by appending it explicitly
	fmt.Println(xml.Header + string(out))

	// use unmarshal to parse a stream of bytes with XML into a data structure
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string {"Mexico", "California"}

	// parent>child>plant tells encoder to nest all plants under <parent><child>
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
