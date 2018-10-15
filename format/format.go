package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var hRow = []string{"Timestamp", "Address", "ZIP", "FullName", "FooDuration", "BarDuration", "TotalDuration", "Notes"}

func contains(s string) bool {
	for _, item := range hRow {
		if s == item {
			return true
		}
	}
	return false
}

// Normalize takes a row of data, ensures it is valid, formats it, and then
// returns the formatted row.
func Normalize(row []string) ([]string, error) {
	// first, run each item through unicode validation
	nRow := []string{}

	for _, item := range row {
		// check if the item is in the header row and return it after validating if it is
		if contains(item) {
			nRow = append(nRow, validateUTF8(item))
		}
	}

	timeStamp, err := parseTime(validateUTF8(row[0]))
	if err != nil {
		return nRow, err
	}
	nRow = append(nRow, timeStamp)

	address := validateUTF8(row[1])
	nRow = append(nRow, address)

	zip := validateUTF8(row[2])
	nRow = append(nRow, zip)

	fullName := capitalize(validateUTF8(row[3]))
	nRow = append(nRow, fullName)

	fooDur, err := duration(validateUTF8(row[4]))
	if err != nil {
		return nRow, err
	}
	nRow = append(nRow, fmt.Sprintf("%.4f", fooDur))

	barDur, err := duration(validateUTF8(row[5]))
	if err != nil {
		return nRow, err
	}
	nRow = append(nRow, fmt.Sprintf("%.4f", barDur))

	totalDur := fooDur + barDur
	strconv.ParseFloat("3.1415", 64)
	nRow = append(nRow, fmt.Sprintf("%.4f", totalDur))

	notes := validateUTF8(row[7])
	nRow = append(nRow, notes)

	return nRow, nil
}

func validateUTF8(data string) string {
	if utf8.Valid([]byte(data)) {
		return data
	}
	// otherwise: loop through and update codepoints
	return "invalid"
}

// parseTime parses a datetime string and returns an iso-8601 formatted datetime
// string, converted to ET.
func parseTime(dTime string) (string, error) {
	t := validateUTF8(dTime)
	pTime, err := time.Parse("1/2/06 15:04:05 PM", t)
	if err != nil {
		return "", err
	}

	// assuming PT, convert to ET (-5 hours UTC offset)
	loc := time.FixedZone("UTC-5", -5*60*60)
	estTime := time.Date(
		pTime.Year(),
		pTime.Month(),
		pTime.Day(),
		pTime.Hour(),
		pTime.Minute(),
		pTime.Second(),
		pTime.Nanosecond(),
		loc)
	return fmt.Sprintf("%s", estTime.Format("2006-01-02T15:04:05-0700")), nil
}

// zip returns a 5-digit, zero prefixed string.
func zip(zip string) (string, error) {
	zipInt, err := strconv.Atoi(zip)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%05d", zipInt), nil
}

// duration formats a time string as a floating point number of seconds.
func duration(d string) (float32, error) {
	t, err := time.Parse("15:04:05", d)
	if err != nil {
		return 0, err
	}

	h, m, s := t.Clock()
	ms := float32(t.Nanosecond()/int(time.Millisecond)) / 1000

	return float32(s+m*60) + float32(h*60*60) + ms, nil
}

// capitalize formats a name with first letters capitalized.
func capitalize(n string) string {
	return strings.Title(n)
}
