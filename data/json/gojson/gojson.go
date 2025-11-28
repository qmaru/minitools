package gojson

import (
	"bytes"
	"io"

	gojson "github.com/goccy/go-json"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type RawMessage = gojson.RawMessage

type GoJSONBasic = common.DataJsonDefault[GoJSON]

// GoJSON
type GoJSON struct{}

func (s GoJSON) NewDecoder(r io.Reader) *gojson.Decoder {
	return gojson.NewDecoder(r)
}

func (s GoJSON) NewEncoder(w io.Writer) *gojson.Encoder {
	return gojson.NewEncoder(w)
}

func (s GoJSON) Marshal(v any) ([]byte, error) {
	return gojson.Marshal(v)
}

func (s GoJSON) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return gojson.MarshalIndent(v, prefix, indent)
}

func (s GoJSON) Unmarshal(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}

func (s GoJSON) Valid(data []byte) bool {
	return gojson.Valid(data)
}

func (s GoJSON) Compact(dst *bytes.Buffer, src []byte) error {
	return gojson.Compact(dst, src)
}

func (s GoJSON) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return gojson.Indent(dst, src, prefix, indent)
}

func New() *GoJSONBasic {
	return new(GoJSONBasic)
}
