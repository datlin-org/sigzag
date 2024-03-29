package crawler

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

var testdata = []byte(`[
  {
    "asset": "pdf/testdata1.pdf",
    "sha256": "47da23b69a3f5329d41886bccef8eb3e112d477e8a5ef3da4994ed2e78376b0f",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "pdf/testdata2.pdf",
    "sha256": "66d443c39adea589b5bcd508d5b2d85356c7921b51d3636be40f28757b384d21",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "pdf/testdata3.pdf",
    "sha256": "976a0c9710bdc2e745c8d5a5e466d3181bd5e899b5d46170efbf379dd46bd320",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "pdf/testdata4.pdf",
    "sha256": "7e66f1ba24486c8c21c5dcce9612fab7df7b851b8630452f9544d3f2c0c2f0cd",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "pdf/testdata5.pdf",
    "sha256": "f7d0391e2d3dfb284aef716743690c035d43d5b72e1f5b569a08c08b095b9dba",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "testdata1.md",
    "sha256": "e6dbed6368694ce31739583c9f93d7fa97aaa7e0d0549ff00810ee193e10e392",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "testdata2.md",
    "sha256": "c8821204fd0c26e338d7895303087bb6be33b197bae47a8523d4c1805d5c96bd",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  },
  {
    "asset": "testdata3.md",
    "sha256": "9b0bfdf44481eb979dd3ed67cede9b4c0e265d2769a8858f35877f705743e3bd",
    "timestamp": "Fri Mar 29 12:15:07 GMT 2024"
  }
]`)

func TestDirectoryCrawler_Compare(t *testing.T) {
	var m Manager
	path := "../../testdata/"
	length := len(strings.Split(path, string(os.PathSeparator)))
	level := 3
	levelStart := length - 1
	config := Config{
		Root:    levelStart,
		Depth:   level + 1,
		TagFile: SIGZAG.Strings(),
	}
	manFile, _, _ := m.GenerateManifest(path, config)
	f, _ := os.ReadFile(manFile)

	var sig1, sig2 []SigTimeless
	_ = json.Unmarshal(f, &sig1)
	_ = json.Unmarshal(testdata, &sig2)
	if !reflect.DeepEqual(sig1, sig2) {
		t.Error("manifest for testdata is not equal")
	}
}

func TestDirectoryCrawler_Diff(t *testing.T) {
	var m Manager
	path := "../../testdata/"
	length := len(strings.Split(path, string(os.PathSeparator)))
	level := 3
	levelStart := length - 1
	config := Config{
		Root:    levelStart,
		Depth:   level + 1,
		TagFile: SIGZAG.Strings(),
	}
	manFile1, _, _ := m.GenerateManifest(path, config)
	manFile2, _, _ := m.GenerateManifest(path, config)
	d := m.Diff(manFile1, manFile2, false)
	if len(d) > 0 {
		t.Errorf("expected diff==0, actual==%d", len(d))
	}
}
