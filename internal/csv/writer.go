package csv

import (
	"CSVPgen/internal/types"
	"encoding/csv"
	"os"
)

func WriteCSV(file *os.File, rows []types.Row) error {
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Obtener los nombres de las columnas desde la primera fila
	if len(rows) == 0 {
		return nil // No hay filas para escribir
	}

	columnNames := make([]string, len(rows[0].Fields))
	for i, field := range rows[0].Fields {
		columnNames[i] = field.Name
	}

	// Escribir los nombres de las columnas
	if err := writer.Write(columnNames); err != nil {
		return err
	}

	// Escribir las filas
	for _, row := range rows {
		record := make([]string, len(row.Fields))
		for i, field := range row.Fields {
			record[i] = field.Value
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
