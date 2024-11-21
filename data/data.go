package data

import "encoding/json"

// DataSuiteBasic
type DataSuiteBasic struct{}

// MapToRawJSON Any MAP to Raw JSON
func (ds *DataSuiteBasic) MapToRawJSON(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

// RawJsonArray2MapArray JSON Array to []map[string]interface{}
func (ds *DataSuiteBasic) RawJsonArray2MapArray(input []byte) (output []map[string]interface{}, err error) {
	err = json.Unmarshal(input, &output)
	return
}

// RawJson2Map JSON Map to map[string]interface{}
func (ds *DataSuiteBasic) RawJson2Map(input []byte) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	err = json.Unmarshal(input, &output)
	return
}

// RawArray2Array JSON Array to []interface{}
func (ds *DataSuiteBasic) RawArray2Array(input []byte) (output []interface{}, err error) {
	err = json.Unmarshal(input, &output)
	return
}

// Float2uint Float to int fixed
func (ds *DataSuiteBasic) Float2uint(f float64) int64 {
	u := int64((f * 100) + 0.5)
	return u
}

func New() *DataSuiteBasic {
	return new(DataSuiteBasic)
}
