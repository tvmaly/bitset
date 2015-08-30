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

	if str != "0f" {
		t.Errorf("Set then ToHexString failed should be 0f")
	}

}

func TestBitSetToHexString(t *testing.T) {
	b := New(15)

	str := b.ToHexString()

	if str != "0f" {
		t.Errorf("ToHexString failed should be 0f")
	}

}

func TestBitSetFromHexString(t *testing.T) {

	b := New(0)

	err := b.FromHexString("0f")

	if err != nil {
		fmt.Printf("error %s\n", err)
		t.Errorf("FromHexString failed")
	}

	str := b.ToHexString()

	if str != "0f" {
		t.Errorf("the FromHexString should equal ToHexString of 0f")

	}

}

func TestBitSetGet(t *testing.T) {

	b := New(8)

	bit := b.Get(3)

	if bit != 1 {
		t.Errorf("Get failed should have returned 1 for 3rd bit in zero based")
	}

}
