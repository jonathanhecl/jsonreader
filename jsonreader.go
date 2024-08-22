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

	arrayGroup := 0
	objectGroup := 0
	objects := []string{}
	object := []byte{}

	for i := 0; i < len(content); i++ {
		if content[i] == '[' && objectGroup == 0 {
			arrayGroup++
		} else if content[i] == '{' {
			objectGroup++
		} else if content[i] == ']' && objectGroup == 0 {
			arrayGroup--
		} else if content[i] == '}' {
			objectGroup--
			if objectGroup == 0 {
				objects = append(objects, string(object))
				object = []byte{}
			}
		} else if objectGroup > 0 {
			object = append(object, content[i])
		}

		if i == 0 && arrayGroup == 0 {
			return JSONStruct{}, errors.New("Invalid JSON array")
		}
	}

	if arrayGroup != 0 || objectGroup != 0 {
		return JSONStruct{}, errors.New("Invalid JSON")
	}

	fmt.Println("Objects: ")
	for i := 0; i < len(objects); i++ {
		fmt.Println(i, objects[i])
	}

	return JSONStruct{}, nil
}
