package main

import (
	"fmt"
	"time"
)

func main() {
	a := dayofyear(2012, 12, 31)
	fmt.Println(a)

}

func dayofyear(year int, month int, day int) int {

	m := time.Month(month)

	base := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	t := time.Date(year, m, day, 0, 0, 0, 0, time.UTC)

	d := t.Sub(base)

	return int(d.Hours()/24.0 + 1)

}
