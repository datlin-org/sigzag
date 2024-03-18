package main

import (
	"flag"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"path/filepath"
)

func main() {
	// creates a ramdisk with running instances of sigzag
	root := flag.String("root", ".", "Root directory")
	flag.Parse()
	path, err := filepath.Abs(*root)
	if err != nil {
		panic("Err")
	}

	config := crawler.Config{
		Ext:  "",
		Size: 0,
	}
	generateManifest(path, config)
}

func generateManifest(path string, config crawler.Config) {
	crawl := crawler.NewCrawler(path, &config)
	err := crawl.Crawl()
	if err != nil {
		return
	}
	crawl.Write(crawler.Manifest)
	crawl.Write(crawler.MerkleTree)
}
