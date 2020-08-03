package zhenai

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"webterren.com/demo/engine"
)

func ParseCityList(node *html.Node) engine.ParseResult {

	city_list := htmlquery.Find(node, "//*[@id='app']/article/dl/dd/a")

	parseResult := engine.ParseResult{}
	limit := 400
	for _, a := range city_list {

		href := htmlquery.SelectAttr(a, "href")
		city := htmlquery.InnerText(a)
		fmt.Printf("city:%s ,url:%s  \n", city, href)

		parseResult.Items = append(parseResult.Items, city)
		parseResult.Requests = append(
			parseResult.Requests, engine.Request{
				Url:       href,
				Parsefunc: ParseCityDetail,
			})
		limit--
		if 0 == limit {
			break
		}
	}
	return parseResult
}
