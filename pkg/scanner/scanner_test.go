package scanner

import "testing"

func TestScan(t *testing.T) {
	var s Scanner
	s.Scan("../../testdata/TestBook.xlsm", EXCEL)
}
