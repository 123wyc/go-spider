package downloader

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var rateLimiter = time.Tick(300 * time.Millisecond)

func DownLoad(url string) (node *html.Node, err error) {

	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {

		return &html.Node{}, err
	}

	if resp.StatusCode != http.StatusOK {

		return &html.Node{}, fmt.Errorf("wrong status code")

	}

	node, err = htmlquery.Parse(resp.Body)

	defer resp.Body.Close()

	if nil != err {
		return &html.Node{}, err
	}
	return node, nil
}
