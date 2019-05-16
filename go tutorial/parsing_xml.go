package main

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	// "os"
)

var XML = []byte(`
<?xml version="1.0" encoding="UTF-8" ?>
<doing>
	<heading>Reminder</heading>
	<Description>thisst</Description>
</doing>
<doing>
	<heading>s</heading>
	<Description>this is a test</Description>
</doing>
<doing>
	<heading>a</heading>
	<Description>this is a test</Description>
</doing>
`)

type SitemapIndex struct {
	Locations []Location `xml:"doing"`
}
type Location struct {
	Loc string `xml:"heading"`
	Des string `xml:"Description"`
}

func (s Location) String() string {
	return fmt.Sprintf("%s - %s", s.Loc, s.Des)
}
func main() {
	// resp, _ := os.Open("note.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// resp.Close()
	var s SitemapIndex
	xml.Unmarshal(XML, &s.Locations)
	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Printf("\n %s", Location.Loc)
	}
}
