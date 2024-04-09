package crawler

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/KevinFasusi/hometree"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type labels int

const (
	MANIFEST labels = iota
	MERKLETREE
	SIGZAG
	ASSET
	DIFF
	HISTORY
	DIRECTORY
	WEB
	URL
	URLS
	DOWNLOAD
)

func (l labels) Strings() string {
	return [...]string{
		"manifest",
		"merkletree",
		"sigzag",
		"asset",
		"diff",
		"history",
		"directory",
		"web",
		"url",
		"urls",
		"download",
	}[l]
}

type Crawler interface {
	Crawl() error
}

type Urls struct {
	Url    string `json:"url"`
	Sha256 string `json:"sha256,omitempty"`
	Size   string `json:"size,omitempty"`
}

type Config struct {
	Root    int
	Depth   int
	TagFile string
	OutDir  string
	Url     string
	Urls    string
}

type DirectoryCrawler struct {
	Dir         string           `json:"dir"`
	Regex       []*regexp.Regexp `json:"regex"`
	Conf        Config           `json:"conf"`
	FileDigests [][]byte         `json:"file_digests"`
	Signatures  []*Sig           `json:"signatures"`
}

type Sig struct {
	Asset     string `json:"asset"`
	Digest    string `json:"sha256"`
	Timestamp string `json:"timestamp"`
}

type SigTimeless struct {
	Asset     string `json:"asset"`
	Digest    string `json:"sha256"`
	Timestamp string `json:"_"`
}

func (d *DirectoryCrawler) Crawl() error {
	return filepath.WalkDir(d.Dir, d.signatureWalk)
}

// signatureWalk traverses a directory and cryptographically signs the files
func (d *DirectoryCrawler) signatureWalk(path string, info fs.DirEntry, _ error) error {
	var wg sync.WaitGroup
	if !info.IsDir() {
		wg.Add(1)
		go func() {
			signature := d.FileSignature(path)
			d.FileDigests = append(d.FileDigests, signature)
			p := strings.Split(path, string(os.PathSeparator))
			p = p[d.Conf.Root:]
			if len(p) <= d.Conf.Depth {
				path = strings.Join(p, string(os.PathSeparator))
				s := Sig{
					Asset:     path,
					Digest:    fmt.Sprintf("%x", signature),
					Timestamp: time.Now().Format(time.UnixDate),
				}
				d.Signatures = append(d.Signatures, &s)
			}
			wg.Done()
		}()
		wg.Wait()
	}
	return nil
}

// FileSignature generates a hash value for files found in the path
func (d *DirectoryCrawler) FileSignature(path string) []byte {
	var sum []byte
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("ERR\n")
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal()
		}
	}(file)

	buf := make([]byte, 8192)
	fileSignature := sha256.New()

	for b := 0; err == nil; {
		b, err = file.Read(buf)
		if err == nil {
			_, err = fileSignature.Write(buf[:b])
		}
	}
	sum = fileSignature.Sum(nil)
	return sum
}

// Write the manifest to disk
func (d *DirectoryCrawler) Write(fileType labels) (string, error) {
	s := sha256.New()
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return "", fmt.Errorf("randomise bytes failed, %s", err)
	}
	s.Write(rb)
	timeStamp := time.Now()
	timeStampFmt := fmt.Sprintf("%d%d%d-%d%d%d", timeStamp.Year(), timeStamp.Month(), timeStamp.Day(),
		timeStamp.Hour(), timeStamp.Minute(), timeStamp.Second())
	var fileName string
	switch fileType {
	case MANIFEST:
		if d.Conf.TagFile == SIGZAG.Strings() {
			fileName = fmt.Sprintf("%s-%s-%x.json", MANIFEST.Strings(), timeStampFmt, s.Sum(nil))
		}
		if d.Conf.TagFile != SIGZAG.Strings() {
			fileName = fmt.Sprintf("%s-%s-%s-%x.json", d.Conf.TagFile, MANIFEST.Strings(), timeStampFmt, s.Sum(nil))
		}
		if d.Conf.OutDir != "" && d.Conf.TagFile == SIGZAG.Strings() {
			fileName = fmt.Sprintf("%s%s", d.Conf.OutDir, fileName)
		}
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		sigJson, _ := json.Marshal(d.Signatures)
		_, err = file.Write(sigJson)
		if err != nil {
			return "", nil
		}
		return fileName, nil
	case MERKLETREE:
		if d.Conf.TagFile == SIGZAG.Strings() {
			fileName = fmt.Sprintf("%s-%s-%x.json", MERKLETREE.Strings(), timeStampFmt, s.Sum(nil))
		}
		if d.Conf.TagFile != SIGZAG.Strings() {
			fileName = fmt.Sprintf("%s-%s-%s-%x.json", d.Conf.TagFile, MERKLETREE.Strings(), timeStampFmt, s.Sum(nil))
		}
		home, err := d.buildMerkleTree()
		if err != nil {
			log.Fatalf("building merkle tree failed")
		}
		if d.Conf.OutDir != "" {
			fileName = fmt.Sprintf("%s%s", d.Conf.OutDir, fileName)
		}
		file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		sigJson, _ := json.Marshal(home)
		_, err = file.Write(sigJson)
		if err != nil {
			log.Print("File failed to write", err)
		}
		return fileName, nil

	default:
		panic("unhandled default case")
	}
	return "", nil
}

// buildMerkleTree builds a merkle tree using the SHA256 of the file as the preimage for the leaf nodes. The interior
// nodes are digests using BLAKE2b
func (d *DirectoryCrawler) buildMerkleTree() (*hometree.Node, error) {
	var homeTree hometree.MerkleTree
	home, merkleErr := homeTree.Build(d.FileDigests)
	if merkleErr.Err != nil {
		return nil, fmt.Errorf("merkle tree err, %s", merkleErr.Error())
	}
	return home, nil
}

// NewDirectoryCrawler instantiates a new directory crawler object
func NewDirectoryCrawler(root string, conf *Config) *DirectoryCrawler {
	return &DirectoryCrawler{
		Dir:  root,
		Conf: *conf}
}

// NewWebCrawler instantiates a new web crawler object
func NewWebCrawler(conf *Config) *WebCrawler {
	return &WebCrawler{
		Conf: *conf,
	}
}
