package main

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
)

type property struct {
	Address     string `json:"address"`
	Price       string `json:"price"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func main() {
	// rentals in london
	url := "https://www.rightmove.co.uk/property-to-rent/London.html"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
	//cookie := &http.Cookie{Name: "_abck", Value: ""}
	//req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response status :", resp.Status)

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	rentals, err := htmlquery.QueryAll(doc, "//div[contains(@class, 'l-searchResult')]")
	if err != nil {
		panic(err)
	}

	for _, rental := range rentals {
		var item property

		item.Address = htmlquery.SelectAttr(htmlquery.FindOne(rental, "//meta[@itemprop='streetAddress']"), "content")
		item.Price = htmlquery.InnerText(htmlquery.FindOne(rental, "//span[@class='propertyCard-priceValue']"))
		item.Url = htmlquery.SelectAttr(htmlquery.FindOne(rental, "//a[contains(@class, 'propertyCard-link')]"), "href")
		item.Description = htmlquery.InnerText(htmlquery.FindOne(rental, "//span[@itemprop='description']"))

		data, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
