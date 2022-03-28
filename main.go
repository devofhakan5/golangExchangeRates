package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.floatrates.com/daily/usd.xml")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var i channel
	xml.Unmarshal(data, &i)
	fmt.Println(i.Items)
}

type items struct {
	XMLName xml.Name `xml:"item"`
	DESC    string   `xml:"description"`
	TITLE   string   `xml:"title"`
}
type channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []items  `xml:"item"`
}

func (x items) String() string {
	return fmt.Sprintf("\t Description:%s --Title:%s \n", x.DESC, x.TITLE)
}
