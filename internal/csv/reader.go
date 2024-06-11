package csv

import (
	"CSVPgen/internal/types"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(file *os.File, headerRow int) ([]string, []types.Row, error) {
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if headerRow >= len(rows) {
		return nil, nil, fmt.Errorf("headerRow out of range")
	}

	columnNames := rows[headerRow]

	var processedRows []types.Row

	for _, row := range rows[headerRow+1:] {
		var fields []types.StructField
		for i, value := range row {
			fields = append(fields, types.StructField{Name: columnNames[i], Value: value})
		}
		processedRows = append(processedRows, types.Row{Fields: fields})
	}

	return columnNames, processedRows, nil
}
