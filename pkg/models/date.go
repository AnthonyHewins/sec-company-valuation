package models

import (
	"fmt"
	"time"
)

type Year byte

type Date struct {
	time.Time
}

func (d *Date) UnmarshalCSV(csv string) (err error) {
	d.Time, err = time.Parse("20060201", csv)
	fmt.Println(csv)
	fmt.Println(err)
	return err
}
