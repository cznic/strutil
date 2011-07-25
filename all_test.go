// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz


package strutil

import (
	"bytes"
	"github.com/cznic/mathutil"
	"math"
	"testing"
)

func TestBase64(t *testing.T) {
	const max = 768
	r, err := mathutil.NewFC32(math.MinInt32, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	bin := []byte{}
	for i := 0; i < max; i++ {
		bin = append(bin, byte(r.Next()))
		cmp, err := Base64Decode(Base64Encode(bin))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(bin, cmp) {
			t.Fatalf("a: % x\nb: % x", bin, cmp)
		}
	}
}

func TestBase32Ext(t *testing.T) {
	const max = 640
	r, err := mathutil.NewFC32(math.MinInt32, math.MaxInt32, true)
	if err != nil {
		t.Fatal(err)
	}

	bin := []byte{}
	for i := 0; i < max; i++ {
		bin = append(bin, byte(r.Next()))
		cmp, err := Base32ExtDecode(Base32ExtEncode(bin))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(bin, cmp) {
			t.Fatalf("a: % x\nb: % x", bin, cmp)
		}
	}
}
