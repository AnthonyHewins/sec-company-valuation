package models

type Pre struct{
	Adsh    string `csv:"adsh"`
	Report  int    `csv:"report"`
	Line    int    `csv:"line"`
	Stmt    string `csv:"stmt"`
	Inpth   bool   `csv:"inpth"`
	Rfile   string `csv:"rfile"`
	Tag     string `csv:"tag"`
	Version string `csv:"version"`
	Plabel  string `csv:"plabel"`
}
