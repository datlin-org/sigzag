package scraper

import "testing"

func TestScrape(t *testing.T) {
	var conf Config
	conf.Url = ""
	scraper := Scraper{
		Conf: conf,
	}
	err := scraper.Crawl()
	if err != nil {
		return
	}
}
