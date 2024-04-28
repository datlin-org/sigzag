package scraper

import (
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

type Config struct {
	Url string
}
type Scraper struct {
	Conf Config
}

func (s *Scraper) Scrape(config Config) {
	sc := Scraper{
		Conf: config,
	}
	err := sc.Crawl()
	if err != nil {
		return
	}
}

func (s *Scraper) Crawl() error {
	c := colly.NewCollector(
	//colly.AllowedDomains(""),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		linkSplit := strings.Split(link, ".")
		extension := linkSplit[len(linkSplit)-1]
		if extension == crawler.TXT.Strings() && strings.ToLower(linkSplit[len(linkSplit)-2]) == crawler.ROBOTS.Strings() {
			filename := strings.ToLower(linkSplit[len(linkSplit)-2])
			fmt.Printf("%s.%s found\n", filename, extension)
			log.Fatal("Crawl terminated")
		}
		if extension == crawler.XLSX.Strings() || extension == crawler.XLSM.Strings() ||
			extension == crawler.XLSB.Strings() || extension == crawler.CSV.Strings() ||
			extension == crawler.ARFF.Strings() || extension == crawler.IPYNB.Strings() ||
			extension == crawler.PARQUET.Strings() || extension == crawler.ZIP.Strings() ||
			extension == crawler.BIN.Strings() || extension == crawler.PDF.Strings() ||
			extension == crawler.GZ.Strings() {

			conf := crawler.Config{
				Url: e.Request.AbsoluteURL(link),
			}
			crawler.DownloadUrl(conf)
		}
		//err := c.Visit(e.Request.AbsoluteURL(link))
		//if err != nil {
		//	return
		//}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(s.Conf.Url)
	if err != nil {
		return err
	}
	return err
}
