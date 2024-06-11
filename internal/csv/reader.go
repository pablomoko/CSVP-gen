package csv

import (
	"encoding/csv"
	"os"
)

type StructField struct {
	Name  string
	Value string
}

type Row struct {
	Fields []StructField
}

func ReadCSV(file *os.File, columnProcessors ColumnProcessorMap, headerRow int) ([]string, []Row, error) {
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	columnNames := rows[headerRow]

	var processedRows []Row

	for _, row := range rows[headerRow+1:] {
		processedRows = append(processedRows, ProcessRow(row, columnNames, columnProcessors))
	}

	return columnNames, processedRows, nil
}
