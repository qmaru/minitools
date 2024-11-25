package gojson

import (
	gojson "github.com/goccy/go-json"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type GoJSONBasic = common.DataJsonDefault[GoJSON]

// GoJSON
type GoJSON struct{}

func (s GoJSON) Marshal(v any) ([]byte, error) {
	return gojson.Marshal(v)
}

func (s GoJSON) Unmarshal(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}

func New() *GoJSONBasic {
	return new(GoJSONBasic)
}
