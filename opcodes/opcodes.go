// Package opcodes defines the shared opcode constants for the Planetopia mesh protocol.
// All opcodes must match the corresponding C definitions in c/opcodes.h.
package opcodes

// Serial command opcodes — byte 0 of the data payload in MessageTypeSerialCmdBroadcast frames.
const (
	OpNodeIdSet  = byte(0xC0) // Server → node: assign logical node ID
	OpConfigSet  = byte(0xC1) // Server → node: set adapter type and config
	OpTxPowerSet = byte(0xC2) // Server → node: set TX power preset

	// Output adapter commands (server → output node)
	OpLEDSolid  = byte(0xD0) // Set LED strip to solid colour: [r, g, b] at bytes [1:4]
	OpLEDOff    = byte(0xD1) // Turn LED strip off
	OpLEDBlink  = byte(0xD2) // Blink LED: [r, g, b, interval_ms_hi, interval_ms_lo] at [1:6]
	OpRelaySet  = byte(0xD8) // Set relay state: [0x00=off, 0x01=on] at byte [1]

	// Input adapter events (node → server, carried in OP_COMMAND_ACK or adapter data frames)
	OpCommandAck = byte(0xE0) // Node → server: acknowledge a received command
)
