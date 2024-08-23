package aes

// AESSuiteCBCBasic
type AESSuiteCBCBasic struct{}

// AESSuiteGCMBasic
type AESSuiteGCMBasic struct{}

func NewCBC() *AESSuiteCBCBasic {
	return new(AESSuiteCBCBasic)
}

func NewGCM() *AESSuiteGCMBasic {
	return new(AESSuiteGCMBasic)
}
