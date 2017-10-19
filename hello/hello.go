package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func main() {
	resp, err := http.Get("https://www.goinggo.net/index.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var d Document
	err = xml.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		fmt.Println(err)
		return
	}

	term := "semantic"

	for i := range d.Channel.Items {
		if !strings.Contains(d.Channel.Items[i].Description, term) {
			continue
		}

		link := fmt.Sprintf("http%s", d.Channel.Items[i].Link)

		fmt.Printf("%s\n%s\n\n", d.Channel.Items[i].Title, link)
	}
}
