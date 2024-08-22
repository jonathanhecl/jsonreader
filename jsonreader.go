package jsonreader

import (
	"errors"
	"fmt"
	"os"
)

type JSONStruct struct {
	Headers []string                  // Headers[0] = "label"
	Rows    map[int]map[string]string // Rows[0]["label"] = "value"
}

func LoadFileJSON(filename string) (JSONStruct, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return JSONStruct{}, err
	}
	defer jsonFile.Close()

	content := []byte{}

	byteValue, _ := os.ReadFile(filename)
	content = append(content, byteValue...)

	ret, err := ReadJSON(string(content))
	if err != nil {
		return JSONStruct{}, err
	}

	return ret, nil
}

func ReadJSON(content string) (JSONStruct, error) {
	fmt.Println(content)

	arrayGroup := 0
	objectGroup := 0
	for i := 0; i < len(content); i++ {
		fmt.Println(i, string(content[i]))
		if content[i] == '[' {
			arrayGroup++
		} else if content[i] == '{' {
			objectGroup++
		} else if content[i] == ']' {
			arrayGroup--
		} else if content[i] == '}' {
			objectGroup--
		}

		if i == 0 && arrayGroup == 0 {
			return JSONStruct{}, errors.New("Invalid JSON array")
		}
	}

	if arrayGroup != 0 || objectGroup != 0 {
		return JSONStruct{}, errors.New("Invalid JSON array")
	}

	return JSONStruct{}, nil
}
