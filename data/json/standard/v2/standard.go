//go:build go1.27

package standardv2

import (
	"bytes"
	"errors"
	"io"

	"encoding/json/jsontext"
	sjson "encoding/json/v2"

	"github.com/qmaru/minitools/v2/data/json/common"
)

type RawMessage = jsontext.Value

type StandardJSONBasic = common.DataJsonDefault[StandardJSON]

// StandardJSON
type StandardJSON struct{}

type Decoder struct {
	dec *jsontext.Decoder
}

type Encoder struct {
	w            io.Writer
	escapeHTML   bool
	indentPrefix string
	indentValue  string
}

func (d *Decoder) Decode(v any) error {
	val, err := d.dec.ReadValue()
	if err != nil {
		return err
	}
	return sjson.Unmarshal(val, v)
}

func (d *Decoder) More() bool {
	k := d.dec.PeekKind()
	return k != ']' && k != '}' && k != 0
}

func (d *Decoder) Token() (jsontext.Token, error) {
	return d.dec.ReadToken()
}

func (d *Decoder) Raw() *jsontext.Decoder {
	return d.dec
}

func (e *Encoder) Encode(v any) error {
	opts := []jsontext.Options{}

	if e.escapeHTML {
		opts = append(opts, jsontext.EscapeForHTML(true))
	}

	if e.indentPrefix != "" || e.indentValue != "" {
		opts = append(opts, jsontext.WithIndentPrefix(e.indentPrefix), jsontext.WithIndent(e.indentValue))
	}

	return sjson.MarshalWrite(e.w, v, opts...)
}

func (e *Encoder) SetEscapeHTML(on bool) {
	e.escapeHTML = on
}

func (e *Encoder) SetIndent(prefix, indent string) {
	e.indentValue = indent
	e.indentPrefix = prefix
}

func (s StandardJSON) NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		dec: jsontext.NewDecoder(r),
	}
}

func (s StandardJSON) NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// NewDecoderWithNumber not supported in jsontext package, so return nil
func (s StandardJSON) NewDecoderWithNumber(r io.Reader) *jsontext.Decoder {
	return nil
}

func (s StandardJSON) Marshal(v any) ([]byte, error) {
	return sjson.Marshal(v)
}

func (s StandardJSON) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return sjson.Marshal(v, jsontext.WithIndentPrefix(prefix), jsontext.WithIndent(indent))
}

func (s StandardJSON) Unmarshal(data []byte, v any) error {
	return sjson.Unmarshal(data, v)
}

func (s StandardJSON) UnmarshalNumber(data []byte, v any) error {
	return sjson.Unmarshal(data, v,
		sjson.WithUnmarshalers(
			sjson.UnmarshalFromFunc(func(dec *jsontext.Decoder, val *any) error {
				if dec.PeekKind() == '0' { // JSON number
					*val = jsontext.Value(nil)
				}
				return errors.ErrUnsupported
			}),
		),
	)
}

func (s StandardJSON) Valid(data []byte) bool {
	return jsontext.Value(data).IsValid()
}

func (s StandardJSON) Compact(dst *bytes.Buffer, src []byte) error {
	v := jsontext.Value(append([]byte(nil), src...))
	if err := v.Compact(); err != nil {
		return err
	}
	_, err := dst.Write(v)
	return err
}

func (s StandardJSON) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	v := jsontext.Value(append([]byte(nil), src...))
	if err := v.Indent(jsontext.WithIndentPrefix(prefix), jsontext.WithIndent(indent)); err != nil {
		return err
	}
	_, err := dst.Write(v)
	return err
}

func New() *StandardJSONBasic {
	return new(StandardJSONBasic)
}
