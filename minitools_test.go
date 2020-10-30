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

	encData := as.Encrypt(plain, key, iv)
	decData := as.Decrypt(encData, key, iv)

	if bytes.Equal(decData, plain) {
		t.Logf("\nEncrypt: %s\nDecrypt: %s", encData, decData)
	} else {
		t.Error("Decryption failed")
	}
}

func TestData(t *testing.T) {
	ds := DataSuite()
	jsonStr := []byte(`{"name": "Alice", "age": 20}`)
	data := ds.RawMap2Map(jsonStr)
	t.Log(data)
}

func TestFile(t *testing.T) {
	fs := FileSuite()
	runPath := fs.LocalPath(true)
	t.Log(runPath)
}

func TestTime(t *testing.T) {
	ts := TimeSuite()
	t1 := "2006.01.02 15:04:05"
	t2 := "2006/01/02"
	now := "2020.10.01 14:30:40"
	data := ts.AnyFormat(t1, t2, now)
	t.Log(data)
}
