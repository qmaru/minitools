package minitools

import (
	"runtime"
	"testing"

	"github.com/qmaru/minitools/v2/data/minijson"
	"github.com/qmaru/minitools/v2/hashx/blake3"
	"github.com/qmaru/minitools/v2/hashx/nanoid"
	"github.com/qmaru/minitools/v2/hashx/sqids"
	"github.com/qmaru/minitools/v2/secret/aes"
	"github.com/qmaru/minitools/v2/secret/xor"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func BenchmarkDataJson(b *testing.B) {
	jdata := minijson.New()
	jsonStr := []byte(`{"id":12345,"name":"John Doe","email":"johndoe@example.com","address":{"street":"123 Main St","city":"Springfield","state":"IL","zip":"62701","country":"USA"},"phone_numbers":[{"type":"home","number":"555-1234"},{"type":"work","number":"555-5678"}],"preferences":{"newsletter":true,"notifications":false,"theme":"dark"},"purchase_history":[{"item_id":987,"item_name":"Laptop","price":1299.99,"quantity":1,"purchase_date":"2024-01-15"},{"item_id":654,"item_name":"Headphones","price":199.99,"quantity":2,"purchase_date":"2024-02-01"}],"active":true,"last_login":"2024-11-21T15:30:00Z","meta":{"created_at":"2020-05-01T10:00:00Z","updated_at":"2024-11-20T08:00:00Z","version":"1.2.3"},"friends":[{"id":67890,"name":"Jane Smith","relationship":"friend"},{"id":23456,"name":"Bob Johnson","relationship":"colleague"}]}`)
	jsonStrLen := int64(len(jsonStr))

	b.ReportAllocs()
	b.SetBytes(jsonStrLen)

	for i := 0; i < b.N; i++ {
		jdata.RawJson2Map(jsonStr)
	}
}

func BenchmarkHashBlake3(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	bhash := blake3.New()

	b.ReportAllocs()
	b.SetBytes(dataLen)
	for i := 0; i < b.N; i++ {
		bhash.Sum256(data)
	}
}

func BenchmarkHashSqids(b *testing.B) {
	data := []uint64{123456}
	dataLen := int64(len(data))
	shash := sqids.New()
	s, _ := shash.New(sqids.SqidsOptions{
		MinLength: 10,
	})

	b.ReportAllocs()
	b.SetBytes(dataLen)
	for i := 0; i < b.N; i++ {
		s.Encode(data)
	}
}

func BenchmarkHashNanoid(b *testing.B) {
	nhash := nanoid.New()

	b.ReportAllocs()
	b.SetBytes(21)
	for i := 0; i < b.N; i++ {
		nhash.New()
	}
}

func BenchmarkSecretAes(b *testing.B) {
	cbc := aes.NewCBC()
	gcm := aes.NewGCM()

	plain := []byte("minitools")
	plainLen := int64(len(plain))
	key := []byte("length is 16 bit")
	iv := []byte("same size as key")

	cbcData := []byte{101, 188, 58, 106, 22, 45, 239, 197, 86, 218, 46, 54, 153, 40, 84, 28}
	cbcDataLen := int64(len(cbcData))
	gcmData := []byte{245, 47, 192, 198, 189, 169, 156, 89, 135, 115, 38, 160, 136, 148, 138, 213, 146, 129, 140, 196, 194, 124, 211, 29, 213, 154, 253, 18, 4, 6, 173, 32, 94, 205, 164, 193, 91}
	gcmDataLen := int64(len(gcmData))

	b.ReportAllocs()

	b.Run("CBCEncrypt", func(b *testing.B) {
		b.SetBytes(plainLen)
		for i := 0; i < b.N; i++ {
			cbc.Encrypt(plain, key, iv)
		}
	})

	b.Run("CBCDecrypt", func(b *testing.B) {
		b.SetBytes(cbcDataLen)
		for i := 0; i < b.N; i++ {
			cbc.Decrypt(cbcData, key, iv)
		}
	})

	b.Run("GCMEncrypt", func(b *testing.B) {
		b.SetBytes(plainLen)
		for i := 0; i < b.N; i++ {
			gcm.Encrypt(plain, key)
		}
	})

	b.Run("GCMDecrypt", func(b *testing.B) {
		b.SetBytes(gcmDataLen)
		for i := 0; i < b.N; i++ {
			gcm.Decrypt(gcmData, key)
		}
	})
}

func BenchmarkSecretXor(b *testing.B) {
	data := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	dataLen := int64(len(data))
	key := []byte("this_is_a_16byte")
	xsecret := xor.New()

	b.ReportAllocs()
	b.SetBytes(dataLen)
	for i := 0; i < b.N; i++ {
		xsecret.Cipher(data, key)
	}
}
