package fetch

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"archive/zip"
	"io/ioutil"
	"net/http"

	"github.com/gocarina/gocsv"

	"github.com/AnthonyHewins/sec-company-valuation/pkg/models"
)

const (
	DeraUrl = "https://www.sec.gov/files/dera/data/financial-statement-data-sets"
)

var (
	ErrQuarter = fmt.Errorf("Quarter must be between 1 and 4")
	ErrYear = fmt.Errorf("Year must be greater than 2009 (SEC's records start then) and can't be in the future")
)

func SecUrl(year, quarter int) error {
	if quarter > 4 || quarter < 1 {
		return ErrQuarter
	}

	if year < 2009 || year > time.Now().Year() {
		return ErrYear
	}

	return Url(fmt.Sprintf("%s/%dq%d.zip", DeraUrl, year, quarter))
}

func Url(url string) error {
	resp, err := http.Get(url)
	if err != nil { return err }

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil { return err }

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil { return err }

	return unzip(r)
}

func Zip(path string) error {
	readerCloser, err := zip.OpenReader(path)
	if err != nil { return err }

	defer readerCloser.Close()

	return unzip(&readerCloser.Reader)
}

func unzip(r *zip.Reader) error {
	var tag []models.Tag
	var sub []models.Sub
	var pre []models.Pre
	var num []models.Num

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		r.LazyQuotes = true
		return r
	})

	for _, file := range r.File {
		// TODO handle serialization of each file into structs
		// then throw everything else away
		var err error

		switch file.Name {
		case "sub.txt":
			fmt.Println("sub")
			err = readZipFile(file, func(r io.Reader) (error) {
				return gocsv.Unmarshal(r, &sub)
			})
		case "tag.txt":
			fmt.Println("tag.txt")
			err = readZipFile(file, func(r io.Reader) (error) {
				return gocsv.Unmarshal(r, &tag)
			})
		case "pre.txt":
			fmt.Println("pre.txt")
			err = readZipFile(file, func(r io.Reader) (error) {
				return gocsv.Unmarshal(r, &pre)
			})
		case "num.txt":
			fmt.Println("num.txt")
			err = readZipFile(file, func(r io.Reader) (error) {
				return gocsv.Unmarshal(r, &num)
			})
		default:
			// skip file
		}

		if err != nil { return err }
	}

	if len(sub) == 0 || len(tag) == 0 || len(pre) == 0 || len(num) == 0 {
		return fmt.Errorf("ZIP file is missing one of the files for DCF, or it had length 0: need {sub, tag, pre, num}.txt")
	}

	return nil
}

func readZipFile(file *zip.File, fn func(r io.Reader) (error)) error {
	fileReader, err := file.Open()
	if err != nil { return err }

	defer fileReader.Close()

	return fn(fileReader)
}
