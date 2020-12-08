package chip8

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMemorySetGet(t *testing.T) {
	m := NewMemory()
	const loc = 10
	const expectedVal = 0xF8

	m.Set(loc, expectedVal)

	val := m.Get(loc)
	if val != expectedVal {
		t.Errorf("Memory at location %d was incorrect.", loc)
	}
}

func TestMemoryGetTwoBytes(t *testing.T) {
	m := NewMemory()
	const loc = 200

	m.Set(loc, 0xF1)
	m.Set(loc+1, 0xF2)

	twoBytes := m.GetTwoBytes(loc)

	if twoBytes != 0xF1F2 {
		t.Errorf("Two bytes at location %d were incorrect.", loc)
	}
}

func TestStackPushPop(t *testing.T) {
	m := NewMemory()
	const addr = 0xF0F

	m.PushStack(addr)
	if m.sp != 1 {
		t.Errorf("Stack Pointer did not Increment upon push")
	}

	if m.PopStack() != addr {
		t.Errorf("Stack did not store correct memory address")

		if m.sp != 0 {
			t.Errorf("Stack Pointer did not decrement upon pop")
		}
	}
}

func TestGetNBytes(t *testing.T) {
	m := NewMemory()

	const numBytes = 5
	const fromAddress = 0x05

	var bytes [numBytes]byte

	for i := 0; i < numBytes; i++ {
		bytes[i] = m.ram[fromAddress+i]
	}

	fmt.Printf("Bytes collected: %s", hex.EncodeToString(bytes[:]))

	if bytes[0] != 0x20 || bytes[1] != 0x60 || bytes[2] != 0x20 || bytes[3] != 0x20 || bytes[4] != 0x70 {
		t.Errorf("The bytes extracted did not match the bytes expected")
	}
}
