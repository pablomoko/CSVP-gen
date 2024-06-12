package csv

import (
	"os"
	"testing"
)

func TestOpenCSV(t *testing.T) {
	// Crear un archivo temporal para la prueba
	tempFile, err := os.CreateTemp("", "test_open_csv_*.csv")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Eliminar el archivo temporal al final de la prueba

	// Intentar abrir el archivo
	file, err := OpenCSV(tempFile.Name())
	if err != nil {
		t.Errorf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Verificar que el archivo no es nulo
	if file == nil {
		t.Error("OpenCSV returned nil file")
	}
}

func TestCreateCSV(t *testing.T) {
	// Crear un archivo temporal para la prueba
	tempFile, err := os.CreateTemp("", "test_create_csv_*.csv")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Eliminar el archivo temporal al final de la prueba

	// Intentar crear el archivo
	file, err := CreateCSV(tempFile.Name())
	if err != nil {
		t.Errorf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	// Verificar que el archivo no es nulo
	if file == nil {
		t.Error("CreateCSV returned nil file")
	}
}
