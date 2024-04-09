package crawler

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type WebCrawler struct {
	Conf Config
}

func (w WebCrawler) Crawl() error {
	//TODO implement me
	panic("implement me")
}

// Download manages direct call to url persisting the download over successive tries until timeout
func (w WebCrawler) Download(file *os.File, retries int) error {
	req, err := http.NewRequest("", w.Conf.Url, nil)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	local := fi.Size()
	if local > 0 {
		start := strconv.FormatInt(local, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}
	webClient := &http.Client{Timeout: 2 * time.Minute}
	res, err := webClient.Do(req)
	if err != nil && timeout(err) {
		if retries > 0 {
			return w.Download(file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errMsg := "http request unsuccessful, Status: %s"
		return fmt.Errorf(errMsg, res.Status)
	}
	if res.Header.Get("Accept-Range") != "bytes" {
		retries = 0
	}
	_, err = io.Copy(file, res.Body)
	if err != nil && timeout(err) {
		if retries > 0 {
			return w.Download(file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	return nil
}

func timeout(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if connErr, ok := err.Err.(net.Error); ok && connErr.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	connectionError := "closed network"
	if err != nil && strings.Contains(err.Error(), connectionError) {
		return true
	}
	return false
}
