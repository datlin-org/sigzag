package main

import (
	"flag"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"github.com/KevinFasusi/sigzag/pkg/utils"
	"path/filepath"
)

func main() {
	root := flag.String("root", ".", "Root directory")
	compareManifest := flag.Bool("compare-manifest", false, "Compare manifests")
	compareMerkle := flag.Bool("compare-merkle", false, "Compare merkle")

	flag.Parse()
	path, err := filepath.Abs(*root)
	if err != nil {
		panic("Err")
	}

	config := crawler.Config{
		Ext:  "",
		Size: 0,
	}
	if !*compareManifest && !*compareMerkle {
		generateManifest(path, config)
	}

	if *compareManifest {
		var m utils.Manager
		m.CompareManifest(flag.Args()[0], flag.Args()[1])
	}
	if *compareMerkle {
		var m utils.Manager
		m.CompareMerkle(flag.Args()[0], flag.Args()[1])
	}
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
