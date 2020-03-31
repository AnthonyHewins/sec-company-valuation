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
	FINDATA = "https://www.sec.gov/files/dera/data/financial-statement-data-sets/2019q4.zip"
)

func Fetch(year, quarter int) (error) {
	if quarter > 4 || quarter < 1 {
		return fmt.Errorf("Quarter must be between 1 and 4")
	}

	if year < 2009 || year > time.Now().Year() {
		return fmt.Errorf("Year must be greater than 2009 (SEC's records start then) and can't be in the future")
	}

	return downloadInPlace(year, quarter)
}

func downloadInPlace(year, quarter int) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%sq%s.zip", FINDATA, year, quarter))
	if err != nil { return nil, err }

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil, err }

	// TODO
	return nil, unzip(buf)
}

func unzip(buf []byte) (error) {
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil { return err }

	for _, reader := range r.File {
		// TODO handle serialization of each file into structs
		// then throw everything else away
		fmt.Println("%s", r.File)
		fmt.Println("read a file")
	}

	return nil
}
