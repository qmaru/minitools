package standard

import (
	sjson "encoding/json"

	"github.com/qmaru/minitools/v2/data/json/common"
)

// StandardJSON
type StandardJSON struct{}

func (s StandardJSON) Marshal(v interface{}) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s StandardJSON) Unmarshal(data []byte, v interface{}) error {
	return sjson.Unmarshal(data, v)
}

func New() *common.DataJsonDefault[StandardJSON] {
	return new(common.DataJsonDefault[StandardJSON])
}
