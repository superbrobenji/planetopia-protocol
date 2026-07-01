package adapter_test

import (
	"testing"

	"github.com/superbrobenji/lattice-protocol/adapter"
)

func TestIsInput(t *testing.T) {
	if !adapter.IsInput(adapter.TypePIR) {
		t.Error("IsInput(TypePIR) = false, want true")
	}
	if adapter.IsInput(adapter.TypeLED) {
		t.Error("IsInput(TypeLED) = true, want false")
	}
}

func TestIsOutput(t *testing.T) {
	if !adapter.IsOutput(adapter.TypeLED) {
		t.Error("IsOutput(TypeLED) = false, want true")
	}
	if adapter.IsOutput(adapter.TypePIR) {
		t.Error("IsOutput(TypePIR) = true, want false")
	}
}
