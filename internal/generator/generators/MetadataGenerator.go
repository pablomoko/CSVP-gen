package generators

import "CSVPgen/internal/types"

type MetadataGenerator struct {
	Metadata *types.CSVMetadata
}

func (m *MetadataGenerator) Generate(row types.Row) ([]types.StructField, error) {
	m.Metadata.FileName = "test"
	m.Metadata.NumRows++
	m.Metadata.NumColumns = len(row.Fields)
	return nil, nil
}
