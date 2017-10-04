package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Zip is a struct that contains selected data
// from each record in zips.csv as fields.
type Zip struct {
	Code  string `json:"code,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// ZipSlice is a slice that contains many Zip.
type ZipSlice []*Zip

// ZipIndex is a map that maps
type ZipIndex map[string]ZipSlice

// LoadZips loads a given .csv file and returns a ZipSlice.
func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	// Reader will only one line at a time.
	reader := csv.NewReader(f)
	// Skip the first row in zips.csv
	// because they are header info.
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	// Pre-allocating the underlying array to be 43000.
	zips := make(ZipSlice, 0, 43000)
	for {
		// fields is an array that contains all tokens in one line.
		fields, err := reader.Read()

		// Jump out of "while" loop when the reader either
		// reaches the end of file or an error occur.
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}

		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}

		zips = append(zips, z)
	}
}
