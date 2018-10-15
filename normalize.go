package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/mercul3s/work_sample/format"
)

/*

Please write a tool that reads a CSV formatted file on `stdin` and
emits a normalized CSV formatted file on `stdout`. Normalized, in this
case, means:

* The entire CSV is in the UTF-8 character set.
* The Timestamp column should be formatted in ISO-8601 format.
* The Timestamp column should be assumed to be in US/Pacific time;
  please convert it to US/Eastern.
* All ZIP codes should be formatted as 5 digits. If there are less
  than 5 digits, assume 0 as the prefix.
* All name columns should be converted to uppercase. There will be
  non-English names.
* The Address column should be passed through as is, except for
  Unicode validation. Please note there are commas in the Address
  field; your CSV parsing will need to take that into account. Commas
  will only be present inside a quoted string.
* The columns `FooDuration` and `BarDuration` are in HH:MM:SS.MS
  format (where MS is milliseconds); please convert them to a floating
  point seconds format.
* The column "TotalDuration" is filled with garbage data. For each
  row, please replace the value of TotalDuration with the sum of
  FooDuration and BarDuration.
* The column "Notes" is free form text input by end-users; please do
  not perform any transformations on this column. If there are invalid
  UTF-8 characters, please replace them with the Unicode Replacement
  Character.

You can assume that the input document is in UTF-8 and that any times
that are missing timezone information are in US/Pacific. If a
character is invalid, please replace it with the Unicode Replacement
Character. If that replacement makes data invalid (for example,
because it turns a date field into something unparseable), print a
warning to `stderr` and drop the row from your output.

You can assume that the sample data we provide will contain all date
and time format variants you will need to handle.

*/

var fileName string

func main() {
	// only take the second argument, which should be the filename
	// and only if it is provided
	if len(os.Args) == 1 {
		fmt.Println("Please provide a file name/path to normalize: './normalize sample.csv' or './normalize /path/to/file.csv'")
		os.Exit(1)
	}
	fileName := os.Args[1]

	csvIn, err := os.Open(fileName)
	if err != nil {
		handleFileError(err, fileName)
	}

	_, fName := path.Split(fileName)

	normFileName := fmt.Sprintf("normalized_" + fName)
	csvOut, err := os.Create(normFileName)
	if err != nil {
		handleFileError(err, fileName)
	}

	fmt.Println("CSV Output filename: ", normFileName)

	reader := csv.NewReader(bufio.NewReader(csvIn))
	writer := csv.NewWriter(csvOut)

	// read through the data until EOF, then write and close the file.
	for {
		row, err := reader.Read()
		if err == io.EOF {
			// flush the writer and exit
			writer.Flush()
			os.Exit(0)
		}
		if err != nil {
			handleFileError(err, fileName)
		}

		normRow, err := format.Normalize(row)
		if err != nil {
			fmt.Printf("Unable to normalize row %s due to error %s - skipping\n", row, err)
			continue
		}
		writer.Write(normRow)
	}
}

func handleFileError(err error, fileName string) {
	fmt.Printf("Error occurred while processing %s: %s\n", fileName, err)
	os.Exit(1)
}
