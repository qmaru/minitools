package sonic

import (
	sjson "github.com/bytedance/sonic"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type SonicJSONBasic = common.DataJsonDefault[SonicJSON]

// SonicJSON
type SonicJSON struct{}

func (s SonicJSON) Marshal(v any) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s SonicJSON) Unmarshal(data []byte, v any) error {
	return sjson.Unmarshal(data, v)
}

func New() *SonicJSONBasic {
	return new(SonicJSONBasic)
}
