package main

import (
	CSVPgen "CSVPgen/internal/csv"
	Processors "CSVPgen/internal/processors"
	"fmt"
)

func main() {
	fmt.Println("CSVP&gen")

	filePath := "data/input/data_test.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	columnProcessors := CSVPgen.ColumnProcessorMap{
		"Scan Rate": Processors.DivisorProcessor{Divisor: 2},
	}

	fmt.Println(CSVPgen.ReadCSV(file, columnProcessors))

}
