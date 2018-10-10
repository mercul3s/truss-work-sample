package format

import "fmt"

type rowData struct {
	timestamp     string
	address       string
	zipcode       string
	fullName      string
	fooDuration   string
	barDuration   string
	totalDuration string
	notes         string
}

// normalize the csv:
// 	address: unicode normalization but contain commas
//  notes - leave as is
// chain a bunch of funcs in here
func Normalize(row []string) ([]string, error) {
	for _, item := range row {
		fmt.Println(item)
	}
	return row, nil
}

//	ensure valid utf-8
func ValidateUTF8() {
}

//	timestamp to iso-8601
//	converted to ET
func Time() {
}

//	zip codes formatted to 5 digits with 0 prefix if they are less
func Zip() {
}

// 	FooDuration and BarDuration to be floating point seconds
// 	TotalDuration = sum of Foo and Bar duration
func Duration() {
}
