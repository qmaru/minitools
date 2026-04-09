//go:build jsonv2
// +build jsonv2

package minitools

import (
	"bytes"
	"reflect"
	"testing"

	standardv1 "github.com/qmaru/minitools/v2/data/json/standard/v1"
	standardv2 "github.com/qmaru/minitools/v2/data/json/standard/v2"
)

func equal(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

func roundtrip(t *testing.T, name string, input any) {
	v1 := standardv1.New()
	v2 := standardv2.New()

	t.Run(name, func(t *testing.T) {
		// --- Marshal ---
		b1, err1 := v1.Json.Marshal(input)
		b2, err2 := v2.Json.Marshal(input)

		if (err1 != nil) != (err2 != nil) {
			t.Fatalf("marshal error mismatch: v1=%v v2=%v", err1, err2)
		}

		// --- cross unmarshal ---
		var o11, o12, o21, o22 any

		if err := v1.Json.Unmarshal(b1, &o11); err != nil {
			t.Fatal(err)
		}
		if err := v2.Json.Unmarshal(b1, &o12); err != nil {
			t.Fatal(err)
		}
		if err := v1.Json.Unmarshal(b2, &o21); err != nil {
			t.Fatal(err)
		}
		if err := v2.Json.Unmarshal(b2, &o22); err != nil {
			t.Fatal(err)
		}

		// --- compare ---
		if !equal(o11, o12) {
			t.Fatalf("v1->v1 vs v1->v2 mismatch\n%#v\n%#v", o11, o12)
		}
		if !equal(o21, o22) {
			t.Fatalf("v2->v1 vs v2->v2 mismatch\n%#v\n%#v", o21, o22)
		}

		// --- valid ---
		if v1.Json.Valid(b1) != v2.Json.Valid(b1) {
			t.Fatalf("valid mismatch on b1")
		}
		if v1.Json.Valid(b2) != v2.Json.Valid(b2) {
			t.Fatalf("valid mismatch on b2")
		}
	})
}

func TestCompatibility(t *testing.T) {
	type S struct {
		A int     `json:"a"`
		B string  `json:"b,omitempty"`
		C *string `json:"c,omitempty"`
	}

	str := "x"

	cases := map[string]any{
		// --- scalar ---
		"nil":   nil,
		"bool":  true,
		"int":   123,
		"float": 0.1 + 0.2,
		"str":   "hello",

		// --- basic ---
		"map": map[string]any{"a": 1, "b": "x"},
		"arr": []any{1, "x", true, nil},

		// --- nested ---
		"nested": map[string]any{
			"a": map[string]any{
				"b": []any{1, map[string]any{"c": true}},
			},
		},

		// --- empty ---
		"empty": map[string]any{
			"m": map[string]any{},
			"a": []any{},
			"s": "",
		},

		// --- string edge ---
		"string_escape": map[string]any{
			"quote": `"a"`,
			"nl":    "a\nb",
			"utf8":  "你好🌍",
		},

		// --- number edge ---
		"big_int": int64(1<<60 + 1),
		"exp":     1e-9,

		// --- struct ---
		"struct_full":  S{A: 1, B: "x", C: &str},
		"struct_empty": S{A: 1},
	}

	for name, v := range cases {
		roundtrip(t, name, v)
	}
}

func TestDecoder_Stream(t *testing.T) {
	v1 := standardv1.New()
	v2 := standardv2.New()

	stream := []byte(`{"a":1} {"b":2}{"c":[1,2,3]}`)

	t.Run("stream decode v1", func(t *testing.T) {
		dec := v1.Json.NewDecoder(bytes.NewReader(stream))

		var results []any
		for {
			var v any
			if err := dec.Decode(&v); err != nil {
				break
			}
			t.Logf("decoded: %#v", v)
			results = append(results, v)
		}

		if len(results) != 3 {
			t.Fatalf("v1 decode count mismatch: %d", len(results))
		}
	})

	t.Run("stream decode v2", func(t *testing.T) {
		dec := v2.Json.NewDecoder(bytes.NewReader(stream))

		var results []any
		for {
			var v any
			if err := dec.Decode(&v); err != nil {
				break
			}
			t.Logf("decoded: %#v", v)
			results = append(results, v)
		}

		if len(results) != 3 {
			t.Fatalf("v2 decode count mismatch: %d", len(results))
		}
	})
}

func TestEncoder_Stream(t *testing.T) {
	v1 := standardv1.New()
	v2 := standardv2.New()

	inputs := []any{
		map[string]any{"a": 1},
		map[string]any{"b": "x"},
		[]any{1, 2, 3},
	}

	t.Run("stream encoder v1", func(t *testing.T) {
		var buf bytes.Buffer
		enc := v1.Json.NewEncoder(&buf)

		for _, v := range inputs {
			if err := enc.Encode(v); err != nil {
				t.Fatal(err)
			}
			t.Logf("encoded: %s", buf.String())
		}

		dec := v1.Json.NewDecoder(bytes.NewReader(buf.Bytes()))

		for range inputs {
			var out any
			if err := dec.Decode(&out); err != nil {
				t.Fatal(err)
			}
			t.Logf("decoded: %#v", out)
		}
	})

	t.Run("stream encoder v2", func(t *testing.T) {
		var buf bytes.Buffer
		enc := v2.Json.NewEncoder(&buf)

		for _, v := range inputs {
			if err := enc.Encode(v); err != nil {
				t.Fatal(err)
			}
			t.Logf("encoded: %s", buf.String())
		}

		dec := v2.Json.NewDecoder(bytes.NewReader(buf.Bytes()))

		for range inputs {
			var out any
			if err := dec.Decode(&out); err != nil {
				t.Fatal(err)
			}
			t.Logf("decoded: %#v", out)
		}
	})
}

func TestInvalidJSON(t *testing.T) {
	v1 := standardv1.New()
	v2 := standardv2.New()

	invalid := [][]byte{
		[]byte(``),
		[]byte(` `),
		[]byte(`{`),
		[]byte(`{"a":}`),
		[]byte(`[1,]`),
		[]byte(`{"a":01}`),
		[]byte(`{"a":1}xxx`),
	}

	for _, b := range invalid {
		if v1.Json.Valid(b) != v2.Json.Valid(b) {
			t.Fatalf("valid mismatch: %s", b)
		}

		var x any
		e1 := v1.Json.Unmarshal(b, &x)
		e2 := v2.Json.Unmarshal(b, &x)

		if (e1 != nil) != (e2 != nil) {
			t.Fatalf("unmarshal error mismatch: %s", b)
		}
	}
}

func TestFormat(t *testing.T) {
	v1 := standardv1.New()
	v2 := standardv2.New()

	raw := []byte(`{"a":1,"b":[1,2,3]}`)

	var c1, c2 bytes.Buffer
	if err := v1.Json.Compact(&c1, raw); err != nil {
		t.Fatal(err)
	}
	if err := v2.Json.Compact(&c2, raw); err != nil {
		t.Fatal(err)
	}

	if c1.String() != c2.String() {
		t.Fatalf("compact mismatch\n%s\n%s", c1.String(), c2.String())
	}

	var i1, i2 bytes.Buffer
	if err := v1.Json.Indent(&i1, raw, "", "  "); err != nil {
		t.Fatal(err)
	}
	if err := v2.Json.Indent(&i2, raw, "", "  "); err != nil {
		t.Fatal(err)
	}

	if i1.String() != i2.String() {
		t.Fatalf("indent mismatch\n%s\n%s", i1.String(), i2.String())
	}
}
