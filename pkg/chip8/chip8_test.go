package chip8

import "testing"

func TestBitExtraction(t *testing.T) {
	c := NewChip8(nil, validTestROM)
	opcode := uint16(0x9F4D)

	nnn, n, x, y, kk := c.extractReferenceBits(opcode)

	if nnn != 0xF4D {
		t.Errorf("nnn was not parsed correctly: %v", nnn)
	}

	if n != 0xD {
		t.Errorf("n was not parsed correctly: %v", n)
	}

	if x != 0xF {
		t.Errorf("x was not parsed correctly: %v", x)
	}

	if y != 0x4 {
		t.Errorf("y was not parsed correctly: %v", y)
	}

	if kk != 0x4D {
		t.Errorf("kk was not parsed correctly: %v", kk)
	}
}

func TestOpRET(t *testing.T) {
	c := NewChip8(nil, validTestROM)

	c.Memory.PushStack(0x026F)
	c.Memory.PushStack(0x031D)
	c.Memory.PushStack(0x0754)

	c.ExecOp(0x00EE)

	if c.Memory.registers.pc != 0x0754 {
		t.Errorf("Failed to set the PC on RET operation")
	}

	if c.Memory.registers.sp != 2 {
		t.Errorf("Failed to decrement the Stack Pointer on RET operation")
	}
}
