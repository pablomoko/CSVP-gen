package csv

import (
	"encoding/csv"
	"os"
)

func WriteCSV(file *os.File, rows []Row, columnNames []string) error {
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(columnNames); err != nil {
		return err
	}

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
