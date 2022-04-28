package minitools

import (
	"bytes"
	"testing"
)

func TestAes(t *testing.T) {
	as := AESSuite()

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
	ds := DataSuite()
	jsonStr := []byte(`{"name": "Alice", "age": 20}`)
	data, err := ds.RawMap2Map(jsonStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestFile(t *testing.T) {
	fs := FileSuite()
	runPath, err := fs.LocalPath(true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(runPath)
}

func TestTime(t *testing.T) {
	ts := TimeSuite()
	t1 := "2006.01.02 15:04:05"
	t2 := "2006/01/02"
	now := "2020.10.01 14:30:40"
	data, err := ts.AnyFormat(t1, t2, now)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
