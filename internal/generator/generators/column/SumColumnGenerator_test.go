package column

import (
	"CSVPgen/internal/types"
	"strconv"
	"strings"
	"testing"
)

func TestSumColumnGenerator_Generate(t *testing.T) {
	generator := &SumColumnGenerator{
		SourceColumns: []string{"Column1", "Column2"},
		NewColumnName: "SumColumn",
	}

	row := types.Row{
		Fields: []types.StructField{
			{Name: "Column1", Value: "10"},
			{Name: "Column2", Value: "20"},
		},
	}

	expectedSum := strconv.FormatFloat(30.0, 'f', -1, 32)

	result, err := generator.Generate(row)
	if err != nil {
		t.Errorf("Error generating sum column: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected one field, got %d", len(result))
	}

	if result[0].Name != "SumColumn" {
		t.Errorf("Expected field name 'SumColumn', got '%s'", result[0].Name)
	}

	if !strings.HasPrefix(result[0].Value, expectedSum) {
		t.Errorf("Expected sum value to start with '%s', got '%s'", expectedSum, result[0].Value)
	}
}
