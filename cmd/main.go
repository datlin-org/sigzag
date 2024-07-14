package main

import (
	"flag"
	"fmt"
	"github.com/datlin-org/sigzag/pkg/crawler"
	"github.com/datlin-org/sigzag/pkg/services"

	//d2 "github.com/datlin-org/sigzag/pkg/daemon"
	"github.com/datlin-org/sigzag/pkg/scraper"
	//"github.com/datlin-org/sigzag/pkg/terminal"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	addr := flag.String("addr", ":0000", "HTTP network address")
	root := flag.String("root", ".", "Root directory")
	level := flag.Int("level", 2, "Maximum directory nesting depth")
	diffManifest := flag.Bool("diff", false, "Compare two manifests and return the difference if any")
	asset := flag.String("asset", crawler.ASSET.Strings(), "Manifests")
	history := flag.Bool("history", false, "Returns the history of an asset")
	tagFile := flag.String("tag-file", crawler.SIGZAG.Strings(), "Prepends the output file with string")
	compareManifest := flag.Bool("compare-manifest", false, "Compare manifests")
	compareMerkle := flag.Bool("compare-merkle", false, "Compare merkle")
	out := flag.String("output-file", "", "Directory for output")
	url := flag.String("url", crawler.URL.Strings(), "Download asset and show sha256 checksum")
	urls := flag.String("urls", crawler.URLS.Strings(), "Download assets from a list of urls in a file and generate a manifest containing checksums")
	datasource := flag.String("datasource", crawler.DATASOURCE.Strings(), "Locate any data source type at location")
	//scan := flag.String("scan", crawler.ASSET.Strings(), "Crawl file")
	term := flag.Bool("terminal", false, "Launch terminal")
	//daemon := flag.Bool("daemon", false, "Launch daemon")

	flag.Parse()
	path, err := filepath.Abs(*root)

	if *level <= 0 {
		log.Fatalf("level must be greater than 0. --level %d flag set\n", *level)
	}

	if err != nil {
		log.Fatalf("path error: %s", err)
	}

	config, _ := CrawlerConfig(level, tagFile, out, url, urls, path)

	if !*compareManifest && !*compareMerkle && !*diffManifest && !*history && *asset == crawler.ASSET.Strings() &&
		*url == crawler.URL.Strings() && *urls == crawler.URLS.Strings() && *datasource == crawler.DATASOURCE.Strings() &&
		!*term && *addr == "0000" {
		var m crawler.Manager
		_, _, err = m.GenerateManifest(path, config)
		if err != nil {
			fmt.Printf("error generating manifests, %s", err)
		}
	}

	//if *term {
	//var t terminal.Terminal
	//t.New()
	//	terminal.Launch()
	//}

	//if *daemon {
	//	d2.Run()
	//}

	if *addr != "0000" {
		err = services.RunService("75000")
		if err != nil {
			return
		}
	}

	if *datasource != crawler.DATASOURCE.Strings() {
		conf := scraper.Config{Url: *datasource}
		sc := scraper.Scraper{Conf: conf}
		sc.Scrape(conf)
	}

	if *url != crawler.URL.Strings() {
		var m crawler.Manager
		m.Download(config, crawler.URL)
	}

	if *urls != crawler.URLS.Strings() {
		var m crawler.Manager
		m.Download(config, crawler.URLS)
	}

	if *diffManifest {
		var m crawler.Manager
		_ = m.Diff(flag.Args()[0], flag.Args()[1], false)
	}

	if *history && *asset != crawler.ASSET.Strings() {
		var m crawler.Manager
		m.History(*asset, flag.Args())
	}
	if *compareManifest {
		var m crawler.Manager
		m.Compare(flag.Args()[0], flag.Args()[1], crawler.MANIFEST)
	}
	if *compareMerkle {
		var m crawler.Manager
		m.Compare(flag.Args()[0], flag.Args()[1], crawler.MERKLETREE)
	}
}

func CrawlerConfig(level *int, tagFile *string, out *string, url *string, urls *string, path string) (crawler.Config, error) {
	length := len(strings.Split(path, string(os.PathSeparator)))
	levelStart := length - 1
	validUrl, valid, err := crawler.ValidateUrl(*url)
	if !valid {
		return crawler.Config{}, err
	}
	return crawler.Config{
		Root:    levelStart,
		Depth:   *level + 1,
		TagFile: *tagFile,
		OutDir:  *out,
		Url:     validUrl,
		Urls:    *urls,
	}, nil
}
