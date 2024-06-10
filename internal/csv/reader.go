package csv

import (
	"encoding/csv"
	"os"
)

type ColumnProcessor interface {
	Process(value string) string
}

type ColumnProcessorMap map[string]ColumnProcessor

type StructField struct {
	Name  string
	Value string
}

type Row struct {
	Fields []StructField
}

func OpenCSV(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ProcessRow(row []string, columnNames []string, columnProcessors ColumnProcessorMap) Row {
	var fields []StructField
	for i, value := range row {
		fieldName := columnNames[i]
		processor, ok := columnProcessors[fieldName]
		if ok {
			value = processor.Process(value)
		}
		fields = append(fields, StructField{Name: fieldName, Value: value})
	}
	return Row{Fields: fields}
}

func ReadCSV(file *os.File, columnProcessors ColumnProcessorMap) ([]Row, error) {
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	columnNames := rows[0]
	var processedRows []Row

	for _, row := range rows[1:] {
		processedRows = append(processedRows, ProcessRow(row, columnNames, columnProcessors))
	}

	return processedRows, nil
}
