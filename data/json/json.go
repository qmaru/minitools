package json

import stdjson "encoding/json"

// DataJsonSuiteBasic
type DataJsonSuiteBasic struct{}

// MapToRawJSON Any MAP to Raw JSON
func (ds *DataJsonSuiteBasic) MapToRawJSON(input any) ([]byte, error) {
	return stdjson.Marshal(input)
}

// RawJsonArray2MapArray JSON Array to []map[string]any
func (ds *DataJsonSuiteBasic) RawJsonArray2MapArray(input []byte) (output []map[string]any, err error) {
	err = stdjson.Unmarshal(input, &output)
	return
}

// RawJson2Map JSON Map to map[string]any
func (ds *DataJsonSuiteBasic) RawJson2Map(input []byte) (output map[string]any, err error) {
	output = make(map[string]any)
	err = stdjson.Unmarshal(input, &output)
	return
}

// RawArray2Array JSON Array to []any
func (ds *DataJsonSuiteBasic) RawArray2Array(input []byte) (output []any, err error) {
	err = stdjson.Unmarshal(input, &output)
	return
}

func New() *DataJsonSuiteBasic {
	return new(DataJsonSuiteBasic)
}
