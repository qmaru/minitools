package common

import "bytes"

type DataJsonBasic interface {
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	Unmarshal(data []byte, v any) error
	Valid(data []byte) bool
	Compact(dst *bytes.Buffer, src []byte) error
	Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
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

// PrettyPrint Pretty Print JSON
func (dj DataJsonDefault[T]) PrettyPrint(input []byte) ([]byte, error) {
	var tmp any
	if err := dj.Json.Unmarshal(input, &tmp); err != nil {
		return nil, err
	}
	return dj.Json.MarshalIndent(tmp, "", "  ")
}

// MergeJSON Merge two JSON objects
func (dj DataJsonDefault[T]) MergeJSON(a, b []byte) ([]byte, error) {
	mapA, err := dj.RawJson2Map(a)
	if err != nil {
		return nil, err
	}
	mapB, err := dj.RawJson2Map(b)
	if err != nil {
		return nil, err
	}
	for k, v := range mapB {
		mapA[k] = v
	}
	return dj.MapToRawJSON(mapA)
}

// DeepCopy Deep Copy JSON object
func (dj DataJsonDefault[T]) DeepCopy(input any) (any, error) {
	data, err := dj.Json.Marshal(input)
	if err != nil {
		return nil, err
	}
	var output any
	if err := dj.Json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output, nil
}

// Compact compact json bytes
func (dj DataJsonDefault[T]) Compact(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := dj.Json.Compact(&buf, input); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// IndentJSON reformat existing JSON bytes
func (dj DataJsonDefault[T]) IndentJSON(input []byte, prefix, indent string) ([]byte, error) {
	var buf bytes.Buffer
	if err := dj.Json.Indent(&buf, input, prefix, indent); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
