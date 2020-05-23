package models

import (
	"fmt"
)

const (
	BalanceSheet            = iota
	IncomeStatement         = iota
	CashFlow                = iota
	Equity                  = iota
	ComprehensiveIncome     = iota
	UnclassifiableStatement = iota
)

type Statement struct { code int }

func (s *Statement) UnmarshalCSV(csv string) error {
	switch csv {
	case "BS": s.code = BalanceSheet
	case "IS": s.code = IncomeStatement
	case "CF": s.code = CashFlow
	case "EQ": s.code = Equity
	case "CI": s.code = ComprehensiveIncome
	case "UN": s.code = UnclassifiableStatement
	default: return fmt.Errorf("unrecognized statement code: %s", csv)
	}

	return nil
}
