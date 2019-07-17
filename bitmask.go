// Package bitmask provides BitMask, BitMask64, and BitMask8 types, with functional getters and setters.
package bitmask

// BitMask is a 32bit mask, suitable for most mask operations
type BitMask uint32

// Set adds the flag to the mask
func (b *BitMask) Set(flag BitMask) {
	*b |= flag
}

// Clear removes the flag from the mask
func (b *BitMask) Clear(flag BitMask) {
	*b ^= flag
}

// IsSet returns true if the flag is set in the mask
func (b *BitMask) IsSet(flag BitMask) bool {
	return *b&flag != 0
}

// BitMask64 is a 64bit mask, suitable for large mask operations
type BitMask64 uint64

// Set adds the flag to the mask
func (b *BitMask64) Set(flag BitMask64) {
	*b |= flag
}

// Clear removes the flag from the mask
func (b *BitMask64) Clear(flag BitMask64) {
	*b ^= flag
}

// IsSet returns true if the flag is set in the mask
func (b *BitMask64) IsSet(flag BitMask64) bool {
	return *b&flag != 0
}

// BitMask8 is an 8bit mask, suitable for small mask operations
type BitMask8 uint8

// Set adds the flag to the mask
func (b *BitMask8) Set(flag BitMask8) {
	*b |= flag
}

// Clear removes the flag from the mask
func (b *BitMask8) Clear(flag BitMask8) {
	*b ^= flag
}

// IsSet returns true if the flag is set in the mask
func (b *BitMask8) IsSet(flag BitMask8) bool {
	return *b&flag != 0
}
