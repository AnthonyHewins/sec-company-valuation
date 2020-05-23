package models

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestParseSub(t *testing.T) {
	tests := []struct {
		name    string
		line []string
		want    *Sub
		wantErr error
	}{
		{
			"Not enough lines, 35 instead of 36",
			[]string{/* "Adsh", */ "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "Wksi", "Fye", "Form", "Period", "Fy", "Fp", "Filed", "Accepted", "Prevrpt", "Detail", "Instance", "Nciks", "Aciks",},
			nil,
			csv.ErrFieldCount,
		},
		{
			"Too many lines, 37 instead of 36",
			[]string{"extra", "Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "Wksi", "Fye", "Form", "Period", "Fy", "Fp", "Filed", "Accepted", "Prevrpt", "Detail", "Instance", "Nciks", "Aciks",},
			nil,
			csv.ErrFieldCount,
		},
		{
			"Wksi must be bool, which is 1 or 0",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "notbool", "Fye", "10-Q", "Period", "Fy", "Fp", "20101010", "Accepted", "true", "true", "Instance", "10", "Aciks",},
			nil,
			fmt.Errorf("Error parsing Wksi: strconv.ParseBool: parsing \"notbool\": invalid syntax"),
		},
		{
			"Form must be one of the FormTypes",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "Not a form", "Period", "Fy", "Fp", "20101010", "Accepted", "true", "true", "Instance", "10", "Aciks",},
			nil,
			fmt.Errorf("Error parsing form: unexpected form code: Not a form"),
		},
		{
			"Period must be date",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "notdate", "Fy", "Fp", "20101010", "Accepted", "true", "true", "Instance", "10", "Aciks",},
			nil,
			fmt.Errorf("Error parsing period: tried parsing two ways, got error1: parsing time \"notdate\" as \"060201\": cannot parse \"notdate\" as \"06\" || error2: parsing time \"notdate\" as \"20060102\": cannot parse \"notdate\" as \"2006\" || on notdate"),
		},
		{
			"Fy must be a year",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "Not a year", "Fp", "Filed", "Accepted", "Prevrpt", "Detail", "Instance", "Nciks", "Aciks",},
			nil,
			fmt.Errorf("Error parsing fy: strconv.ParseUint: parsing \"Not a year\": invalid syntax"),
		},
		{
			"Filed must be Date",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "2010", "2010", "not date", "Accepted", "Prevrpt", "Detail", "Instance", "Nciks", "Aciks",},
			nil,
			fmt.Errorf("Error parsing filed: tried parsing two ways, got error1: parsing time \"not date\" as \"060201\": cannot parse \"not date\" as \"06\" || error2: parsing time \"not date\" as \"20060102\": cannot parse \"not date\" as \"2006\" || on not date"),
		},
		{
			"Prevrpt must be bool",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "2010", "2010", "20101010", "Accepted", "notbool", "Detail", "Instance", "Nciks", "Aciks",},
			nil,
			fmt.Errorf("Error parsing prevrpt: strconv.ParseBool: parsing \"notbool\": invalid syntax"),
		},
		{
			"Detail must be bool",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "2010", "2010", "20060131", "Accepted", "true", "notbool", "Instance", "Nciks", "Aciks",},
			nil,
			fmt.Errorf("Error parsing detail: strconv.ParseBool: parsing \"notbool\": invalid syntax"),
		},
		{
			"Nciks must be int",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "2010", "2010", "20060131", "Accepted", "true", "true", "Instance", "notint", "Aciks",},
			nil,
			fmt.Errorf("Error parsing nciks: strconv.ParseInt: parsing \"notint\": invalid syntax"),
		},
		{
			"Valid parse",
			[]string{"Adsh", "Cik", "Name", "Sic", "Countryba", "Stprba", "Cityba", "Zipba", "Bas1", "Bas2", "Baph", "Countryma", "Stprma", "Cityma", "Zipma", "Mas1", "Mas2", "Countryin", "Stprinc", "Ein", "Former", "Changed", "Afs", "true", "Fye", "10-Q", "20101010", "2010", "2010", "20060131", "Accepted", "true", "true", "Instance", "0", "Aciks",},
			&Sub{
				Adsh: "Adsh",
				Cik: "Cik",
				Name: "Name",
				// ...unused fields
				Former: "Former",
				Changed: "Changed",
				Afs: "Afs",
				Wksi: true,
				Fye: "Fye",
				Form: 1,
				Period: time.Date(2010, 10, 10, 0, 0, 0, 0, time.UTC),
				Fy: 2010,
				Fp: "2010",
				Filed: time.Date(2006, 01, 31, 0, 0, 0, 0, time.UTC),
				Accepted: "Accepted",
				Prevrpt: true,
				Detail: true,
				Instance: "Instance",
				Nciks: 0,
				Aciks: "Aciks",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSub(tt.line)
			if err != tt.wantErr && (err == nil || tt.wantErr == nil || tt.wantErr.Error() != err.Error()) {
				t.Errorf("FromCsv() error = %v, want %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}
