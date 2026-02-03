package convert

// ConvertSuiteBasic
type ConvertSuiteBasic struct{}

// Float2uint Float to int fixed
func (ds *ConvertSuiteBasic) Float2uint(f float64) int64 {
	u := int64((f * 100) + 0.5)
	return u
}

func New() *ConvertSuiteBasic {
	return new(ConvertSuiteBasic)
}
