// Copyright 2015 Tyson Maly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file tests bit sets

package bitset

import (
	"fmt"
	"testing"
)

func TestBitSetNew(t *testing.T) {
	b := New(0)

	b.Set(15)

	str := b.ToHexString()

	/*
		if str != "0f" {
			t.Errorf("Set then ToHexString failed")
		}
	*/

	fmt.Printf("Set then ToHexString returned %s\n", str)

}

func TestBitSetToHexString(t *testing.T) {
	b := New(15)

	str := b.ToHexString()

	if str != "0f" {
		t.Errorf("ToHexString failed")
	}

	fmt.Printf("ToHexString returned %s\n", str)

}
