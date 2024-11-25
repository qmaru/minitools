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
	"github.com/qmaru/minitools/v2/secret/aes/cbc"
	"github.com/qmaru/minitools/v2/secret/aes/gcm"
	"github.com/qmaru/minitools/v2/secret/xor"
	"github.com/qmaru/minitools/v2/time"
)

func TestAes(t *testing.T) {
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

func TestDataJson(t *testing.T) {
	jdata := gojson.New()
	jsonStr := []byte(`{"name": "Alice", "age": 20}`)
	data, err := jdata.RawJson2Map(jsonStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestFile(t *testing.T) {
	fs := file.New()
	runPath, err := fs.RootPath("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(runPath)
}

func TestTime(t *testing.T) {
	ts := time.New()
	t1 := "2006.01.02 15:04:05"
	t2 := "2006/01/02"
	now := "2020.10.01 14:30:40"
	data, err := ts.AnyFormat(t1, t2, now)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
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

func TestHashBlake3(t *testing.T) {
	bhash := blake3.New()
	s := bhash.Sum256([]byte("123456")).ToBase64()
	t.Log(s)
}

func TestHashNanoid(t *testing.T) {
	nhash := nanoid.New()
	s, err := nhash.New(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestHashMurmur3(t *testing.T) {
	mhash := murmur3.New()
	data := mhash.Sum32([]byte("hello world")).ToUint32()
	t.Log(data)
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
