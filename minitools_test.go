package minitools

import (
	"bytes"
	"testing"

	"github.com/qmaru/minitools/v2/data/json/gojson"
	"github.com/qmaru/minitools/v2/file"
	"github.com/qmaru/minitools/v2/hashx/blake3"
	"github.com/qmaru/minitools/v2/hashx/murmur3"
	"github.com/qmaru/minitools/v2/hashx/nanoid"
	"github.com/qmaru/minitools/v2/hashx/sqids"
	"github.com/qmaru/minitools/v2/hashx/text"
	"github.com/qmaru/minitools/v2/secret/aes/cbc"
	"github.com/qmaru/minitools/v2/secret/aes/gcm"
	"github.com/qmaru/minitools/v2/secret/chacha20"
	"github.com/qmaru/minitools/v2/secret/xor"
	"github.com/qmaru/minitools/v2/time"
)

func TestDataJson(t *testing.T) {
	jdata := gojson.New()
	jsonStr := []byte(`{"name": "Alice", "age": 20}`)

	// Decoder
	var d map[string]any
	decoder := jdata.Json.NewDecoder(bytes.NewBuffer(jsonStr))
	err := decoder.Decode(&d)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decode: %v", d)

	// Encoder
	var buf bytes.Buffer
	encoder := jdata.Json.NewEncoder(&buf)
	err = encoder.Encode(d)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Encode: %v", buf.String())

	// Common method
	data, err := jdata.RawJson2Map(jsonStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("RawJson2Map: %v", data)
}

func TestFile(t *testing.T) {
	fs := file.New()
	runPath, err := fs.RootPath("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(runPath)
}

func TestHashBlake3(t *testing.T) {
	bhash := blake3.New()
	s := bhash.Sum256([]byte("123456")).ToBase64()
	t.Log(s)
}

func TestHashMurmur3(t *testing.T) {
	mhash := murmur3.New()
	data := mhash.Sum32([]byte("hello world")).ToUint32()
	t.Log(data)
}

func TestHashNanoid(t *testing.T) {
	nhash := nanoid.New()
	s, err := nhash.New(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestHashSqids(t *testing.T) {
	shash := sqids.New()
	s, err := shash.New(sqids.SqidsOptions{
		MinLength: 10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s.Encode([]uint64{123456}))
}

func TestHashText(t *testing.T) {
	thash := text.New()
	s := "123456"

	t.Logf("raw: %s\n", s)

	// base64
	b64e := thash.Base64Encode([]byte(s))
	t.Logf("b64e: %s\n", b64e)

	b64ding, err := thash.Base64Decoding(b64e)
	if err != nil {
		t.Fatal(err)
	}
	b64s := b64ding.DecodeString()
	t.Logf("b64s: %v\n", b64s)
	b64r := b64ding.DecodeRaw()
	t.Logf("b64r: %v\n", b64r)

	// base62
	b62e := thash.Base62Encode([]byte(s))
	t.Logf("b62e: %s\n", b62e)

	b62ding, err := thash.Base62Decoding(b62e)
	if err != nil {
		t.Fatal(err)
	}
	b62s := b62ding.DecodeString()
	t.Logf("b62s: %v\n", b62s)
	b62r := b62ding.DecodeRaw()
	t.Logf("b62r: %v\n", b62r)

	// hex
	hexe := thash.HexEncode([]byte(s))
	t.Logf("hexe: %s\n", hexe)

	hexding, err := thash.HexDecoding(hexe)
	if err != nil {
		t.Fatal(err)
	}
	hexs := hexding.DecodeString()
	t.Logf("hexs: %v\n", hexs)
	hexr := hexding.DecodeRaw()
	t.Logf("hexr: %v\n", hexr)

	// Nonce
	n, err := thash.Nonce(16)
	if err != nil {
		t.Fatal(err)
	}
	nonceB64 := thash.Base64Encode(n)
	t.Logf("nonceB64: %v\n", nonceB64)
	nonceB62 := thash.Base62Encode(n)
	t.Logf("nonceB62: %v\n", nonceB62)
	nonceHex := thash.HexEncode(n)
	t.Logf("nonceHex: %v\n", nonceHex)
}

func TestSecretAes(t *testing.T) {
	cbc := cbc.New()
	gcm := gcm.New()

	plain := []byte("minitools")
	key := []byte("length is 16 bit")
	iv := []byte("same size as key")

	encCBCData, err := cbc.Encrypt(plain, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	decCBCData, err := cbc.Decrypt(encCBCData, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	encGCMData, err := gcm.Encrypt(plain, key)
	if err != nil {
		t.Fatal(err)
	}

	decGCMData, err := gcm.Decrypt(encGCMData, key)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(decCBCData, plain) {
		t.Logf("CBC:\nEncrypt: %v\nDecrypt: %s", encCBCData, decCBCData)
	} else {
		t.Error("Decryption failed")
	}

	if bytes.Equal(decGCMData, plain) {
		t.Logf("GCM\nEncrypt: %v\nDecrypt: %s", encGCMData, decGCMData)
	} else {
		t.Error("Decryption failed")
	}
}

func TestSecretChacha20(t *testing.T) {
	c := chacha20.New()

	plain := []byte("minitools")
	key := []byte("this is a 32bit key for chacha20")
	nonce, _ := c.GenerateNonce()

	encData, err := c.Encrypt(plain, key, nonce)
	if err != nil {
		t.Fatal(err)
	}

	decData, err := c.Decrypt(encData, key)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(decData, plain) {
		t.Logf("Encrypt: %v\nDecrypt: %s\n", encData, decData)
	} else {
		t.Error("Decryption failed")
	}
}

func TestSecretXor(t *testing.T) {
	data := []byte("hello world")
	key := []byte("123456")

	x := xor.New()
	cipherData := x.Cipher(data, key)

	plainData := x.Cipher(cipherData, key)

	if x.ToString(plainData) == string(data) {
		t.Logf("GCM\nEncrypt: %s\nDecrypt: %s", x.ToString(cipherData), x.ToString(plainData))
	} else {
		t.Error("Decryption failed")
	}
}

func TestTime(t *testing.T) {
	ts := time.New()
	t1 := "2006.01.02 15:04:05"
	t2 := "2006/01/02"
	now := "2020.10.01 14:30:40"
	data, err := ts.ConvertFormat(t1, t2, now)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
