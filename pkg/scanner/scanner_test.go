package scanner

import "testing"

func TestScan(t *testing.T) {
	conf := &Config{
		ScanType: EXCEL,
		Path:     "../../testdata/TestBook.xlsm",
	}
	s := Scanner{Conf: *conf}
	err := s.Crawl()
	if err != nil {
		return
	}
}
