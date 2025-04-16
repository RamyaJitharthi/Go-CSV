package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("new.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	allRecords, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read csv: %v", err)
	}

	if len(allRecords) == 0 {
		log.Fatal("CSV File is empty")
	}

	headers := allRecords[0]
	dataRows := allRecords[1:]

	recordCount := make(map[string]int)
	firstSeen := make(map[string]bool)
	rowMap := make(map[string][]string)

	uniqueRecords := [][]string{headers}
	duplicates := [][]string{{"lastname", "count"}}

	for _, row := range dataRows {
		key := strings.Join(row, "|")
		recordCount[key]++

		if !firstSeen[key] {
			uniqueRecords = append(uniqueRecords, row)
			firstSeen[key] = true
			rowMap[key] = row
		}
	}

	for key, count := range recordCount {
		if count > 1 {
			lastName := rowMap[key][0]
			duplicates = append(duplicates, []string{lastName, strconv.Itoa(count)})
		}
	}

	writeCSV("clean.csv", uniqueRecords)
	writeCSV("duplicates.csv", duplicates)
}

func writeCSV(filename string, records [][]string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filename, err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write to file %s: %v", filename, err)
		}
	}
}
