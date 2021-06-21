package Crawler

import (
	"fmt"
)

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func Results() {

	fmt.Printf("%-80s %10s %10s\n", "URL", "Parsed", "Links")
	for URL, info := range AllLinks.GetLinkMap() {
		URL = PadRight(URL, " ", 80) // Just to make it look a bit Pretty!
		fmt.Printf("%.80s %10v %10d\n", URL, info.Visited, info.TotalLinks)
	}

}
