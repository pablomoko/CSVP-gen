package types

type StructField struct {
	Name  string
	Value string
}

type Row struct {
	Fields []StructField
}
