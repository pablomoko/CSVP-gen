package main

import (
	CSVPgen "CSVPgen/internal/csv"
	Processors "CSVPgen/internal/processors"
	"fmt"
)

func main() {

	filePath := "data/input/data_test.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	columnProcessors := CSVPgen.ColumnProcessorMap{
		"Scan Rate": Processors.DivisorProcessor{Divisor: 2},
	}

	columnNames, processedRows, err := CSVPgen.ReadCSV(file, columnProcessors, 0)

	outputFilePath := "data/output/data_test_processed.csv"
	outputFile, err := CSVPgen.CreateCSV(outputFilePath)

	if err != nil {
		fmt.Println("Error creando archivo de salida: ", err)
		return
	}

	if err := CSVPgen.WriteCSV(outputFile, processedRows, columnNames); err != nil {
		fmt.Println("Error escribiendo archivo CSV:", err)
		return
	}

	fmt.Println("Archivo CSV procesado y guardado en:", outputFilePath)

}
