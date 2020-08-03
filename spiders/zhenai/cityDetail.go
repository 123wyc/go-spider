package zhenai

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"webterren.com/demo/engine"
	"webterren.com/demo/log"
)

func ParseCityDetail(city *html.Node) engine.ParseResult {

	item_list := htmlquery.Find(city, "//*[@id='app']/div[2]/div[2]/div[1]/div[2]/div[@class='list-item']")

	result := engine.ParseResult{}

	for _, item := range item_list {
		log.Add()
		name := htmlquery.InnerText(htmlquery.FindOne(item, "./div[2]/table/tbody/tr[1]/th/a"))
		//sex := htmlquery.InnerText(htmlquery.FindOne(item,"/div[2]/table/tbody/tr[2]/td[1]"))
		//url := htmlquery.SelectAttr(htmlquery.FindOne(item, "//div/table/tbody/tr[1]/th/a"), "href")
		//fmt.Printf("name:%s , sex:%s,url :%s  \n",name,sex,url)
		result.Items = append(result.Items, name)
		// result.Requests = append(result.Requests, engine.Request{
		// 	Url: url,
		// 	Parsefunc: func(node *html.Node) engine.ParseResult {
		// 		return ProfileParse(node, name, url)
		// 	},
		// })
	}
	li_arr := htmlquery.Find(city, "//*[@id='app']/div[2]/div[2]/div[1]/div[2]/div[21]/ul/li")
	var name = "下一页"
	for _, li := range li_arr {
		var item = li
		text := htmlquery.InnerText(item)
		fmt.Println(text)
		if strings.Contains(text, name) {
			url := htmlquery.SelectAttr(htmlquery.FindOne(item, "./a"), "href")
			//fmt.Println("url", url)
			result.Items = append(result.Items, name)
			result.Requests = append(result.Requests, engine.Request{
				Url: url,
				Parsefunc: func(node *html.Node) engine.ParseResult {
					return ParseCityDetail(node)
				},
			})
		}
	}

	return result
}
