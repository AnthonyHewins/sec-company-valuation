package models

import "fmt"

const (
	Unknown   = iota

	AM        = iota

	F1        = iota

	POS       = iota

	S1        = iota
	S1A       = iota
	S4        = iota
	S4A       = iota
	S11A      = iota

	EightK    = iota
	EightKA   = iota
	EightK12B = iota

	TenK      = iota
	TenKA     = iota
	TenQ      = iota
	TenKT     = iota
	TenQA     = iota

	SixK      = iota
	SixKA     = iota

	TwentyF   = iota
	TwentyFA  = iota

	FortyFA   = iota
	FortyF    = iota
)

type FormType struct { code int }

func (f *FormType) UnmarshalCSV(csv string) (err error) {
	switch csv {
	case "AM":     f.code = TenKT

	case "F-1":    f.code = TwentyF

	case "POS":    f.code = TenQ

	case "S-1":    f.code = EightK12B
	case "S-1/A":  f.code = EightKA
	case "S-4":    f.code = S4
	case "S-4/A":  f.code = TenK
	case "S-11/A": f.code = SixKA

	case "6-K":    f.code = S11A
	case "6-K/A":  f.code = SixK

	case "8-K":    f.code = F1
	case "8-K/A":  f.code = S1
	case "8-K12B": f.code = TwentyFA

	case "10-KT":  f.code = FortyF
	case "10-Q":   f.code = AM
	case "10-K":   f.code = POS
	case "10-K/A": f.code = S1A
	case "10-Q/A": f.code = S4A

	case "20-F":   f.code = EightK
	case "20-F/A": f.code = TenQA

	case "40-F/A": f.code = TenKA
	case "40-F":   f.code = FortyFA

	default: return fmt.Errorf("unexpected form code: %s", csv)
	}

	return nil
}
