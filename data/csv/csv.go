package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// CSVSuiteBasic
type CSVSuiteBasic[T any] struct{}

func (c *CSVSuiteBasic[T]) openFile(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// ReadFile
func (c *CSVSuiteBasic[T]) ReadFile(filepath string) ([]T, error) {
	records, err := c.openFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("csv content not found")
	}

	headers := records[0]
	var results []T

	var t T
	structType := reflect.TypeOf(t)

	if structType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type T must be a struct")
	}

	fieldMap := make(map[string]int)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("csv")
		if tag != "" {
			fieldMap[tag] = i
		}
	}

	for _, row := range records[1:] {
		entryPtr := reflect.New(structType).Elem()

		for i, value := range row {
			if i >= len(headers) {
				continue
			}
			fieldIndex, exists := fieldMap[headers[i]]
			if !exists {
				continue
			}

			field := entryPtr.Field(fieldIndex)
			if !field.IsValid() || !field.CanSet() {
				continue
			}

			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Bool:
				if boolVal, err := strconv.ParseBool(value); err == nil {
					field.SetBool(boolVal)
				}
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
					field.SetInt(intVal)
				}
			case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
				if uintVal, err := strconv.ParseUint(value, 10, 64); err == nil {
					field.SetUint(uintVal)
				}
			case reflect.Float64, reflect.Float32:
				if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
					field.SetFloat(floatVal)
				}
			}
		}

		results = append(results, entryPtr.Interface().(T))
	}

	return results, nil
}

func New[T any]() *CSVSuiteBasic[T] {
	return new(CSVSuiteBasic[T])
}
