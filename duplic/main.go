package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
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

	uniqueMap := make(map[string]bool)
	uniqueRecords := [][]string{headers}

	valueCount := make(map[string]int)

	for _, row := range dataRows {
		rowKey := joinRow(row)

		if !uniqueMap[rowKey] {
			uniqueRecords = append(uniqueRecords, row)
			uniqueMap[rowKey] = true
		}

		for _, val := range row {
			valueCount[val]++
		}
	}

	duplicates := [][]string{{"value", "count"}}
	for val, count := range valueCount {
		if count > 1 {
			duplicates = append(duplicates, []string{val, strconv.Itoa(count)})
		}
	}

	writeCSV("clean.csv", uniqueRecords)
	writeCSV("duplicates.csv", duplicates)
}

func joinRow(row []string) string {
	return "\"" + joinWithDelimiter(row, "|") + "\""
}

func joinWithDelimiter(row []string, delim string) string {
	result := ""
	for i, v := range row {
		if i > 0 {
			result += delim
		}
		result += v
	}
	return result
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
