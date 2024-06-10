package csv

import (
	"encoding/csv"
	"os"
	//"path/filepath"
	//"strings"
)

// Campo de dato
type StructField struct {
	Name  string
	Value string
}

// Fila del csv
type Row struct {
	Fields []StructField
}

func ReadCSV(filepath string) ([]Row, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	columnNames := rows[0]

	var processedRows []Row

	for _, row := range rows[1:] {
		var fields []StructField
		for i, value := range row {
			fields = append(fields, StructField{Name: columnNames[i], Value: value})
		}
		processedRows = append(processedRows, Row{Fields: fields})
	}

	return processedRows, nil
}
