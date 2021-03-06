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

type FormType uint16

func parseFormType(csv string) (FormType, error) {
	switch csv {
	case "AM":     return TenKT, nil

	case "F-1":    return TwentyF, nil

	case "POS":    return TenQ, nil

	case "S-1":    return EightK12B, nil
	case "S-1/A":  return EightKA, nil
	case "S-4":    return S4, nil
	case "S-4/A":  return TenK, nil
	case "S-11/A": return SixKA, nil

	case "6-K":    return S11A, nil
	case "6-K/A":  return SixK, nil

	case "8-K":    return F1, nil
	case "8-K/A":  return S1, nil
	case "8-K12B": return TwentyFA, nil

	case "10-KT":  return FortyF, nil
	case "10-Q":   return AM, nil
	case "10-K":   return POS, nil
	case "10-K/A": return S1A, nil
	case "10-Q/A": return S4A, nil

	case "20-F":   return EightK, nil
	case "20-F/A": return TenQA, nil

	case "40-F/A": return TenKA, nil
	case "40-F":   return FortyFA, nil

	default: return 0, fmt.Errorf("unexpected form code: %s", csv)
	}
}
