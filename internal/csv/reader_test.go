package csv

import (
	"os"
	"testing"

	"CSVPgen/internal/types"
)

func TestReadCSV(t *testing.T) {
	// Crear un archivo CSV en memoria para la prueba
	csvData := `Name,Last Name,Age
John,Doe,30
Jane,Smith,25`
	file := createTempCSVFile(csvData)
	defer os.Remove(file.Name()) // Eliminar el archivo temporal al final de la prueba

	// Abrir el archivo CSV
	openedFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("Error opening CSV file: %v", err)
	}
	defer openedFile.Close()

	// Leer el archivo CSV
	headerRow := 0
	columnNames, rows, err := ReadCSV(openedFile, headerRow)
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	// Verificar que las columnas se hayan leído correctamente
	expectedColumnNames := []string{"Name", "Last Name", "Age"}
	if !equalSlice(expectedColumnNames, columnNames) {
		t.Errorf("Expected column names %v, got %v", expectedColumnNames, columnNames)
	}

	// Verificar que las filas se hayan leído correctamente
	expectedRows := []types.Row{
		{Fields: []types.StructField{{Name: "Name", Value: "John"}, {Name: "Last Name", Value: "Doe"}, {Name: "Age", Value: "30"}}},
		{Fields: []types.StructField{{Name: "Name", Value: "Jane"}, {Name: "Last Name", Value: "Smith"}, {Name: "Age", Value: "25"}}},
	}
	if !equalRows(expectedRows, rows) {
		t.Errorf("Expected rows %v, got %v", expectedRows, rows)
	}
}

// Función de utilidad para crear un archivo CSV temporal con los datos proporcionados
func createTempCSVFile(data string) *os.File {
	file, _ := os.CreateTemp("", "test_*.csv")
	file.WriteString(data)
	file.Close()
	return file
}

// Función de utilidad para comparar dos slices de strings
func equalSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// Función de utilidad para comparar dos slices de structs Row
func equalRows(rows1, rows2 []types.Row) bool {
	if len(rows1) != len(rows2) {
		return false
	}
	for i := range rows1 {
		if !equalFields(rows1[i].Fields, rows2[i].Fields) {
			return false
		}
	}
	return true
}

// Función de utilidad para comparar dos slices de structs StructField
func equalFields(fields1, fields2 []types.StructField) bool {
	if len(fields1) != len(fields2) {
		return false
	}
	for i := range fields1 {
		if fields1[i].Name != fields2[i].Name || fields1[i].Value != fields2[i].Value {
			return false
		}
	}
	return true
}
