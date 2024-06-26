package generator

import (
	"CSVPgen/internal/types"
)

type ColumnGenerator interface {
	Generate(row types.Row) ([]types.StructField, error)
}

type ColumnGeneratorMap map[string]ColumnGenerator

func ProcessRowsWithGenerators(rows []types.Row, columnGenerators ColumnGeneratorMap) ([]types.Row, error) {
	var processedRows []types.Row

	for _, row := range rows {
		for _, generator := range columnGenerators {
			newFields, err := generator.Generate(row)
			if err != nil {
				return nil, err
			}
			row.Fields = append(row.Fields, newFields...)
		}
		processedRows = append(processedRows, row)
	}

	return processedRows, nil
}
