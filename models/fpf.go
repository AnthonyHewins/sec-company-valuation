package models

import (
	"fmt"
)

const (
	FY = iota
	Q1 = iota
	Q2 = iota
	Q3 = iota
	Q4 = iota
	H1 = iota
	H2 = iota
	M9 = iota
	T1 = iota
	T2 = iota
	T3 = iota
	M8 = iota
	CY = iota
)

// Fiscal period focus
type Fpf struct { code int }

func (f *Fpf) UnmarshalCSV(csv string) (err error) {
	switch csv {
	case "CY": f.code = CY
	case "FY": f.code = FY
	case "Q1": f.code = Q1
	case "Q2": f.code = Q2
	case "Q3": f.code = Q3
	case "Q4": f.code = Q4
	case "H1": f.code = H1
	case "H2": f.code = H2
	case "T1": f.code = T1
	case "T2": f.code = T2
	case "T3": f.code = T3
	case "M8": f.code = M8
	case "M9": f.code = M9
	default: return fmt.Errorf("unexpected code: %v", csv)
	}

	return nil
}
