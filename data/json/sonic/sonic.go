package sonic

import (
	"bytes"
	"io"

	sjson "github.com/bytedance/sonic"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type RawMessage = sjson.NoCopyRawMessage

type SonicJSONBasic = common.DataJsonDefault[SonicJSON]

var (
	ConfigStd     = sjson.ConfigStd
	ConfigDefault = sjson.ConfigDefault
	ConfigFastest = sjson.ConfigFastest
)

// SonicJSON
type SonicJSON struct{}

func (s SonicJSON) NewDecoder(r io.Reader) sjson.Decoder {
	return sjson.ConfigDefault.NewDecoder(r)
}

func (s SonicJSON) NewEncoder(w io.Writer) sjson.Encoder {
	return sjson.ConfigDefault.NewEncoder(w)
}

func (s SonicJSON) NewDecoderWithNumber(r io.Reader) sjson.Decoder {
	dec := sjson.ConfigDefault.NewDecoder(r)
	dec.UseNumber()
	return dec
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

func (s SonicJSON) UnmarshalNumber(data []byte, v any) error {
	dec := sjson.ConfigDefault.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	return dec.Decode(v)
}

func (s SonicJSON) Valid(data []byte) bool {
	return sjson.Valid(data)
}

func (s SonicJSON) Compact(dst *bytes.Buffer, src []byte) error {
	var v any
	if err := ConfigStd.Unmarshal(src, &v); err != nil {
		return err
	}

	dst.Grow(len(src))
	data, err := ConfigStd.Marshal(v)
	if err != nil {
		return err
	}

	dst.Write(data)
	return nil
}

func (s SonicJSON) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	var v any
	if err := ConfigStd.Unmarshal(src, &v); err != nil {
		return err
	}

	dst.Grow(2 * len(src))
	data, err := ConfigStd.MarshalIndent(v, prefix, indent)
	if err != nil {
		return err
	}

	_, err = dst.Write(data)
	return err
}

func New() *SonicJSONBasic {
	return new(SonicJSONBasic)
}
