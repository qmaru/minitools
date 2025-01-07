package standard

import (
	"io"

	sjson "encoding/json"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type StandardJSONBasic = common.DataJsonDefault[StandardJSON]

// StandardJSON
type StandardJSON struct{}

func (s StandardJSON) NewDecoder(r io.Reader) *sjson.Decoder {
	return sjson.NewDecoder(r)
}

func (s StandardJSON) NewEncoder(w io.Writer) *sjson.Encoder {
	return sjson.NewEncoder(w)
}

func (s StandardJSON) Marshal(v any) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s StandardJSON) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return sjson.MarshalIndent(v, prefix, indent)
}

func (s StandardJSON) Unmarshal(data []byte, v any) error {
	return sjson.Unmarshal(data, v)
}

func (s StandardJSON) Valid(data []byte) bool {
	return sjson.Valid(data)
}

func New() *StandardJSONBasic {
	return new(StandardJSONBasic)
}
