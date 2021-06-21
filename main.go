/*
	WebCrawler
*/
package main

import (
	"flag"
	"os"
	"os/signal"

	l "github.com/acky666/ackyLog"

	Crawler "github.com/acky666/WebCrawler/Crawler"
)

func main() {

	// As this can take a long time,  Catch the Control C and Print out the Progress so far.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Print Progress and Exit
	go func() {
		oscall := <-c
		l.WARNING("User or System Interupted: %v, dumping progress so far. ", oscall)
		Crawler.Results()
		os.Exit(1)
	}()

	URLtoCrawl := flag.String("url", "", "The URL to Crawl")
	Concurrent := flag.Int("c", 10, "Max number of Concurrent Goroutines")
	Debug := flag.Bool("debug", true, "Display Debug messages")
	flag.Parse()

	if *Debug {
		l.SHOWDEBUG = true
	}

	// Nothing Specified
	if *URLtoCrawl == "" {
		l.INFO("Please specify a URL on using the -url command argument\n")
		os.Exit(1)
	}

	// Start Crawling!
	l.INFO("Crawling '%s'\n", *URLtoCrawl)

	// Add the Starting URL (the one from the command line)
	Crawler.AllLinks.Add(*URLtoCrawl)

	// Go Go Go
	Crawler.Go(*Concurrent)

	// Output Results
	Crawler.Results()
}
