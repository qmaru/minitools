package sonic

import (
	"io"

	sjson "github.com/bytedance/sonic"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type SonicJSONBasic = common.DataJsonDefault[SonicJSON]

// SonicJSON
type SonicJSON struct{}

func (s SonicJSON) NewDecoder(r io.Reader) sjson.Decoder {
	return sjson.ConfigDefault.NewDecoder(r)
}

func (s SonicJSON) NewEncoder(w io.Writer) sjson.Encoder {
	return sjson.ConfigDefault.NewEncoder(w)
}

func (s SonicJSON) Marshal(v any) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s SonicJSON) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return sjson.MarshalIndent(v, prefix, indent)
}

func (s SonicJSON) Unmarshal(data []byte, v any) error {
	return sjson.Unmarshal(data, v)
}

func (s SonicJSON) Valid(data []byte) bool {
	return sjson.Valid(data)
}

func New() *SonicJSONBasic {
	return new(SonicJSONBasic)
}
