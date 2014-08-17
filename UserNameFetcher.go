package main

import (
	"encoding/csv"
	"errors"
	"net/http"
)

const (
	URL = "http://docs.google.com/feeds/download/spreadsheets/Export?key=10eX2Uu23XUiIshBJhjlpeo4ZdAM--3HQPUHYwt62lY8&exportFormat=csv&gid=0"
)

func fetchUsernames() (result map[string]string, err error) {
	resp, err := http.Get(URL)
	if err != nil {
		err = errors.New("Could not fetch the URL")
		return
	}

	reader := csv.NewReader(resp.Body)

	// remove the first line with the headers
	reader.Read()

	records, err := reader.ReadAll()
	if err != nil {
		err = errors.New("Could not Parse the csv file")
		return
	}

	result = make(map[string]string)
	for i := range records {
		result[records[i][0]] = records[i][1]
	}

	return
}
