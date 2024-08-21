# jsonreader
JSON Simple Reader

## Requirements
* A JSON array

## Features
* Read JSON files with headers and rows
* Don't require an deterministic struct to read

## Returned struct
- Headers []string
- Rows    map[int]map[string]string

## Installation
`go get github.com/jonathanhecl/jsonreader`

## Example
```go
import (
    "github.com/jonathanhecl/jsonreader"
)

func main() {
     // ...
}

```