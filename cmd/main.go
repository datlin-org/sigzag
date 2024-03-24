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
	level := flag.Int("level", 3, "Directory nesting depth")
	compareManifest := flag.Bool("compare-manifest", false, "Compare manifests")
	compareMerkle := flag.Bool("compare-merkle", false, "Compare merkle")

	flag.Parse()
	path, err := filepath.Abs(*root)

	if err != nil {
		log.Fatalf("path error: %s", err)
	}

	levelOK, err, length := checkLevel(path, *level)
	if err != nil {
		log.Fatalf("%s Inapproriate level set. Maximum nesting exepected < %d, got %d\n", err, length-1, *level)
	}

	if levelOK {
		config := crawler.Config{
			Level: *level,
		}
		if !*compareManifest && !*compareMerkle {
			err = generateManifest(path, config)
			if err != nil {
				fmt.Printf("error generating manifests, %s", err)
			}
		}
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

func checkLevel(path string, level int) (bool, error, int) {
	p := strings.Split(path, string(os.PathSeparator))
	if level >= len(p) {
		return false, fmt.Errorf("absolute path == %d nested levels. ", len(p)), len(p)
	}
	return true, nil, 0
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
