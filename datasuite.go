package minitools

import "encoding/json"

// DataSuiteBasic 数据转换基类
type DataSuiteBasic struct{}

// MapToRawJSON 任意 map 类型转换为原始 JSON 数据
func (ds *DataSuiteBasic) MapToRawJSON(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

// RawMaps2Maps 原始 map 数组转 []map[string]interface{}
func (ds *DataSuiteBasic) RawMaps2Maps(data []byte) (m []map[string]interface{}) {
	m = make([]map[string]interface{}, 0)
	_ = json.Unmarshal(data, &m)
	return
}

// RawMap2Map 原始 map 转 map[string]interface{}
func (ds *DataSuiteBasic) RawMap2Map(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	return
}

// RawArray2Array 原始数组类型转 []interface{}
func (ds *DataSuiteBasic) RawArray2Array(data []byte) (m []interface{}) {
	m = make([]interface{}, 0)
	_ = json.Unmarshal(data, &m)
	return
}

// Float2uint 浮点转整型处理
func (ds *DataSuiteBasic) Float2uint(f float64) int64 {
	u := int64((f * 100) + 0.5)
	return u
}
