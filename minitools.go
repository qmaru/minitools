package minitools


// AESSuite AES加解密
func AESSuite() *AESSuiteBasic {
	return new(AESSuiteBasic)
}

// DataSuite 数据转换
func DataSuite() *DataSuiteBasic {
	return new(DataSuiteBasic)
}

// FileSuite 文件操作
func FileSuite() *FileSuiteBasic {
	return new(FileSuiteBasic)
}

// TimeSuite 时间转换
func TimeSuite() *TimeSuiteBasic {
	return new(TimeSuiteBasic)
}
