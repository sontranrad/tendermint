package common

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	strChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 62 characters
)

func init() {
	// Seed math/rand with "secure" int64
	b := CRandBytes(8)
	var seed uint64
	for i := 0; i < 8; i++ {
		seed |= uint64(b[i])
		seed <<= 8
	}
	rand.Seed(int64(seed))
}

// Constructs an alphanumeric string of given length.
func RandStr(length int) string {
	chars := []byte{}
MAIN_LOOP:
	for {
		val := rand.Int63()
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) // rightmost 6 bits
			if v >= 62 {         // only 62 characters in strChars
				val >>= 6
				continue
			} else {
				chars = append(chars, strChars[v])
				if len(chars) == length {
					break MAIN_LOOP
				}
				val >>= 6
			}
		}
	}

	return string(chars)
}

func RandUInt16() uint16 {
	return uint16(rand.Uint32() & (1<<16 - 1))
}

func RandUInt32() uint32 {
	return rand.Uint32()
}

func RandUInt64() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

// Distributed pseudo-exponentially to test for various cases
func RandUInt16Exp() uint16 {
	bits := rand.Uint32() % 16
	if bits == 0 {
		return 0
	}
	n := uint16(1 << (bits - 1))
	n += uint16(rand.Int31()) & ((1 << (bits - 1)) - 1)
	return n
}

// Distributed pseudo-exponentially to test for various cases
func RandUInt32Exp() uint32 {
	bits := rand.Uint32() % 32
	if bits == 0 {
		return 0
	}
	n := uint32(1 << (bits - 1))
	n += uint32(rand.Int31()) & ((1 << (bits - 1)) - 1)
	return n
}

// Distributed pseudo-exponentially to test for various cases
func RandUInt64Exp() uint64 {
	bits := rand.Uint32() % 64
	if bits == 0 {
		return 0
	}
	n := uint64(1 << (bits - 1))
	n += uint64(rand.Int63()) & ((1 << (bits - 1)) - 1)
	return n
}

func RandTime() time.Time {
	return time.Unix(int64(RandUInt64Exp()), 0)
}

func RandBytes(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = byte(rand.Intn(256))
	}
	return bs
}

//-----------------------------------------------------------------------------
// CRand* methods are crypto safe.

func CRandBytes(numBytes int) []byte {
	b := make([]byte, numBytes)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// RandHex(24) gives 96 bits of randomness, strong enough for most purposes.
func CRandHex(numDigits int) string {
	return hex.EncodeToString(CRandBytes(numDigits / 2))
}
