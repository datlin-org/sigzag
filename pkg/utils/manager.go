package utils

import (
	"encoding/json"
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"log"
	"os"
	"reflect"

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
	f1, err := os.ReadFile(file1)
	if err != nil {
		log.Print("File 1 not found,", err)
	}
	f2, err := os.ReadFile(file2)
	if err != nil {
		log.Fatal()
	}
	var s1, s2 []crawler.Sig

	_ = json.Unmarshal(f1, &s1)
	_ = json.Unmarshal(f2, &s2)

	fmt.Printf("Equal:%v\n", reflect.DeepEqual(s1, s2))
}
