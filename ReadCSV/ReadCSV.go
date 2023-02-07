
package main

import (
"encoding/csv"
import (
"encoding/csv"
"fmt"
"os"
)

type record struct {
	Name string
	Age int
	City string
}

func main() {
	// Open the CSV file
	file, err := os.Open("records.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()


	// Read the contents of the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Parse the records into a slice of structs
	var data []record
	for i, record := range records {
		if i == 0 {
			// Skip the header row
			continue
		}

		// Parse the Age column as an int
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println("Error parsing Age:", err)
			return
		}

		data = append(data, record{
			Name: record[0],
			Age:  age,
			City: record[2],
		})
	}

	// Output the data as a table
	fmt.Println("Name\tAge\tCity")
	fmt.Println("----\t---\t----")
	for _, record := range data {
		fmt.Printf("%s\t%d\t%s\n", record.Name, record.Age, record.City)
	}
}

