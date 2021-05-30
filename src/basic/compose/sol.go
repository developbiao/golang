package main

import "fmt"

type sol int

type report struct {
	sol
	location
	temperature
}

type temperature struct {
	hight, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (s sol) days(s2 sol) int {
	fmt.Println("Call sol days()")
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{sol: 15, location: location{-33.33, 29.87}}

	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(365))
	fmt.Println(report.lat)
}
