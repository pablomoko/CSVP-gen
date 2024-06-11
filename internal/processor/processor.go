package processor

import (
	"CSVPgen/internal/types"
)

type ColumnProcessor interface {
	Process(value string) string
}
type ColumnProcessorMap map[string]ColumnProcessor

func ProcessRow(row types.Row, columnNames []string, columnProcessors ColumnProcessorMap) types.Row {
	var fields []types.StructField
	for i, value := range row.Fields {
		fieldName := columnNames[i]
		processor, ok := columnProcessors[fieldName]
		if ok {
			value.Value = processor.Process(value.Value)
		}
		fields = append(fields, value)
	}
	return types.Row{Fields: fields}
}

func ProcessRows(columnNames []string, rows []types.Row, columnProcessors ColumnProcessorMap) ([]types.Row, error) {
	var processedRows []types.Row

	for _, row := range rows {
		processedRows = append(processedRows, ProcessRow(row, columnNames, columnProcessors))
	}

	return processedRows, nil
}
