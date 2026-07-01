package opcodes_test

import (
	"testing"

	"github.com/superbrobenji/lattice-protocol/opcodes"
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

func TestOpHealthReq_Value(t *testing.T) {
	if opcodes.OpHealthReq != 0xB0 {
		t.Errorf("OpHealthReq = 0x%02X, want 0xB0", opcodes.OpHealthReq)
	}
}

func TestOpHealthReport_Value(t *testing.T) {
	if opcodes.OpHealthReport != 0xB1 {
		t.Errorf("OpHealthReport = 0x%02X, want 0xB1", opcodes.OpHealthReport)
	}
}

func TestOpNodeHealth_Value(t *testing.T) {
	if opcodes.OpNodeHealth != 0xB2 {
		t.Errorf("OpNodeHealth = 0x%02X, want 0xB2", opcodes.OpNodeHealth)
	}
}
