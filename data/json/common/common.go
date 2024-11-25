package common

type DataJsonBasic interface {
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type DataJsonDefault[T DataJsonBasic] struct {
	Json T
}

// MapToRawJSON Any MAP to Raw JSON
func (dj DataJsonDefault[T]) MapToRawJSON(input any) ([]byte, error) {
	return dj.Json.Marshal(input)
}

// RawJsonArray2MapArray JSON Array to []map[string]any
func (dj DataJsonDefault[T]) RawJsonArray2MapArray(input []byte) (output []map[string]any, err error) {
	err = dj.Json.Unmarshal(input, &output)
	return
}

// RawJson2Map JSON Map to map[string]any
func (dj DataJsonDefault[T]) RawJson2Map(input []byte) (output map[string]any, err error) {
	output = make(map[string]any)
	err = dj.Json.Unmarshal(input, &output)
	return
}

// RawArray2Array JSON Array to []any
func (dj DataJsonDefault[T]) RawArray2Array(input []byte) (output []any, err error) {
	err = dj.Json.Unmarshal(input, &output)
	return
}
