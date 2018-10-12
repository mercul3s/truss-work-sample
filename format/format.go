package format

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

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
//  notes: validate utf-8 (replace invalid chars with utf-8 replacement char)
//  uppercase names
// chain a bunch of funcs in here
func Normalize(row []string) ([]string, error) {
	// first, run each item through unicode validation
	fmt.Println(row)
	for _, item := range row {
		parsedTime, err := Time(item)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(parsedTime)
	}
	return row, nil
}

//	ensure valid utf-8
//  Go is utf-8 encoded, so I think all of this should just work?
func ValidateUTF8(data string) string {
	if utf8.Valid([]byte(data)) {
		return data
	}
	return "invalid"
}

//	timestamp to iso-8601
//	converted to ET
func Time(t string) (string, error) {
	fmt.Println(t)
	parsedTime, err := time.Parse("1/2/06 15:04:05 PM", t)
	if err != nil {
		return "", err
	}
	parsedTime.Add(-4 * 60 * 60)
	fmt.Println("Parsed time is:", parsedTime)
	// assuming PT, convert to ET (+3 hours)
	return fmt.Sprintf("%s", parsedTime), err
}

//	zip codes formatted to 5 digits with 0 prefix if they are less
func Zip(zip string) (string, error) {
	zipInt, err := strconv.Atoi(zip)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%05d", zipInt), nil
}

// 	FooDuration and BarDuration to be floating point seconds
// 	TotalDuration = sum of Foo and Bar duration
func Duration(d string) (string, error) {
	t, err := time.Parse("15:04:05", d)
	if err != nil {
		return "", err
	}

	h, m, s := t.Clock()
	ms := t.Nanosecond() / int(time.Millisecond)
	return fmt.Sprintf("%d.%d", (s + m*60 + h*60*60), ms), nil
}
