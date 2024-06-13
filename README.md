# CSVP-gen (CSV Processing & Generating Library)

## Introduction

This Go library provides tools for reading, writing, generating, and processing CSV files.

## Installation

To use this library in your Go project, import it using:

```shell
go get -u github.com/pablomoko/CSVP-gen
```
## Project Structure

- **cmd**: Contains the main entry point of the application.
- **internal**: Contains the internal source code of the library.
  - **csv**: Functionality for reading and writing CSV files.
  - **generator**: Tools for generating data and creating CSV files.
  - **processor**: Functions for processing CSV data through transformations.
  - **types**: Definitions of shared types used throughout the library.
- **data**: Example directory containing input and output CSV files for testing.

## Usage

### Reading CSV Files

To read a CSV file, you can use the `ReadCSV` function from the `csv` package. Here's an example:



```go
package main

import (
    "CSVPgen/internal/csv"
    "fmt"
    "os"
)

func main() {
    filePath := "path/to/your/csv/file.csv"
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    columnNames, rows, err := csv.ReadCSV(file, 0)
    if err != nil {
        fmt.Println("Error reading CSV:", err)
        return
    }

    // Process the CSV data here...
}

```
`ReadCSV`
### Writing CSV Files
To write a CSV file, you can use the `WriteCSV` function from the `csv` package. Here's an example:
```go
package main

import (
  "CSVPgen/internal/csv"
  "CSVPgen/internal/types"
  "fmt"
  "os"
)

func main() {
  filePath := "path/to/your/output/file.csv"
  file, err := os.Create(filePath)
  if err != nil {
    fmt.Println("Error creating file:", err)
    return
  }
  defer file.Close()

  // Example rows to write
  rows := []types.Row{
    {Fields: []types.StructField{{Name: "Name", Value: "John"}, {Name: "Age", Value: "30"}}},
    {Fields: []types.StructField{{Name: "Name", Value: "Jane"}, {Name: "Age", Value: "25"}}},
  }

  if err := csv.WriteCSV(file, rows); err != nil {
    fmt.Println("Error writing CSV:", err)
    return
  }

  fmt.Println("CSV file written successfully")
}

```

### Generating CSV Data
To generate CSV data, you can create `generators` that implement the `ColumnGenerator` interface and use them with the `ProcessRowsWithGenerators` function. Here's an example:
```go
package main

import (
  CSVPgen "CSVPgen/internal/csv"
  Generator "CSVPgen/internal/generator"
  ColGenerators "CSVPgen/internal/generator/generators/column"
  MetaDataGenerator "CSVPgen/internal/generator/generators/metadata"
  "CSVPgen/internal/types"
  "fmt"
  "os"
)

func main() {
  filePath := "data/input/data_test.csv"
  file, err := CSVPgen.OpenCSV(filePath)
  if err != nil {
    fmt.Println(err)
  }

  _, rows, err := CSVPgen.ReadCSV(file, 0)
  if err != nil {
    fmt.Println("Error reading CSV:", err)
    return
  }

  sumGenerator := &ColGenerators.SumColumnGenerator{
    SourceColumns: []string{"Age", "Dog's Age"},
    NewColumnName: "SumColumn",
  }

  metadata := &types.CSVMetadata{
    FileName:   "test_metadata_2",
    NumRows:    0,
    NumColumns: 0,
  }

  metadataGenerator := &MetaDataGenerator.MetadataGenerator{
    Metadata: metadata,
  }

  columnGenerators := Generator.ColumnGeneratorMap{
    "SumColumn": sumGenerator,
    "Metadata":  metadataGenerator,
  }

  finalRows, err := Generator.ProcessRowsWithGenerators(rows, columnGenerators)
  if err != nil {
    fmt.Println("Error processing rows:", err)
    return
  }

  outputFilePath := "data/output/" + metadata.FileName + ".csv"
  outputFile, err := CSVPgen.CreateCSV(outputFilePath)
  if err != nil {
    fmt.Println("Error creating output file:", err)
    return
  }

  if err := CSVPgen.WriteCSV(outputFile, finalRows); err != nil {
    fmt.Println("Error writing CSV:", err)
    return
  }

  fmt.Println("CSV file generated and saved at:", outputFilePath)
}

```

### Processing CSV Data
To process CSV data, you can create `processors` that implement the `ColumnProcessor` interface and use them with the `ProcessRows` function. Here's an example:
```go
package main

import (
  CSVPgen "CSVPgen/internal/csv"
  Processor "CSVPgen/internal/processor"
  Processors "CSVPgen/internal/processor/processors"
  "CSVPgen/internal/types"
  "fmt"
  "os"
)

func main() {
  filePath := "data/input/data_test.csv"
  file, err := CSVPgen.OpenCSV(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  _, rows, err := CSVPgen.ReadCSV(file, 0)
  if err != nil {
    fmt.Println("Error reading CSV:", err)
    return
  }

  columnProcessors := Processor.ColumnProcessorMap{
    "Age": Processors.DivisorProcessor{Divisor: 2},
  }

  processedRows, err := Processor.ProcessRows(rows, columnProcessors)
  if err != nil {
    fmt.Println("Error processing rows:", err)
    return
  }

  metadata := &types.CSVMetadata{
    FileName:   "test_metadata_1",
    NumRows:    0,
    NumColumns: 0,
  }

  outputFilePath := "data/output/" + metadata.FileName + ".csv"
  outputFile, err := CSVPgen.CreateCSV(outputFilePath)
  if err != nil {
    fmt.Println("Error creating output file:", err)
    return
  }

  if err := CSVPgen.WriteCSV(outputFile, processedRows); err != nil {
    fmt.Println("Error writing CSV:", err)
    return
  }

  fmt.Println("CSV file processed and saved at:", outputFilePath)
}

```

### Creating Custom Processors
To create a custom `processor`, implement the `ColumnProcessor` interface:
```go
package processors

type DivisorProcessor struct {
    Divisor int
}

func (p DivisorProcessor) Process(value string) string {
    // Custom logic to process the value
    return value // Implement the actual processing logic
}

```

### Creating Custom Generators
To create a custom `generator`, implement the `ColumnGenerator` interface:
```go
package column

import "CSVPgen/internal/types"

type SumColumnGenerator struct {
    SourceColumns []string
    NewColumnName string
}

func (g *SumColumnGenerator) Generate(row types.Row) ([]types.StructField, error) {
    sum := 0
    for _, col := range g.SourceColumns {
        for _, field := range row.Fields {
            if field.Name == col {
                value, _ := strconv.Atoi(field.Value)
                sum += value
            }
        }
    }
    newField := types.StructField{
        Name:  g.NewColumnName,
        Value: strconv.Itoa(sum),
    }
    return []types.StructField{newField}, nil
}

```


### Metadata Generator
The `MetadataGenerator` is a special type of `generator` that collects metadata about the CSV file, such as the number of rows and columns. This metadata can be used for various purposes, such as logging or generating summary reports.
```go
package metadata

import (
    "CSVPgen/internal/types"
)

type MetadataGenerator struct {
  Metadata *types.CSVMetadata
}

func (m *MetadataGenerator) Generate(row types.Row) ([]types.StructField, error) {
  m.Metadata.FileName = "file_name"
  m.Metadata.NumRows++
  m.Metadata.NumColumns = len(row.Fields)
  return nil, nil
}




```


## Contributing

Contributions are welcome! If you would like to contribute to this library, please create a pull request detailing your proposed changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
