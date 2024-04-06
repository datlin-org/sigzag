package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/KevinFasusi/hometree"
)

type Util interface {
	Compare(file1 string, file2 string, value labels)
	Diff(m1 string, m2 string, timeless bool) []Sig
	Write(label labels)
	History(asset string, args []string)
}

type Manager struct {
	Sig    []Sig
	Hist   []History
	Merkle hometree.Node
}

func (m *Manager) Compare(file1 string, file2 string, value labels) {
	switch value {
	case MANIFEST:
		s1 := Read(file1, MANIFEST, true).Sig
		s2 := Read(file2, MANIFEST, true).Sig
		//strip timestamp
		fmt.Printf("Equal:%v\n", reflect.DeepEqual(s1, s2))
	case MERKLETREE:
		n1 := Read(file1, MERKLETREE, true).Merkle
		n2 := Read(file2, MERKLETREE, true).Merkle
		//strip timestamp
		fmt.Printf("Equal:%v\n", reflect.DeepEqual(n1, n2))
	default:
		panic("unhandled default case")
	}
}

func (m *Manager) Diff(m1 string, m2 string, timeless bool) []Sig {
	var remove []Sig
	s1 := Read(m1, MANIFEST, timeless).Sig
	s2 := Read(m2, MANIFEST, timeless).Sig
	for _, i := range s2 {
		for _, j := range s1 {
			if i.Digest == j.Digest {
				remove = append(remove, j)
			}
		}
	}
	for _, sig1 := range s1 {
		for n, sig2 := range s2 {
			for _, k := range remove {
				if sig1.Digest == k.Digest && sig1.Digest == sig2.Digest {
					s2 = append(s2[:n], s2[n+1:]...)
				}
			}
		}
	}
	m.Sig = s2
	m.Write(DIFF)
	return m.Sig
}

func (m *Manager) Write(label labels) {
	switch label {
	case DIFF:
		sigJson, _ := json.Marshal(m.Sig)
		toFile(label, sigJson)
	case HISTORY:
		sigJson, _ := json.Marshal(m.Hist)
		toFile(label, sigJson)
	default:
		log.Fatalf("unknown signature type, expected %s or %s, got==%s", DIFF.Strings(), HISTORY.Strings(), label.Strings())
	}
}

func toFile(label labels, toJson []byte) {
	timeStamp := time.Now()
	outFile := fmt.Sprintf("%s-%d%d%d-%d%d%d.json", label.Strings(), timeStamp.Year(), timeStamp.Month(), timeStamp.Day(),
		timeStamp.Hour(), timeStamp.Minute(), timeStamp.Second())
	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0666)
	_, err = file.Write(toJson)
	if err != nil {
		return
	}
}

func Read(file string, label labels, timeless bool) Manager {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("File not found,", err)
	}

	switch label {
	case MANIFEST:
		var sig []Sig
		var sigTime []SigTimeless
		if timeless {
			_ = json.Unmarshal(f, &sigTime)
		} else {
			_ = json.Unmarshal(f, &sig)
		}

		return Manager{
			Sig:    sig,
			Hist:   nil,
			Merkle: hometree.Node{},
		}
	case MERKLETREE:
		var n hometree.Node
		_ = json.Unmarshal(f, &n)
		return Manager{
			Sig:    nil,
			Hist:   nil,
			Merkle: n}
	default:
		log.Fatalf("unknown signature type, expected %s or %s, got==%s", MANIFEST.Strings(), MERKLETREE.Strings(), label.Strings())
	}
	return Manager{}
}

type History struct {
	Asset   string
	History []Sig
}

func (m *Manager) History(asset string, args []string) {
	var rec []Sig
	var history History
	var historyRec []History
	for _, i := range args {
		f := Read(i, MANIFEST, false).Sig
		for _, j := range f {
			if asset == j.Asset {
				rec = append(rec, j)
			}
		}
	}
	history.Asset = asset
	history.History = rec
	historyRec = append(historyRec, history)
	m.Hist = historyRec
	m.Write(HISTORY)
}

func (m *Manager) GenerateManifest(path string, config Config) (string, string, error) {
	crawl := NewDirectoryCrawler(path, &config)
	err := crawl.Crawl()
	if err != nil {
		return "", "", fmt.Errorf("unable to crawl directory, %s", err)
	}
	manifestFile, err := crawl.Write(MANIFEST)
	if err != nil {
		return "", "", fmt.Errorf("unable to write manifest. %s", err)
	}
	merkleFile, err := crawl.Write(MERKLETREE)
	if err != nil {
		return "", "", fmt.Errorf("unable to write Merkle tree. %s", err)
	}
	return manifestFile, merkleFile, nil
}

func (m *Manager) Download(config Config, label labels) {

	switch label {
	case URL:
		downloadUrl(config)
	case URLS:
		downloadUrls(config)
	default:
		panic("unhandled default case")
	}
}

func downloadUrl(config Config) {
	w := NewWebCrawler(&config)
	urlParts := strings.Split(config.Url, "/")
	f := urlParts[len(urlParts)-1]
	file, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)
	err = w.Download(file, 3)
	if err != nil {
		return
	}
	fileStat, err := file.Stat()
	var d DirectoryCrawler
	fmt.Printf("%v bytes downloaded\n%v MB downloaded\n", fileStat.Size(), fileStat.Size()/1000000)
	f2, err := os.OpenFile(f, os.O_RDONLY, 0644)
	_ = f2.Close()
	fmt.Printf("sha256: %x\n", d.FileSignature(f2.Name()))
}

func downloadUrls(config Config) {
	f1, ok := os.ReadFile(config.Urls)
	w := NewWebCrawler(&config)
	var res []*UrlResult
	if ok != nil {
		return
	}
	var u []Urls
	err := json.Unmarshal(f1, &u)
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, i := range u {
		s := i
		w.Conf.Url = s.Url
		urlParts := strings.Split(s.Url, "/")
		f := urlParts[len(urlParts)-1]
		file, err := os.Create(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				return
			}
		}(file)
		err = w.Download(file, 3)
		if err != nil {
			return
		}
		res = append(res, compareDownloadSHA(s, file, f))
	}
	writeDownloadManifest(res)
}

func compareDownloadSHA(s Urls, file *os.File, name string) *UrlResult {
	fileStat, _ := file.Stat()
	var match bool
	var d DirectoryCrawler
	size := fileStat.Size()
	fmt.Printf("%v bytes downloaded\n%vMB downloaded\n", size, size/1000000)
	f2, _ := os.OpenFile(name, os.O_RDONLY, 0644)
	_ = f2.Close()
	sha := fmt.Sprintf("%x", d.FileSignature(f2.Name()))
	fmt.Printf("sha256: %s\n", sha)
	if sha == s.Sha256 {
		match = true
		fmt.Printf("Match: true\n")
	} else {
		match = false
		fmt.Printf("Match: false\n")
	}
	return &UrlResult{
		File:   f2.Name(),
		Sha256: sha,
		Size:   size,
		Match:  match,
	}

}

type UrlResult struct {
	File   string `json:"file"`
	Sha256 string `json:"sha256,omitempty"`
	Size   int64  `json:"size,omitempty"`
	Match  bool   `json:"match,omitempty"`
}

func writeDownloadManifest(res []*UrlResult) {
	marshal, err := json.Marshal(res)
	if err != nil {
		return
	}
	toFile(DOWNLOAD, marshal)
}
