package main

import (
	"flag"
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := flag.String("root", ".", "Root directory")
	level := flag.Int("level", 2, "Maximum directory nesting depth")
	diffManifest := flag.Bool("diff", false, "Compare two manifests and return the difference if any")
	asset := flag.String("asset", crawler.ASSET.Strings(), "Manifests")
	history := flag.Bool("history", false, "Returns the history of an asset")
	outputFile := flag.String("output-file", crawler.SIGZAG.Strings(), "Prepends the output file with string")
	compareManifest := flag.Bool("compare-manifest", false, "Compare manifests")
	compareMerkle := flag.Bool("compare-merkle", false, "Compare merkle")

	flag.Parse()
	path, err := filepath.Abs(*root)

	if *level <= 0 {
		log.Fatalf("level must be greater than 0. --level %d flag set\n", *level)
	}

	if err != nil {
		log.Fatalf("path error: %s", err)
	}
	length := len(strings.Split(path, string(os.PathSeparator)))

	levelStart := length - 1
	config := crawler.Config{
		Root:       levelStart,
		Depth:      *level + 1,
		OutputFile: *outputFile,
	}

	if !*compareManifest && !*compareMerkle && !*diffManifest && !*history && *asset == crawler.ASSET.Strings() {
		err = generateManifest(path, config)
		if err != nil {
			fmt.Printf("error generating manifests, %s", err)
		}
	}

	if *diffManifest {
		var m crawler.Manager
		m.Diff(flag.Args()[0], flag.Args()[1])
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

func generateManifest(path string, config crawler.Config) error {
	crawl := crawler.NewCrawler(path, &config)
	err := crawl.Crawl()
	if err != nil {
		return fmt.Errorf("unable to crawl directory, %s", err)
	}
	err = crawl.Write(crawler.MANIFEST)
	if err != nil {
		return fmt.Errorf("unable to write manifest. %s", err)
	}
	err = crawl.Write(crawler.MERKLETREE)
	if err != nil {
		return fmt.Errorf("unable to write Merkle tree. %s", err)
	}
	return nil
}
