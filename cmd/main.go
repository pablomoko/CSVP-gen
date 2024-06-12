package main

import (
	CSVPgen "CSVPgen/internal/csv"
	Generator "CSVPgen/internal/generator"
	Generators "CSVPgen/internal/generator/generators"
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

	columnProcessors := Processor.ColumnProcessorMap{
		"Scan Rate": Processors.DivisorProcessor{Divisor: 2},
	}

	processedRows, err := Processor.ProcessRows(rows, columnProcessors)

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

	test1()

}

func test1() {
	filePath := "data/input/data_test_2.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	_, rows, err := CSVPgen.ReadCSV(file, 0)

	sumGenerator := &Generators.SumColumnGenerator{
		SourceColumns: []string{"Scan Rate", "Address"}, // Definir las columnas de origen para la suma
		NewColumnName: "SumColumn",                      // Definir el nombre de la nueva columna
	}

	// Crear un mapa de generadores de columnas
	columnGenerators := Generator.ColumnGeneratorMap{
		"SumColumn": sumGenerator, // Agregar el generador de suma al mapa
	}

	finalRows, err := Generator.ProcessRowsWithGenerators(rows, columnGenerators)

	fmt.Println(finalRows)

}
