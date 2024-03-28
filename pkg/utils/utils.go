package utils

import (
	"encoding/json"
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/KevinFasusi/hometree"
)

type Manager struct{}

func (m *Manager) CompareMerkle(file1 string, file2 string) {
	f1, err := os.ReadFile(file1)
	if err != nil {
		log.Print("File 1 not found,", err)
	}
	f2, err := os.ReadFile(file2)
	if err != nil {
		log.Fatal()
	}
	var n1, n2 hometree.Node

	_ = json.Unmarshal(f1, &n1)
	_ = json.Unmarshal(f2, &n2)

	fmt.Printf("Equal:%v\n", reflect.DeepEqual(n1, n2))
}

func (m *Manager) CompareManifest(file1 string, file2 string) {
	s1, s2 := readManifest(file1, file2)
	fmt.Printf("Equal:%v\n", reflect.DeepEqual(s1, s2))
}

func (m *Manager) Diff(m1 string, m2 string) {
	var diff []crawler.Sig
	var remove []crawler.Sig
	s1, s2 := readManifest(m1, m2)
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
	diff = s2
	WriteDiff(diff)
}

func WriteDiff(diff []crawler.Sig) {
	timeStamp := time.Now()
	diffFile := fmt.Sprintf("diff-%d%d%d-%d%d%d.json", timeStamp.Year(), timeStamp.Month(), timeStamp.Day(),
		timeStamp.Hour(), timeStamp.Minute(), timeStamp.Second())
	file, err := os.OpenFile(diffFile, os.O_WRONLY|os.O_CREATE, 0666)
	sigJson, _ := json.Marshal(diff)
	_, err = file.Write(sigJson)
	if err != nil {
		return
	}
}

func readManifest(file1 string, file2 string) ([]crawler.Sig, []crawler.Sig) {
	f1, err := os.ReadFile(file1)
	if err != nil {
		log.Fatal("File 1 not found,", err)
	}
	f2, err := os.ReadFile(file2)
	if err != nil {
		log.Fatal("File 2 not found,", err)
	}
	var s1, s2 []crawler.Sig

	_ = json.Unmarshal(f1, &s1)
	_ = json.Unmarshal(f2, &s2)

	return s1, s2
}
