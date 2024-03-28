package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/KevinFasusi/hometree"
)

type Util interface {
	Compare(file1 string, file2 string, value labels)
	Diff(m1 string, m2 string)
	Write(label labels)
	History(asset string, args []string)
}

type Manager struct {
	Sig  []Sig
	Hist []History
}

func (m *Manager) Compare(file1 string, file2 string, value labels) {
	switch value {
	case MANIFEST:
		s1 := readManifest(file1)
		s2 := readManifest(file2)
		fmt.Printf("Equal:%v\n", reflect.DeepEqual(s1, s2))
	case MERKLETREE:
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
	default:
		panic("unhandled default case")
	}
}

func (m *Manager) Diff(m1 string, m2 string) {
	var remove []Sig
	s1 := readManifest(m1)
	s2 := readManifest(m2)
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
		panic("unhandled default case")
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

func readManifest(file string) []Sig {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("File 1 not found,", err)
	}
	var sig []Sig
	_ = json.Unmarshal(f, &sig)
	return sig
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
		f := readManifest(i)
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
