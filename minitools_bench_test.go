//go:build go1.27

package minitools

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/qmaru/minitools/v2/data/json/gojson"
	"github.com/qmaru/minitools/v2/data/json/sonic"
	standardv1 "github.com/qmaru/minitools/v2/data/json/standard/v1"
	standardv2 "github.com/qmaru/minitools/v2/data/json/standard/v2"
	"github.com/qmaru/minitools/v2/hashx/blake3"
	"github.com/qmaru/minitools/v2/hashx/md5"
	"github.com/qmaru/minitools/v2/hashx/murmur3"
	"github.com/qmaru/minitools/v2/hashx/sha256"
	"github.com/qmaru/minitools/v2/hashx/sha512"
	"github.com/qmaru/minitools/v2/random/nanoid"
	"github.com/qmaru/minitools/v2/random/sqids"
	"github.com/qmaru/minitools/v2/secret/aes/cbc"
	"github.com/qmaru/minitools/v2/secret/aes/gcm"
	"github.com/qmaru/minitools/v2/secret/chacha20"
	"github.com/qmaru/minitools/v2/secret/xor"
)

func init() {
	runtime.GOMAXPROCS(1)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	Profile struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"profile"`

	Tags    []string `json:"tags"`
	History []Item   `json:"history"`
}

type Item struct {
	ItemID int     `json:"item_id"`
	Price  float64 `json:"price"`
}

func GenerateData(count int) []User {
	users := make([]User, count)

	for i := range count {
		u := User{
			ID:    i,
			Name:  "user_" + strconv.Itoa(i),
			Email: "user_" + strconv.Itoa(i) + "@example.com",
			Tags:  []string{"tag1", "tag2", "tag3"},
			History: []Item{
				{ItemID: 1, Price: 100},
				{ItemID: 2, Price: 200},
			},
		}

		u.Profile.City = "City_" + strconv.Itoa(i%100)
		u.Profile.Country = "Country_" + strconv.Itoa(i%10)

		users[i] = u
	}

	return users
}

func GenerateJSONBytes(count int) []byte {
	data := GenerateData(count)
	stdJ := standardv1.New()
	b, _ := stdJ.Json.Marshal(data)
	return b
}

func BenchmarkDataJson(b *testing.B) {
	stdJ := standardv1.New()
	stdJ2 := standardv2.New()
	sonicJ := sonic.New()
	goJ := gojson.New()

	b.ReportAllocs()

	sizes := []int{1, 10, 100, 1000}

	for _, n := range sizes {
		n := n

		b.Run(fmt.Sprintf("%dKB", n), func(b *testing.B) {
			jsonByte := GenerateJSONBytes(n)
			size := int64(len(jsonByte))

			// Unmarshal to struct
			b.Run("Std/Struct", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = stdJ.Json.Unmarshal(jsonByte, &u)
				}
			})

			b.Run("StdV2/Struct", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = stdJ2.Json.Unmarshal(jsonByte, &u)
				}
			})

			b.Run("GoJson/Struct", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = goJ.Json.Unmarshal(jsonByte, &u)
				}
			})

			b.Run("Sonic/Struct", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = sonicJ.Json.Unmarshal(jsonByte, &u)
				}
			})

			// API decode + encode
			b.Run("Std/API", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = stdJ.Json.Unmarshal(jsonByte, &u)
					_, _ = stdJ.Json.Marshal(&u)
				}
			})

			b.Run("StdV2/API", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = stdJ2.Json.Unmarshal(jsonByte, &u)
					_, _ = stdJ2.Json.Marshal(&u)
				}
			})

			b.Run("GoJson/API", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = goJ.Json.Unmarshal(jsonByte, &u)
					_, _ = goJ.Json.Marshal(&u)
				}
			})

			b.Run("Sonic/API", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var u []User
					_ = sonicJ.Json.Unmarshal(jsonByte, &u)
					_, _ = sonicJ.Json.Marshal(&u)
				}
			})

			// map
			b.Run("Std/Map", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var d []map[string]any
					_ = stdJ.Json.Unmarshal(jsonByte, &d)
				}
			})

			b.Run("StdV2/Map", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var d []map[string]any
					_ = stdJ2.Json.Unmarshal(jsonByte, &d)
				}
			})

			b.Run("GoJson/Map", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var d []map[string]any
					_ = goJ.Json.Unmarshal(jsonByte, &d)
				}
			})

			b.Run("Sonic/Map", func(b *testing.B) {
				b.SetBytes(size)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					var d []map[string]any
					_ = sonicJ.Json.Unmarshal(jsonByte, &d)
				}
			})
		})
	}
}

func BenchmarkHashBlake3(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	bhash := blake3.New()

	b.ReportAllocs()

	b.Run("Sum256", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			bhash.Sum256(data)
		}
	})

	b.Run("Sum512", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			bhash.Sum512(data)
		}
	})

	b.Run("StreamSum256", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			h := blake3.New()
			h.SetSize(32)
			h.Write(data)
			h.SumStream()
		}
	})

	b.Run("StreamSum512", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			h := blake3.New()
			h.SetSize(64)
			h.Write(data)
			h.SumStream()
		}
	})
}

func BenchmarkHashMD5(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	mhash := md5.New()

	b.ReportAllocs()

	b.Run("Sum", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			mhash.Sum(data)
		}
	})

	b.Run("StreamSum", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			h := md5.New()
			h.Write(data)
			h.SumStream()
		}
	})
}

func BenchmarkHashSha256(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	shash := sha256.New()

	b.ReportAllocs()

	b.Run("Sum256", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			shash.Sum256(data)
		}
	})

	b.Run("StreamSum256", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			shash := sha256.New()
			shash.Write(data)
			shash.SumStream()
		}
	})
}

func BenchmarkHashSha512(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	shash := sha512.New()

	b.ReportAllocs()

	b.Run("Sum512", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			shash.Sum512(data)
		}
	})

	b.Run("StreamSum512", func(b *testing.B) {
		b.SetBytes(dataLen)
		for i := 0; i < b.N; i++ {
			shash := sha512.New()
			shash.Write(data)
			shash.SumStream()
		}
	})
}

func BenchmarkHashMurmur3(b *testing.B) {
	data := []byte("key=helloworld!!")
	dataLen := int64(len(data))
	bhash := murmur3.New()

	b.ReportAllocs()
	b.SetBytes(dataLen)
	for b.Loop() {
		bhash.Sum64(data)
	}
}

func BenchmarkHashNanoid(b *testing.B) {
	nhash := nanoid.New()

	b.ReportAllocs()
	b.SetBytes(21)
	for b.Loop() {
		nhash.New()
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
	for b.Loop() {
		s.Encode(data)
	}
}

func BenchmarkSecretAes(b *testing.B) {
	aescbc := cbc.New()
	aesgcm := gcm.New()

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
			aescbc.Encrypt(plain, key, iv)
		}
	})

	b.Run("CBCDecrypt", func(b *testing.B) {
		b.SetBytes(cbcDataLen)
		for i := 0; i < b.N; i++ {
			aescbc.Decrypt(cbcData, key, iv)
		}
	})

	b.Run("GCMEncrypt", func(b *testing.B) {
		b.SetBytes(plainLen)
		for i := 0; i < b.N; i++ {
			aesgcm.Encrypt(plain, key)
		}
	})

	b.Run("GCMDecrypt", func(b *testing.B) {
		b.SetBytes(gcmDataLen)
		for i := 0; i < b.N; i++ {
			aesgcm.Decrypt(gcmData, key)
		}
	})
}

func BenchmarkSecretChacha20(b *testing.B) {
	c := chacha20.New()

	plain := []byte("minitools")
	plainLen := int64(len(plain))
	nonce, _ := c.GenerateNonce()

	encData := []byte{17, 53, 44, 213, 99, 96, 75, 168, 10, 6, 99, 178, 26, 150, 207, 112, 40, 50, 73, 200, 125}
	encDataLen := int64(len(encData))

	key := []byte("this is a 32bit key for chacha20")

	b.ReportAllocs()

	b.Run("Encrypt", func(b *testing.B) {
		b.SetBytes(plainLen)
		for i := 0; i < b.N; i++ {
			c.Encrypt(plain, key, nonce)
		}
	})

	b.Run("Decrypt", func(b *testing.B) {
		b.SetBytes(encDataLen)
		for i := 0; i < b.N; i++ {
			c.Decrypt(encData, key)
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
	for b.Loop() {
		xsecret.Cipher(data, key)
	}
}
