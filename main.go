package main

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

func generateAllLinks() []string {
	serverStrings := []string{"as", "na", "krjp", "kakao", "sa", "eu", "oc", "sea"}
	//queueStrings := []string{"tpp", "fpp"}
	startString := "https://pubg.op.gg/leaderboard/?server="
	//  nextString := ""
	returnStrings := []string{}

	for _, s := range serverStrings {
		nextString := startString + s
		//  fmt.Printf("%v\n", nextString)
		//returnStrings = append(returnStrings, nextString)
		for i := 1; i < 3; i++ {
			nextString := nextString + "&mode=tpp&queue_size=" + strconv.Itoa(i)
			returnStrings = append(returnStrings, nextString)
		}
	}

	return returnStrings
}

func collectData(urls []string) {

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("HEIMDALL Error: There was an error getting %v, %v\n", url, err)
			continue
		}
		parsedPage, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Printf("NORN Error: There was an error parsing the page, %v", err)
			continue
		}
	}
}

func main() {
	links := generateAllLinks()

	for _, l := range links {
		fmt.Printf("%v\n", l)
	}

	return
	//this makes a request for the asian server, thrid person, singe queue
	resp, err := http.Get("https://pubg.op.gg/leaderboard/?server=as&mode=tpp&queue_size=1")

	//was there any error any making the request
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	r := resp.Body
	parsedPage, err := html.Parse(r)

	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, b := range n.Attr {
				if b.Val == "leader-board-top3__rating-value" {
					fmt.Printf("%v\n", n.FirstChild.Data)
				}
			}
			//  fmt.Println()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(parsedPage)
}
