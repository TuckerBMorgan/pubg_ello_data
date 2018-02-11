package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"time"
	"golang.org/x/net/html"
)

func generateAllLinks() []string {
	//strings that are used in the URL for iding each server region
	//Asia, North America, Korea/Japan, Kakao(seems to be a korean ISP, not sure why it is its own server), South America, Europe, Oceania, South East Asia
	serverStrings := []string{"as", "na"} //, "krjp", "kakao", "sa", "eu", "oc", "sea"}
	startString := "https://pubg.op.gg/leaderboard/?server="
	returnStrings := []string{}

	//for each server string region code
	for _, s := range serverStrings {
		//add the code to the url to give it the region
		nextString := startString + s
		for i := 1; i < 3; i++ {
			//we want the single and duo queue, which is done by adding either 1, or 2 to the end of the url
			nextString := nextString + "&mode=tpp&queue_size=" + strconv.Itoa(i)
			returnStrings = append(returnStrings, nextString)
		}
	}

	return returnStrings
}

func collectData(url string) []string {

	scoreStore := []string{}
	//make a request for the page(leaderboard)
	resp, err := http.Get(url)
	if err != nil {
		//check to makesure that we got some request
		fmt.Printf("HEIMDALL Error: There was an error getting %v, %v\n", url, err)
		return []string{}
	}
	parsedPage, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("NORN Error: There was an error parsing the page, %v", err)
		return []string{}
	}
	store := scraper(parsedPage)
	for _, s := range store {
		scoreStore = append(scoreStore, s)
	}

	return scoreStore
}

func scraper(n *html.Node) []string {
	scoreStore := []string{}
	//is this an element node
	if n.Type == html.ElementNode {
		//for each of the attributes that the element has
		for _, b := range n.Attr {
			//if they are the value for the leaderboard rating text
			if b.Val == "leader-board-top3__rating-value" {
				//	fmt.Printf("%v\n", n.FirstChild.Data)
				scoreStore = append(scoreStore, n.FirstChild.Data)
			}
		}
	}

	//traverse the tree of the page
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returns := scraper(c)
		for _, s := range returns {
			scoreStore = append(scoreStore, s)
		}
	}

	return scoreStore
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//create all the links for all of the different leaderboard regions, and the for single and duo queue
	links := generateAllLinks()
	servers := []string{"as", "na"}
	count := 0
	mod := 2
	f, err := os.OpenFile("../../../pubg_project_site/data_file.txt", os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	f.WriteString(fmt.Sprintf("%v\n", time.Now()))
	check(err)
	for _, link := range links {
		innerCount := 1
		scoreStore := collectData(link)
		for _, score := range scoreStore {
			dumb := []byte(fmt.Sprintf("%v %v %v %v\n", score, servers[count / mod], (count + 1) - (2 * (count / mod)), innerCount))
			_, err := f.Write(dumb);
			check(err)
	//		fmt.Printf("%v\n", n2)
			innerCount++
		}
		count++
	}
	defer f.Close()
}
