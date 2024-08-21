package jsonreader

import "os"

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
	_, err = jsonFile.Read(content)
	if err != nil {
		return JSONStruct{}, err
	}

	ret, err := ReadJSON(string(content))
	if err != nil {
		return JSONStruct{}, err
	}

	return ret, nil
}

func ReadJSON(content string) (JSONStruct, error) {
	// Read char by char
	// @todo

	return JSONStruct{}, nil
}
