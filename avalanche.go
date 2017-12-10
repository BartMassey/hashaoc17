// Copyright © 2017 Bart Massey
// This program is licensed under the "MIT License".
// Please see the file LICENSE in this distribution
// for license terms.

// Avalanche test for Advent of Code 2017 Day 10 hash.

package main

import (
	"fmt"
	"hashaoc17"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func explode(input []uint8) []int {
	result := make([]int, 8 * len(input))
	for i, v := range input {
		for j := 0; j < 8; j++ {
			if ((v >> uint(j)) & 1) == 1 {
				result[8 * i + j] = 1
			}
		}
	}
	return result
}

func implode(input []int) []uint8 {
	if len(input) % 8 != 0 {
		panic("uneven input")
	}
	result := make([]uint8, len(input) / 8)
	for i, v := range input {
		switch v {
		case 0:
			// Do nothing.
		case 1:
			result[i / 8] |= 1 << uint(i % 8)
		default:
			panic("bad bit")
		}
	}
	return result
}

func xor(x [16]uint8, y [16]uint8) [16]uint8 {
	var result [16]uint8
	for i, v := range x {
		result[i] = v ^ y[i]
	}
	return result
}

func accum(x []int, y []int) {
	if len(x) != len(y) {
		panic("mismatched lengths")
	}
	for i, v := range y {
		x[i] += v
	}
}

func sum(input []int) int {
	t := 0
	for _, v := range input {
		t += v
	}
	return t
}

func randSample() []uint8 {
	sample := make([]uint8, 16)
	for i := range sample {
		sample[i] = uint8(rand.Intn(256))
	}
	return sample
}

func main() {
	if len(os.Args) > 1 {
		seed, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		rand.Seed(seed)
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	nrounds := 128
	nflipped := 0
	var bit_flips [128]int
	for r := 0; r < nrounds; r++ {
		sample := randSample()
		h0 := hashaoc17.HashAoC17(sample)
		sampleBits := explode(sample)
		for i := range sampleBits {
			sampleBits[i] = 1 - sampleBits[i]
			bs := implode(sampleBits)
			h := hashaoc17.HashAoC17(bs)
			x := xor(h0, h)
			xx := explode(x[:])
			accum(bit_flips[:], xx)
			nflipped += sum(xx)
			sampleBits[i] = 1 - sampleBits[i]
		}
	}
	fmt.Printf("aflip: %.03f\n",
		float64(nflipped) / 128.0 / float64(nrounds))
	for i, v := range bit_flips {
		fmt.Printf("bit %03d: %.03f\n",
			i, float64(v) / 128.0 / float64(nrounds))
	}
}
