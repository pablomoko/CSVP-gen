package csv

type ColumnProcessor interface {
	Process(value string) string
}
type ColumnProcessorMap map[string]ColumnProcessor

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
