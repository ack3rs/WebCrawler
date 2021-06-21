package Crawler

import "sync"

// Start the Crawler
func Go(MaxNumberOfGoRoutines int) {

	// Just in Case!
	if MaxNumberOfGoRoutines < 1 {
		MaxNumberOfGoRoutines = 1
	}

	var wg sync.WaitGroup

	// Loop through all the Links, until there isn't any Visited = false left!
	for {

		// Have all the links / URL's been visited?
		HaveAllTheLinksBeenVisited := true

		Concurrent := 0

		for url, link := range AllLinks.GetLinkMap() {
			if !link.Visited {
				go Crawl(url, &wg)
				wg.Add(1)
				HaveAllTheLinksBeenVisited = false

				// Unless this code is here, the go routines will keep be added for every pass,  to many and yu can drown the server!.
				Concurrent++
				if Concurrent > MaxNumberOfGoRoutines {
					break // Stop adding more GoRoutines!
				}
			}
		}

		if HaveAllTheLinksBeenVisited {
			break // the main loop
		}

		// Wait Until all the Go Routines are done, before another cycle
		wg.Wait()
	}
}
