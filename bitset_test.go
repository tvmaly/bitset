// Copyright 2016 Tyson Maly. All rights reserved.
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

	str := b.String()

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

func TestBitInt64(t *testing.T) {

	b := New(0)

	err := b.FromHexString("0f")

	if err != nil {
		fmt.Printf("error %s\n", err)
		t.Errorf("FromHexString failed")
	}

	i := b.Int64()

	if i != 15 {
		t.Errorf("the Int64 should return int64 representation 15")
	}

	fmt.Printf("Int is %d\n", i)

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

func TestTwoBits(t *testing.T) {

	b := New(0)

	b.SetBit(2)
	b.SetBit(5)

	str := b.ToHexString()

	if str != "24" {
		t.Errorf("the ToHexString should return 07")

	}

	bit0 := b.Get(0)
	bit1 := b.Get(1)
	bit2 := b.Get(2)
	bit5 := b.Get(5)

	if bit0 != 0 {
		fmt.Println("zero bit not zero!")
		t.Errorf("Get failed should have returned 0 for zero bit in 07 in zero based")
	}
	if bit1 != 0 {
		fmt.Println("first bit not zero!")
		t.Errorf("Get failed should have returned 0 for first bit in 07 in zero based")
	}
	if bit2 != 1 {
		fmt.Println("second bit not one!")
		t.Errorf("Get failed should have returned 1 for second bit in 07 in zero based")
	}
	if bit5 != 1 {
		fmt.Println("fifth bit not one!")
		t.Errorf("Get failed should have returned 1 for fifth bit in 07 in zero based")
	}
}
