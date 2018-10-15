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

	reader := csv.NewReader(bufio.NewReader(csvIn))
	writer := csv.NewWriter(csvOut)

	// read through the data until EOF, then write and close the file.
	for {
		row, err := reader.Read()
		if err == io.EOF {
			// flush the writer and exit
			fmt.Printf("finished normalizing input file, writing to %s\n", normFileName)
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
