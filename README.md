# HashAoC17: A "cryptographic" hash
Copyright (c) 2017 Bart Massey

This codebase contains an implementation of the hash
function from Advent of Code 2017 Day 10 Part 2, along with
an "avalanche" test driver to analyze a cryptographic
property.

Set your GOPATH environment variable to this directory, then
run with

    go run avalanche.go

to get statistics. You can optionally give a 64-bit unsigned
decimal seed argument to the tester to repeat a sequence.

---

This work is licensed under the "MIT License".  Please see
the file `LICENSE` in the source distribution of this
software for license terms.
