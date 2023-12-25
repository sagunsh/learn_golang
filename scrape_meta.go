package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	resp, err := http.Get("https://sagunshrestha.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("status =", resp.Status)
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	list, err := htmlquery.QueryAll(doc, "//meta")
	if err != nil {
		panic(err)
	}

	for _, n := range list {
		content := htmlquery.SelectAttr(n, "content")
		name := htmlquery.SelectAttr(n, "name")
		if name != "" && content != "" {
			fmt.Println("name =", name, ", content =", content)
		}
	}

	title := strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(doc, "//title")))
	fmt.Println("title =", title)
}
