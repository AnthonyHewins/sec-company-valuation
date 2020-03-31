package fetch

import (
	"fmt"
	"time"
	"bytes"

	"net/http"
	"io/ioutil"
	"archive/zip"
)

const (
	FINDATA = "https://www.sec.gov/files/dera/data/financial-statement-data-sets"
)

func SecURL(year, quarter int) (error) {
	if quarter > 4 || quarter < 1 {
		return fmt.Errorf("Quarter must be between 1 and 4")
	}

	if year < 2009 || year > time.Now().Year() {
		return fmt.Errorf("Year must be greater than 2009 (SEC's records start then) and can't be in the future")
	}

	return URL(fmt.Sprintf("%s/%dq%d.zip", FINDATA, year, quarter))
}

func URL(url string) (error) {
	resp, err := http.Get(url)
	if err != nil { return err }

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil { return err }

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil { return err }

	return unzip(r)
}

func ZIP(path string) (error) {
	readerCloser, err := zip.OpenReader(path)
	if err != nil { return err }

	defer readerCloser.Close()

	return unzip(&readerCloser.Reader)
}

func unzip(r *zip.Reader) (error) {
	var sub, tag, pre, num []byte

	for _, file := range r.File {
		// TODO handle serialization of each file into structs
		// then throw everything else away
		var err error

		switch file.Name {
		case "sub.txt":
			sub, err = readZipFile(file)
		case "tag.txt":
			tag, err = readZipFile(file)
		case "pre.txt":
			pre, err = readZipFile(file)
		case "num.txt":
			num, err = readZipFile(file)
		default:
			// skip file
		}

		if err != nil { return err }
	}

	if len(sub) == 0 || len(tag) == 0 || len(pre) == 0 || len(num) == 0 {
		return fmt.Errorf("ZIP file is missing one of the files for DCF: need {sub, tag, pre, num}.txt")
	}

	return nil
}

func readZipFile(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil { return nil, err }

	defer f.Close()

	return ioutil.ReadAll(f)
}
