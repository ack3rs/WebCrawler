package Crawler

import (
	"testing"
)

func TestSmallCrawl(t *testing.T) {

	// Add the Starting URL (the one from the command line)
	AllLinks.Add("http://www.acky.com")

	// Called
	Go(1)

	// Called but Not really test
	Results()

	jsonResults, _ := ResultsJSON()

	if string(jsonResults) != `{"http://www.acky.com":{"Visited":true,"TotalLinks":0}}` {
		t.Errorf("Failed TestSmallCrawl Test!")
	}

}

func TestBigCrawl(t *testing.T) {

	// Add the Starting URL (the one from the command line)
	AllLinks.Add("http://www.acky.com/test.html")

	// Called
	Go(1)

	// Called but Not really test
	Results()

	links := 0
	for _, info := range AllLinks.GetLinkMap() {
		links = links + info.TotalLinks
	}

	if links != 3 {
		t.Errorf("Failed TestBigCrawl Test!")
	}
}
