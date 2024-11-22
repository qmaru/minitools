package convert

// DataConvertSuiteBasic
type DataConvertSuiteBasic struct{}

// Float2uint Float to int fixed
func (ds *DataConvertSuiteBasic) Float2uint(f float64) int64 {
	u := int64((f * 100) + 0.5)
	return u
}

func New() *DataConvertSuiteBasic {
	return new(DataConvertSuiteBasic)
}
