// Copyright © 2017 Bart Massey
// This program is licensed under the "MIT License".
// Please see the file LICENSE in this distribution
// for license terms.

// Advent of Code 2017 Day 10 Part 2 hash function.

package hashaoc17;

// In-place reverse a subsequence of the circular buffer
// ring starting at posn and with length length.
func circularReverse(ring *[256]uint8, posn int, length uint8) {
	// Swap start position within the reversal.
	i := 0
	// Loop until the indices meet/cross, which
	// will mean that everything has been reversed.
	for {
		// Swap end position within the reversal.
		j := int(length) - 1 - i
		// If the end has met/crossed the start,
		// get out.
		if j <= i {
			break
		}
		// Swap start and end position within the
		// ring.
		pi := (posn + i) % 256
		pj := (posn + j) % 256
		// Make the swap.
		ring[pi], ring[pj] = ring[pj], ring[pi]
		// Advance the position.
		i++
	}
}

// The hash function from Advent of Code Day 2017 Day 10
// Part 2. Given an input byte slice, return a 16-byte
// "cryptographic" hash of that input.
func HashAoC17(input []uint8) [16]uint8 {
	// The ring.
	var ring [256]uint8
	for i := range ring {
		ring[i] = uint8(i)
	}
	// Position in ring.
	posn := 0
	// Skip count in ring.
	skip := 0
	// Treat the input as ASCII (Unicode)
	// characters.
	// "Extra" characters to stick on the end.
	extras := []uint8 {17, 31, 73, 47, 23}
	// Build the lengths array as specified.
	lengths := make([]uint8, len(input) + len(extras))
	for i, v := range input {
		lengths[i] = v
	}
	for i, v := range extras {
		lengths[len(input) + i] = v
	}
	// Run the rounds.
	for i := 0; i < 64; i++ {
		// Run a round.
		for _, length := range lengths {
			// Do the reversal.
			circularReverse(&ring, posn, length)
			// Advance the position.
			posn = (posn + int(length) + skip) % 256
			// Increment the skip.
			skip++
		}
	}
	// Set up the result.
	var result [16]uint8
	nb := 0
	// Get the checksum for each block.
	for b := 0; b < 256; b += 16 {
		// Can start with 0 since x ^ 0 == x.
		h := uint8(0)
		// Xor in all the characters.
		for i := 0; i < 16; i++ {
			h = h ^ ring[b + i]
		}
		// Store the block
		result[nb] = h
		nb++
	}
	// Return the checksum.
	return result
}
