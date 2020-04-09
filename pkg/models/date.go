package models

import (
	"fmt"
	"time"
)

type Year byte

type Date struct { time.Time }

func (d *Date) UnmarshalCSV(csv string) error {
	var err1, err2 error

	if d.Time, err1 = time.Parse("060201", csv); err1 != nil {
		if d.Time, err2 = time.Parse("20060102", csv); err2 != nil {
			return fmt.Errorf("tried parsing two ways, got error1: %v || error2: %v || on %s", err1, err2, csv)
		}
	}

	return nil
}
