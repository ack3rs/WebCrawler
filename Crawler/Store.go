package Crawler

import (
	"sync"

	myLogger "github.com/acky666/ackyLog"
)

type Link struct {
	Visited    bool
	TotalLinks int
}

type AllLinksMap struct {
	Links map[string]Link
	Mux   sync.Mutex
}

var AllLinks = AllLinksMap{Links: make(map[string]Link)}

// Add a URL to the LinkMap
func (l *AllLinksMap) Add(URL string) {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	// Does the Key Already Exist?
	if _, exists := l.Links[URL]; exists {
		return
	}

	Parsed := 0
	for _, StoredLinks := range l.Links {
		if StoredLinks.Visited {
			Parsed++
		}
	}

	l.Links[URL] = Link{Visited: false, TotalLinks: 0}
	PercentageComplete := float32(float32(Parsed)/float32(len(l.Links))) * 100
	myLogger.DEBUG("Adding %s total urls:[F-CYAN]%d[F-NORMAL] parsed:[F-YELLOW]%d[F-NORMAL] %.2f%% complete", URL, len(l.Links), Parsed, PercentageComplete)
}

// Update a URL already in the LinkMap
func (l *AllLinksMap) Update(URL string, TotalLinks int) {
	l.Mux.Lock()
	l.Links[URL] = Link{Visited: true, TotalLinks: TotalLinks}
	l.Mux.Unlock()
}

func (l *AllLinksMap) GetLinkMap() map[string]Link {
	l.Mux.Lock()
	defer l.Mux.Unlock()
	return l.Links
}
