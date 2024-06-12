package main

import (
	CSVPgen "CSVPgen/internal/csv"
	Generator "CSVPgen/internal/generator"
	ColGenerators "CSVPgen/internal/generator/generators/column"
	MetaDataGenerator "CSVPgen/internal/generator/generators/metadata"
	Processor "CSVPgen/internal/processor"
	Processors "CSVPgen/internal/processor/processors"
	"CSVPgen/internal/types"

	"fmt"
)

func main() {

	filePath := "data/input/data_test.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	_, rows, err := CSVPgen.ReadCSV(file, 0)

	columnProcessors := Processor.ColumnProcessorMap{
		"Age": Processors.DivisorProcessor{Divisor: 2},
	}

	processedRows, err := Processor.ProcessRows(rows, columnProcessors)

	metadata := &types.CSVMetadata{
		FileName:   "test_metadata_1",
		NumRows:    0,
		NumColumns: 0,
	}

	outputFilePath := "data/output/" + metadata.FileName + ".csv"
	outputFile, err := CSVPgen.CreateCSV(outputFilePath)

	if err != nil {
		fmt.Println("Error creando archivo de salida: ", err)
		return
	}

	if err := CSVPgen.WriteCSV(outputFile, processedRows); err != nil {
		fmt.Println("Error escribiendo archivo CSV:", err)
		return
	}

	fmt.Println("Archivo CSV procesado y guardado en:", outputFilePath)

	test1()

}

func test1() {
	filePath := "data/input/data_test.csv"
	file, err := CSVPgen.OpenCSV(filePath)

	if err != nil {
		fmt.Println(err)
	}

	_, rows, err := CSVPgen.ReadCSV(file, 0)

	sumGenerator := &ColGenerators.SumColumnGenerator{
		SourceColumns: []string{"Age", "Dog's Age"}, // Definir las columnas de origen para la suma
		NewColumnName: "SumColumn",                  // Definir el nombre de la nueva columna
	}

	metadata := &types.CSVMetadata{
		FileName:   "test_metadata_2",
		NumRows:    0,
		NumColumns: 0,
	}

	metadataGenerator := &MetaDataGenerator.MetadataGenerator{
		Metadata: metadata,
	}

	columnGenerators := Generator.ColumnGeneratorMap{
		"SumColumn": sumGenerator,
		"Metadata":  metadataGenerator, // Agregar el generador de metadatos al mapa
	}

	finalRows, err := Generator.ProcessRowsWithGenerators(rows, columnGenerators)

	outputFilePath := "data/output/" + metadata.FileName + ".csv"
	outputFile, err := CSVPgen.CreateCSV(outputFilePath)

	if err := CSVPgen.WriteCSV(outputFile, finalRows); err != nil {
		fmt.Println("Error escribiendo archivo CSV:", err)
		return
	}

}
