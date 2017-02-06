package tests

import (
	"testing"
	"github.com/uncleandy/test_bits_count_golang/bit_counter"
)

func TestBitsCounter(t *testing.T) {
	if bit_counter.BitsCount(uint64(0)) != 0 {
		t.Error("Wrong BitsCount(0). Expected 0, got", bit_counter.BitsCount(uint64(0)))
	}
	if bit_counter.BitsCount(uint64(1)) != 1 {
		t.Error("Wrong BitsCount(1). Expected 1, got", bit_counter.BitsCount(uint64(1)))
	}
	if bit_counter.BitsCount(uint64(0xFF)) != 8 {
		t.Error("Wrong BitsCount(0xFF). Expected 8, got", bit_counter.BitsCount(uint64(0xFF)))
	}
	if bit_counter.BitsCount(uint64(0xFFFF)) != 16 {
		t.Error("Wrong BitsCount(0xFFFF). Expected 16, got", bit_counter.BitsCount(uint64(0xFFFF)))
	}
	if bit_counter.BitsCount(uint64(0xFFFFFFFF)) != 32 {
		t.Error("Wrong BitsCount(0xFFFFFFFF). Expected 32, got", bit_counter.BitsCount(uint64(0xFFFFFFFF)))
	}
	if bit_counter.BitsCount(uint64(0xFFFFFFFFFFFFFFFF)) != 64 {
		t.Error("Wrong BitsCount(0xFFFFFFFFFFFFFFFF). Expected 64, got", bit_counter.BitsCount(uint64(0xFFFFFFFFFFFFFFFF)))
	}
	if bit_counter.BitsCount(uint64(0xAAAAAAAAAAAAAAAA)) != 32 {
		t.Error("Wrong BitsCount(0xAAAAAAAAAAAAAAAA). Expected 32, got", bit_counter.BitsCount(uint64(0xAAAAAAAAAAAAAAAA)))
	}
	if bit_counter.BitsCount(uint64(0xD7D7D7D7D7D7D7D7)) != 48 {
		t.Error("Wrong BitsCount(0xD7D7D7D7D7D7D7D7). Expected 48, got", bit_counter.BitsCount(uint64(0xD7D7D7D7D7D7D7D7)))
	}
}
