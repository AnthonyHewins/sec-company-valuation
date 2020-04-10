package models

type Num struct{
	Adsh     string    `csv:"adsh"`
	Tag      string    `csv:"tag"`
	Version  string    `csv:"version"`
	Ddate    Date      `csv:"ddate"`
	Qtrs     int       `csv:"qtrs"`
	Uom      string    `csv:"uom"`
	Coreg    string       `csv:"coreg"` // TODO needs to be int
	Value    int     `csv:"value"`
	Footnote string    `csv:"footnote"`
}
