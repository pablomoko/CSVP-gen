package main

import (
	CSVPgen "CSVPgen/internal/csv"
	Processor "CSVPgen/internal/processor"
	Processors "CSVPgen/internal/processor/processors"

	"fmt"
)

func main() {

	filePath := "data/input/data_test.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	columnNames, rows, err := CSVPgen.ReadCSV(file, 0)

	outputFilePath := "data/output/data_test_processed.csv"
	outputFile, err := CSVPgen.CreateCSV(outputFilePath)

	if err != nil {
		fmt.Println("Error creando archivo de salida: ", err)
		return
	}

	columnProcessors := Processor.ColumnProcessorMap{
		"Scan Rate": Processors.DivisorProcessor{Divisor: 2},
	}

	processedRows, err := Processor.ProcessRows(columnNames, rows, columnProcessors)

	if err := CSVPgen.WriteCSV(outputFile, processedRows, columnNames); err != nil {
		fmt.Println("Error escribiendo archivo CSV:", err)
		return
	}

	fmt.Println("Archivo CSV procesado y guardado en:", outputFilePath)

}
