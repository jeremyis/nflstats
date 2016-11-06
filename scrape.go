package main

import (
  "fmt"
  "net/http"

  "github.com/yhat/scrape"
  "golang.org/x/net/html"
  "golang.org/x/net/html/atom"
)

func main() {
  // request and parse the front page
  url := "http://www.nfl.com/stats/categorystats?tabSeq=2&offensiveStatisticCategory=GAME_STATS&conference=ALL&role=TM&season=2016&seasonType=REG&d-447263-s=TOTAL_YARDS_GAME_AVG&d-447263-o=2&d-447263-n=1"
  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  root, err := html.Parse(resp.Body)
  if err != nil {
    panic(err)
  }

  // define a matcher
  matcher := func(n *html.Node) bool {
    // must check for nil values

    // table#result tbody + tbody
    return n.DataAtom == atom.Tr &&
        (scrape.Attr(n, "class") == "odd" || scrape.Attr(n, "class") == "even") &&
        n.Parent != nil &&
        n.Parent.DataAtom == atom.Tbody &&
        n.Parent.Parent != nil &&
        n.Parent.Parent.DataAtom == atom.Table &&
        scrape.Attr(n.Parent.Parent, "id") == "result";
  }
  // grab all articles and print them
  nodes := scrape.FindAll(root, matcher)

  /*
  c := n.FirstChild // tr
  fmt.Printf("%s\n", c.DataAtom)
  c = c.NextChild
  fmt.Printf("%s\n", c.DataAtom) // td -> rank
  */
  for _, n := range nodes {
    tds := scrape.FindAll(n, func(n *html.Node) bool { return n.DataAtom == atom.Td })
    first := tds[0]
    second := tds[1]
    //td := n.FirstChild
    // tr. td
    //rankCell := r.FirstChild
    //c := n.FirstChild
    fmt.Printf("t.: %s %s\n", scrape.Text(first), scrape.Text(second))
    }
  /*
  for r := n.FirstChild; r != nil; r = r.NextSibling { // r = tr
  }
  */
  fmt.Printf("yo")
  /*
  for i, article := range articles {
    fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
  }
  */
}
