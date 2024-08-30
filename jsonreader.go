package jsonreader

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type DataStruct interface {
	GetHeaders() []string
	GetRows() map[int]map[string]string
}

type JSONStruct struct {
	Headers []string                  // Headers[0] = "label"
	Rows    map[int]map[string]string // Rows[0]["label"] = "value"
}

func (c JSONStruct) GetHeaders() []string {
	return c.Headers
}

func (c JSONStruct) GetRows() map[int]map[string]string {
	return c.Rows
}

func LoadFileJSON(filename string) (JSONStruct, error) {
	content, err := getContentFile(filename)
	if err != nil {
		return JSONStruct{}, err
	}

	ret, err := ReadJSON(content)
	if err != nil {
		return JSONStruct{}, err
	}

	return ret, nil
}

func LoadFileJSONLine(filename string) (JSONStruct, error) {
	content, err := getContentFile(filename)
	if err != nil {
		return JSONStruct{}, err
	}

	ret, err := ReadJSONLine(content)
	if err != nil {
		return JSONStruct{}, err
	}

	return ret, nil
}

func getContentFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content := []byte{}

	byteValue, _ := os.ReadFile(filename)
	content = append(content, byteValue...)

	return string(content), nil
}

func ReadJSON(content string) (JSONStruct, error) {
	arrayGroup := 0
	objectGroup := 0
	objects := []string{}
	object := []byte{}

	for i := 0; i < len(content); i++ {
		if content[i] == '[' && objectGroup == 0 { // Start of array
			arrayGroup++
		} else if content[i] == ']' && objectGroup == 0 { // End of array
			arrayGroup--
		} else if content[i] == '{' {
			objectGroup++
		} else if content[i] == '}' {
			objectGroup--
			if objectGroup == 0 { // End of main object
				objects = append(objects, "{"+string(object)+"}")
				object = []byte{}
			}
		} else if objectGroup > 0 { // In object
			object = append(object, content[i])
		}

		if i == 0 && arrayGroup == 0 { // If not is an array then it's invalid
			return JSONStruct{}, errors.New("Invalid JSON array")
		}
	}

	if arrayGroup != 0 || objectGroup != 0 { // If some groups are not closed then it's invalid
		return JSONStruct{}, errors.New("Invalid JSON")
	}

	res, err := getStruct(objects)
	if err != nil {
		return JSONStruct{}, err
	}

	return res, nil
}

func ReadJSONLine(content string) (JSONStruct, error) {
	objects := strings.Split(content, "\n")

	res, err := getStruct(objects)
	if err != nil {
		return JSONStruct{}, err
	}

	return res, nil
}

func getStruct(objects []string) (JSONStruct, error) {
	res := JSONStruct{
		Headers: []string{},
		Rows:    make(map[int]map[string]string),
	}

	headerMap := make(map[string]bool)

	for i, obj := range objects {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(obj), &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		row := make(map[string]string)

		for key, value := range data {
			if !headerMap[key] {
				res.Headers = append(res.Headers, key)
				headerMap[key] = true
			}

			valueStr, err := json.Marshal(value)
			if err != nil {
				return JSONStruct{}, err
			}

			row[key] = removeQuotes(string(valueStr))
		}

		res.Rows[i] = row
	}

	return res, nil
}

func removeQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}

	return s
}
