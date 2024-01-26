package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(error error) {
	if error != nil {
		fmt.Println(error)
	}
}

func getHtml(url string) *http.Response {
	response, error := http.Get(url)
	check(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	return response
}

func scrapePageData(doc *goquery.Document) {
	doc.Find("ul.srp-results>li.s-item").Each(func(index int, item *goquery.Selection) {
		a := item.Find("a.s-item__link")

		title := strings.TrimSpace(a.Text())
		url, _ := a.Attr("href")

		price_span := item.Find("span.s-item__price").Text()
		fmt.Println(price_span, title, url)
	})
}

func main() {
	url := "https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2334524.m570.l1313&_nkw=beatles+puzzle&_sacat=0&_odkw=beatles+pluzzle&_osacat=0"

	response := getHtml(url)
	defer response.Body.Close()

	doc, error := goquery.NewDocumentFromReader(response.Body)
	check(error)

	scrapePageData(doc)
}
