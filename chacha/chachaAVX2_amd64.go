// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// +build go1.7, amd64, !gccgo, !appengine, !go1.6

package chacha

// go1.7 is still beta (and this AVX2 implementation is experimenal) disabled
var useAVX2 = supportAVX2() == 1 && false

// xorBlocks crypts full block ( len(src) - (len(src) mod 64) bytes ) from src to
// dst using the state. Src and dst may be the same slice but otherwise should not
// overlap. This function increments the counter of state.
func xorBlocks(dst, src []byte, state *[64]byte, rounds int) {
	if useAVX2 && len(src) >= 128 {
		xorBlocksAVX2(dst, src, state, rounds)
	} else if useSSSE3 {
		xorBlocksSSSE3(dst, src, state, rounds)
	} else {
		xorBlocksSSE2(dst, src, state, rounds)
	}
}

// supportAVX2 returns 1 if the runtime (the executing machine) supports AVX2.
//go:noescape
func supportAVX2() int

// xorBlocksAVX2 crypts full block ( len(src) - (len(src) mod 64) bytes ) from src to
// dst using the state.
//go:noescape
func xorBlocksAVX2(dst, src []byte, state *[64]byte, rounds int)
