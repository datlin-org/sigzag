package crawler

import (
	"strings"
	"testing"
)

func TestValidateExt(t *testing.T) {
	url := "https://randomurl.somewhere.zip"
	res := ValidateExt(url)
	urlComponents := strings.Split(url, "/")
	extComponents := strings.Split(urlComponents[len(urlComponents)-1], ".")
	ext := extComponents[len(extComponents)-1]
	v := []string{
		XLSX.Strings(),
		XLSB.Strings(),
		XLSM.Strings(),
		CSV.Strings(),
		ARFF.Strings(),
		IPYNB.Strings(),
		PARQUET.Strings(),
		ZIP.Strings(),
		BIN.Strings(),
		PDF.Strings(),
		GZ.Strings(),
		TXT.Strings(),
	}

	if !res {
		t.Errorf("Expected: known extension as suffix %v, Actual: %s", v, ext)
	}

	url = "https://randomurl.somewhere.zp"
	res = ValidateExt(url)
	urlComponents = strings.Split(url, "/")
	extComponents = strings.Split(urlComponents[len(urlComponents)-1], ".")
	ext = extComponents[len(extComponents)-1]
	if res {
		t.Errorf("Expected: known extension as suffix %v, Actual: %s", v, ext)
	}
}

func TestValidateUrl(t *testing.T) {
	url := "https://randomurl.n.uk"
	_, valid, err := ValidateUrl(url)
	if !valid {
		t.Errorf("Url validation error: %s", err.Error())
	}
}
