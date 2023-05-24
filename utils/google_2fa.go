package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

var (
	google2FASecretTable = []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', // 7
		'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', // 15
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', // 23
		'Y', 'Z', '2', '3', '4', '5', '6', '7', // 31
	}
)

func NewSecret() string {
	l := len(google2FASecretTable)
	secret := make([]byte, 16)
	for i := range secret {
		secret[i] = google2FASecretTable[rand.Intn(l)]
	}
	return string(secret)
}
func Google2FACode(key string) (string, error) {
	return google2FACode(key, time.Now().Unix())
}
func google2FACode(key string, t int64) (string, error) {
	hs, e := hmacSha1(key, t/30)
	if e != nil {
		return "", e
	}
	snum := lastBit4byte(hs)
	d := snum % 1000000
	return fmt.Sprintf("%06d", d), nil
}

func lastBit4byte(hmacSha1 []byte) int32 {
	if len(hmacSha1) != sha1.Size {
		return 0
	}
	offsetBits := int8(hmacSha1[len(hmacSha1)-1]) & 0x0f
	p := (int32(hmacSha1[offsetBits]) << 24) |
		(int32(hmacSha1[offsetBits+1]) << 16) |
		(int32(hmacSha1[offsetBits+2]) << 8) |
		(int32(hmacSha1[offsetBits+3]) << 0)
	return (p & 0x7fffffff)
}

func hmacSha1(key string, t int64) ([]byte, error) {
	decodeKey, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(key)
	if err != nil {
		return nil, err
	}

	cData := make([]byte, 8)
	binary.BigEndian.PutUint64(cData, uint64(t))

	h1 := hmac.New(sha1.New, decodeKey)
	_, e := h1.Write(cData)
	if e != nil {
		return nil, e
	}
	return h1.Sum(nil), nil
}
