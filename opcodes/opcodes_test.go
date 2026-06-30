package opcodes_test

import (
	"testing"

	"github.com/superbrobenji/planetopia-protocol/opcodes"
)

func TestOpNodeIdSet_Value(t *testing.T) {
	if opcodes.OpNodeIdSet != 0xC0 {
		t.Errorf("OpNodeIdSet = 0x%02X, want 0xC0", opcodes.OpNodeIdSet)
	}
}

func TestOpCommandAck_Value(t *testing.T) {
	if opcodes.OpCommandAck != 0xE0 {
		t.Errorf("OpCommandAck = 0x%02X, want 0xE0", opcodes.OpCommandAck)
	}
}

func TestOpLEDSolid_Value(t *testing.T) {
	if opcodes.OpLEDSolid != 0xD0 {
		t.Errorf("OpLEDSolid = 0x%02X, want 0xD0", opcodes.OpLEDSolid)
	}
}
