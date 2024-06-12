package types

type StructField struct {
	Name  string
	Value string
}

type Row struct {
	Fields []StructField
}

type CSVMetadata struct {
	FileName   string // Nombre del archivo CSV
	NumRows    int    // Número total de filas
	NumColumns int    // Número total de columnas
}
