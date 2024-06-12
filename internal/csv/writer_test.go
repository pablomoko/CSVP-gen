package csv

import (
	"CSVPgen/internal/types"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	// Crear un archivo temporal para la prueba
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Eliminar el archivo temporal al final de la prueba

	// Definir las filas de ejemplo
	rows := []types.Row{
		{Fields: []types.StructField{
			{Name: "Name", Value: "John"},
			{Name: "Last Name", Value: "Doe"},
			{Name: "Age", Value: "30"},
		}},
		{Fields: []types.StructField{
			{Name: "Name", Value: "Jane"},
			{Name: "Last Name", Value: "Smith"},
			{Name: "Age", Value: "25"},
		}},
	}

	// Escribir las filas en el archivo CSV
	if err := WriteCSV(tmpfile, rows); err != nil {
		t.Fatalf("Error writing CSV file: %v", err)
	}

	// Leer el contenido del archivo CSV
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	// Verificar el contenido del archivo CSV
	expectedContent := []byte(`Name,Last Name,Age
John,Doe,30
Jane,Smith,25
`)
	if !reflect.DeepEqual(content, expectedContent) {
		t.Errorf("Expected CSV content:\n%s\nGot:\n%s", expectedContent, content)
	}
}
