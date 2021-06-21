package Crawler

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	l "github.com/acky666/ackyLog"
	"golang.org/x/net/html"
)

// Crawl through a URL,  making a note of all External and Internal Links.
func Crawl(StartingURL string, wg *sync.WaitGroup) {

	TotalLinks := 0

	// When are done ... We can exit the crawl in a number of places, so makes use the All Links are Updated and Tell the WG!
	defer wg.Done()

	// Make the Web Call to the URL
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(StartingURL)

	// TODO:
	//      Only work if it's a valid 200?
	//
	//      At the momemt it doesn't care if it's a 200 or 404

	if err != nil {
		l.ERROR("Unable to Crawl %s Error: %s", StartingURL, err)
		AllLinks.Update(StartingURL, TotalLinks)
		return
	}

	// Parse the Results
	DocToken := html.NewTokenizer(response.Body)
	for {
		DocTokenNext := DocToken.Next()

		switch DocTokenNext {
		case html.ErrorToken:

			// End of Document or Error
			AllLinks.Update(StartingURL, TotalLinks)
			return

		case html.StartTagToken, html.EndTagToken:

			token := DocToken.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {

						// Clean Up the URL
						URLtoAdd, err := cleanURL(StartingURL, attr.Val)
						if err == nil {
							TotalLinks++
							AllLinks.Add(URLtoAdd)
						}

					} // href
				} // for
			} // if a
		}
	}
}

// Test 2 URLS and see if the domains Match.
func doDomainsMatch(D1, D2 string) bool {

	PD1, err1 := url.Parse(D1)
	if err1 != nil {
		l.ERROR("Domain %s, failed to Parse Ignoring (%v)", D1, err1)
		return false
	}

	PD2, err2 := url.Parse(D2)
	if err2 != nil {
		l.ERROR("Domain %s, failed to Parse Ignoring (%v)", D2, err2)
		return false
	}

	if PD1.Host == PD2.Host {
		return true
	}
	return false
}

// It's my Job to Clean up the URL based on the Source URL ..   Building an Absolute URL
//
// "http://www.monzo.com/cdn-cgi/l/email-protection#241b024549541f5751464e4147501967"
// There are a ton of URLS on Monzo that have a randomly generated hash at the end,  they are the same page.  So let's clean them
//
func cleanURL(SourceURL, LinkURL string) (string, error) {

	//l.DEBUG("Found Link FROM at %s FROM %s", LinkURL, SourceURL)

	ParsedURL, err := url.Parse(LinkURL)
	if err != nil {
		return "", err
	}

	if (strings.HasPrefix(LinkURL, "https://")) || (strings.HasPrefix(LinkURL, "http://")) {
		// We have an Absolute URL

		if doDomainsMatch(SourceURL, LinkURL) {

			if ParsedURL.RawQuery == "" {
				return ParsedURL.Scheme + "://" + ParsedURL.Host + ParsedURL.Path, nil
			} else {
				return ParsedURL.Scheme + "://" + ParsedURL.Host + ParsedURL.Path + "?" + ParsedURL.RawQuery, nil
			}

		} else {
			// the Domains Don't Match, return nothing
			return "", errors.New("The Domains Don't Match")
		}

	}

	// We have a Relative URL
	SourceParsed, err := url.Parse(SourceURL)
	if err != nil {
		return "", errors.New("Unable to Parse Source URL")
	}

	// Make a Proper Relative URL
	if !strings.HasPrefix(LinkURL, "/") {
		LinkURL = "/" + LinkURL
	}

	newURL := SourceParsed.Scheme + "://" + SourceParsed.Host + LinkURL
	GoodRelativeURL, err := sanatiseURL(newURL)
	if err != nil {
		return "", err
	}

	return GoodRelativeURL, nil
}

// It's my Job to just check the URL is OK, and valid from a Schema point of view.
//
func sanatiseURL(urlToCheck string) (string, error) {

	ParsedURL, err := url.Parse(urlToCheck)
	if err != nil {
		return "", errors.New("Invalid")
	}
	if ParsedURL.RawQuery == "" {
		return ParsedURL.Scheme + "://" + ParsedURL.Host + ParsedURL.Path, nil
	} else {
		return ParsedURL.Scheme + "://" + ParsedURL.Host + ParsedURL.Path + "?" + ParsedURL.RawQuery, nil
	}
}
