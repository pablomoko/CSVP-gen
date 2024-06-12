package processor

import (
	"CSVPgen/internal/types"
)

type ColumnProcessor interface {
	Process(value string) string
}
type ColumnProcessorMap map[string]ColumnProcessor

func ProcessRow(row types.Row, columnProcessors ColumnProcessorMap) types.Row {
	var fields []types.StructField
	for _, value := range row.Fields {
		processor, ok := columnProcessors[value.Name]
		if ok {
			value.Value = processor.Process(value.Value)
		}
		fields = append(fields, value)
	}
	return types.Row{Fields: fields}
}

func ProcessRows(rows []types.Row, columnProcessors ColumnProcessorMap) ([]types.Row, error) {
	var processedRows []types.Row

	for _, row := range rows {
		processedRows = append(processedRows, ProcessRow(row, columnProcessors))
	}

	return processedRows, nil
}
