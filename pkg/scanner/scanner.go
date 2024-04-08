package scanner

import (
	"archive/zip"
	"fmt"
	"github.com/KevinFasusi/sigzag/pkg/crawler"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ScanType int

const (
	EXCEL ScanType = iota
	DATABASE
)

type Scanner struct {
}

func (s *Scanner) Scan(path string, scanType ScanType) {

	switch scanType {

	case EXCEL:
		scanExcel(path)
	case DATABASE:

	default:
		panic("unhandled default case")

	}
}

func scanExcel(path string) {
	sigzagDir := os.Getenv("SIGZAG_DIR")
	if sigzagDir == "" {
		hd, _ := os.UserHomeDir()
		sigzagDir = hd + string(os.PathSeparator) + ".sigzag"
	}

	destination := sigzagDir + string(os.PathSeparator) + "out"
	archive, _ := zip.OpenReader(path)
	defer func(archive *zip.ReadCloser) {
		err := archive.Close()
		if err != nil {
			return
		}
	}(archive)

	for _, f := range archive.File {
		newPath := filepath.Join(destination, f.Name)
		fmt.Printf("unzipping file: %s\n", newPath)
		if !strings.HasPrefix(newPath, filepath.Clean(destination)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
		}

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			err := os.MkdirAll(newPath, os.ModePerm)
			if err != nil {
				continue
			}
		}

		if err := os.MkdirAll(filepath.Dir(newPath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		err = dstFile.Close()
		if err != nil {
			return
		}
		err = fileInArchive.Close()
		if err != nil {
			return
		}
	}
	err := archive.Close()
	if err != nil {
		return
	}
	length := len(strings.Split(destination, string(os.PathSeparator)))
	levelStart := length - 1
	config := crawler.Config{
		Root:    levelStart,
		Depth:   2 + 1,
		TagFile: crawler.SIGZAG.Strings(),
		OutDir:  destination,
		Url:     crawler.URL.Strings(),
		Urls:    crawler.URLS.Strings(),
	}
	var m crawler.Manager
	_, _, err = m.GenerateManifest(destination, config)
}
