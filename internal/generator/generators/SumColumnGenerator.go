package generators

import (
	"CSVPgen/internal/types"
	"fmt"
	"strconv"
)

type SumColumnGenerator struct {
	SourceColumns []string
	NewColumnName string
}

func (g *SumColumnGenerator) Generate(row types.Row) ([]types.StructField, error) {
	sum := 0.0
	for _, col := range g.SourceColumns {
		found := false
		for _, field := range row.Fields {
			if field.Name == col {
				value, err := strconv.ParseFloat(field.Value, 32)
				if err != nil {
					return nil, fmt.Errorf("error converting value to float: %v", err)
				}
				sum += value
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("column %s not found in row", col)
		}
	}
	newField := types.StructField{
		Name:  g.NewColumnName,
		Value: fmt.Sprintf("%f", sum),
	}
	return []types.StructField{newField}, nil
}
