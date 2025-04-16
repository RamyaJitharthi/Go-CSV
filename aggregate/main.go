package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// records := [][]string {
	// 	 	 	{"companyName","sales", "profit"},
	// 			{"Hill", "$3000", "$1000"},
	// 			{"Sea View", "$2000", "$800"},
	// 			{"Mountain", "$2500", "$1200"},
	// }

	// f, err := os.Create("sales.csv")
	
	// if err != nil {
	// 	log.Fatalln("Failed to open file ", err)
	// }
	
	// defer f.Close()

	// w := csv.NewWriter(f)
	// defer w.Flush()

	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatalln("Error writng record to file", err)
	// 	}

	// }

	f, err := os.Open("sales.csv") 
	if err != nil {
		log.Fatalln("Failed to open File", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Failed to to read csv", err)
	}

	// var totalSales,totalProfit float64
	// var count int

	salesTotals := make(map[string]float64)
	profitTotals := make(map[string]float64)
	entryCounts := make(map[string]int)

	for i, record := range records {
		if i== 0 {
			continue
		}

		company := record[1]
		salesStr := strings.TrimPrefix(record[2], "$")
		profitStr := strings.TrimPrefix(record[3], "$")

		sales, err1 := strconv.ParseFloat(salesStr, 64)
		profit, err2 := strconv.ParseFloat(profitStr, 64)

		if err1 != nil || err2 != nil {
			log.Printf("Skipping invalid record at line %d",i+1)
			continue
		}

		// totalSales += sales
		// totalProfit += profit
		// count++

		salesTotals[company] += sales
		profitTotals[company] += profit
		entryCounts[company]++
	}

	// fmt.Printf("Total Sales: $%.2f\n", totalSales)
	// fmt.Printf("Total Profit: $%.2f\n", totalProfit)

	// if count > 0 {
	// 	fmt.Printf("Average Daily sales: $%.2f\n", totalSales/float64(count ))
	// 	fmt.Printf("Average Daily profit: $%.2f\n", totalProfit/float64(count ))
	// }

	fmt.Println("Company-wise Totals and Averages:")
	for company, totalSales := range salesTotals {
		totalProfit := profitTotals[company]
		count := entryCounts[company]

		fmt.Printf("\nCompany: %s\n", company)
		fmt.Printf("  Total Sales: $%.2f\n", totalSales)
		fmt.Printf("  Total Profit: $%.2f\n", totalProfit)
		fmt.Printf("  Average Daily Sales: $%.2f\n", totalSales/float64(count))
		fmt.Printf("  Average Daily Profit: $%.2f\n", totalProfit/float64(count))
	}
}
