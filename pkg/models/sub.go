package models

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type SubCsv struct {
	Subs []Sub
}

type Sub struct {
	Adsh       string  `csv:"adsh"`
	Cik        string  `csv:"cik"`
	Name       string  `csv:"name"`
	// Sic        *string  `csv:"-"`
	// Countryba  *string  `csv:"-"`
	// Stprba     *string  `csv:"-"`
	// Cityba     *string  `csv:"-"`
	// Zipba      *string  `csv:"-"`
	// Bas1       *string  `csv:"-"`
	// Bas2       *string  `csv:"-"`
	// Baph       *string  `csv:"-"`
	// Countryma  *string  `csv:"-"`
	// Stprma     *string  `csv:"-"`
	// Cityma     *string  `csv:"-"`
	// Zipma      *string  `csv:"-"`
	// Mas1       *string  `csv:"-"`
	// Mas2       *string  `csv:"-"`
	// Countryinc *string  `csv:"-"`
	// Stprinc    *string  `csv:"-"`
	// Ein        *string  `csv:"-"`
	Former     string  `csv:"former"`
	Changed    string  `csv:"changed"`
	Afs        string  `csv:"afs"`
	Wksi       bool    `csv:"wksi"`
	Fye        string  `csv:"fye"`
	Form       FormType `csv:"form"`
	Period     time.Time     `csv:"period"`
	Fy         Year     `csv:"fy"`
	Fp         string  `csv:"fp"` // TODO probably needs to be FPF
	Filed      time.Time     `csv:"filed"`
	Accepted   string  `csv:"accepted"`
	Prevrpt    bool     `csv:"prevrpt"`
	Detail     bool     `csv:"detail"`
	Instance   string  `csv:"instance"`
	Nciks      int     `csv:"nciks"`
	Aciks      string  `csv:"aciks"`
}

func parseSub(line []string) (*Sub, error) {
	if (len(line) != 36) {
		return nil, csv.ErrFieldCount
	}

	wksi, err := strconv.ParseBool(line[23])
	if err != nil {
		return nil, fmt.Errorf("Error parsing Wksi: %v", err)
	}

	form, err := parseFormType(line[25])
	if err != nil {
		return nil, fmt.Errorf("Error parsing form: %v", err)
	}

	period, err := parseDate(line[26])
	if err != nil {
		return nil, fmt.Errorf("Error parsing period: %v", err)
	}

	fy, err := strconv.ParseUint(line[27], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Error parsing fy: %v", err)
	}

	filed, err := parseDate(line[29])
	if err != nil {
		return nil, fmt.Errorf("Error parsing filed: %v", err)
	}

	prevrpt, err := strconv.ParseBool(line[31])
	if err != nil {
		return nil, fmt.Errorf("Error parsing prevrpt: %v", err)
	}

	detail, err := strconv.ParseBool(line[32])
	if err != nil {
		return nil, fmt.Errorf("Error parsing detail: %v", err)
	}

	nciks, err := strconv.ParseInt(line[34], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Error parsing nciks: %v", err)
	}

	return &Sub{
		Adsh: line[0],
		Cik: line[1],
		Name: line[2],
		// ...unused fields
		Former: line[20],
		Changed: line[21],
		Afs: line[22],
		Wksi: wksi,
		Fye: line[24],
		Form: form,
		Period: *period,
		Fy: Year(fy),
		Fp: line[28],
		Filed: *filed,
		Accepted: line[30],
		Prevrpt: prevrpt,
		Detail: detail,
		Instance: line[33],
		Nciks: int(nciks),
		Aciks: line[35],
	}, nil
}

func (s *Sub) UnmarshalSub(wg *sync.WaitGroup) {

}
