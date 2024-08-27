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
    data, err := jsonreader.LoadFileJSON("example.json")
    // ...
    data2, err := jsonreader.ReadJSON("{\"id\":1,\"father\":\"Mark\",\"mother\":\"Charlotte\",\"children\":2}")
    // ...

}

```