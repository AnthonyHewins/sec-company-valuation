package models

import (
	"fmt"
	"time"
)

type Year uint16

func parseDate(csv string) (*time.Time, error) {
	var d time.Time

	d, err1 := time.Parse("060201", csv)
	if err1 == nil {
		return &d, nil
	}

	d, err2 := time.Parse("20060102", csv)
	if err2 == nil {
		return &d, nil
	}

	return nil, fmt.Errorf("tried parsing two ways, got error1: %v || error2: %v || on %s", err1, err2, csv)
}
