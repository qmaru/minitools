package sonic

import (
	sjson "github.com/bytedance/sonic"

	"github.com/qmaru/minitools/v2/data/json/common"
)

// SonicJSON
type SonicJSON struct{}

func (s SonicJSON) Marshal(v interface{}) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s SonicJSON) Unmarshal(data []byte, v interface{}) error {
	return sjson.Unmarshal(data, v)
}

func New() *common.DataJsonDefault[SonicJSON] {
	return new(common.DataJsonDefault[SonicJSON])
}
