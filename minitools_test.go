package minitools

import (
	"bytes"
	"testing"

	"github.com/qmaru/minitools/v2/aes"
	"github.com/qmaru/minitools/v2/data"
	"github.com/qmaru/minitools/v2/file"
	"github.com/qmaru/minitools/v2/time"
)

func TestAes(t *testing.T) {
	as := aes.New()

	plain := []byte("minitools")
	key := []byte("length is 16 bit")
	iv := []byte("same size as key")

	encData, err := as.Encrypt(plain, key, iv)
	if err != nil {
		t.Fatal(err)
	}
	decData, err := as.Decrypt(encData, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(decData, plain) {
		t.Logf("\nEncrypt: %s\nDecrypt: %s", encData, decData)
	} else {
		t.Error("Decryption failed")
	}
}

func TestData(t *testing.T) {
	ds := data.New()
	jsonStr := []byte(`{"name": "Alice", "age": 20}`)
	data, err := ds.RawJson2Map(jsonStr)
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
