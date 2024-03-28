package main

import (
	"flag"
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"github.com/KevinFasusi/sigzag/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := flag.String("root", ".", "Root directory")
	level := flag.Int("level", 2, "Maximum directory nesting depth")
	diffManifest := flag.Bool("diff", false, "Compare two manifests and return the difference if any")
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

	if !*compareManifest && !*compareMerkle && !*diffManifest {
		err = generateManifest(path, config)
		if err != nil {
			fmt.Printf("error generating manifests, %s", err)
		}
	}

	if *diffManifest {
		var m utils.Manager
		m.Diff(flag.Args()[0], flag.Args()[1])
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

func generateManifest(path string, config crawler.Config) error {
	crawl := crawler.NewCrawler(path, &config)
	err := crawl.Crawl()
	if err != nil {
		return fmt.Errorf("unable to crawl directory, %s", err)
	}
	err = crawl.Write(crawler.Manifest)
	if err != nil {
		return fmt.Errorf("unable to write manifest. %s", err)
	}
	err = crawl.Write(crawler.MerkleTree)
	if err != nil {
		return fmt.Errorf("unable to write Merkle tree. %s", err)
	}
	return nil
}
