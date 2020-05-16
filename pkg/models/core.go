package models

type OptionCsvReader struct {
	Err error
	Ok interface{}
}

type CsvReadable interface {
	FromCsv(line []string) (interface{}, error)
}
