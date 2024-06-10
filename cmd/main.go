package main

import (
	CSVPgen "CSVPgen/internal/csv"
	"fmt"
)

func main() {
	fmt.Println("CSVP&gen")

	filePath := "data/input/data_test.csv"

	fmt.Println(CSVPgen.ReadCSV(filePath))

}
