// Package adapter defines the input/output adapter type constants for the Lattice mesh.
// Input adapters send events to the server (e.g. PIR motion sensor).
// Output adapters receive commands from the server (e.g. LED strip, relay).
// All constants must match the corresponding C definitions in c/adapter_types.h.
package adapter

// Adapter type identifiers — carried in health report frames and config set commands.
const (
	TypeUnknown = int32(0) // not yet configured
	TypeSerial  = int32(1) // serial management (internal)
	TypePIR     = int32(2) // passive infrared motion sensor (INPUT)
	TypeLED     = int32(3) // LED strip (OUTPUT)
	TypeRelay   = int32(4) // relay switch (OUTPUT)
)

// IsInput returns true for adapter types that send events to the server.
func IsInput(t int32) bool {
	return t == TypePIR
}

// IsOutput returns true for adapter types that receive commands from the server.
func IsOutput(t int32) bool {
	return t == TypeLED || t == TypeRelay
}
