/*
Package bitset implements a simple bitset using math/big
the motivation was that I needed a simple way to store
bitstrings as hex values that could easily work on the frontend
side of a web app with javascript.  Since BitSet is composed of a big.Int
you can use methods like Int64 to get back a decimal representation of the bitset

Example use:

	import "bitset"

	b := bitset.New(0)

	b = b.FromHexString("0f")

*/

package bitset

import (
	"errors"
	"fmt"
	"math/big"
)

// type BitSet wraps a *big.Int
type BitSet struct {
	*big.Int
}

// New(x int64) *BitSet given a int64 return a new bitset represented by this number
func New(x int64) *BitSet {
	return &BitSet{big.NewInt(x)}
}

// FromHexString(hexstr string) error load a bitset into the existing bitset instance using a hexadecimal value
func (b *BitSet) FromHexString(hexstr string) error {

	var ok bool

	b.Int, ok = b.SetString(hexstr, 16)

	if !ok {
		return errors.New("cannot convert " + hexstr + " to hex value to big int")
	}

	return nil
}

// String() string implements Stringer interface returns the hex value of the bitset
func (b *BitSet) String() string {

	return b.ToHexString()

}

// ToHexString() string returns the hex representation of the bit set
func (b *BitSet) ToHexString() string {

	return fmt.Sprintf("%x", b.Bytes())

}

// ToInt() return int version if possible
func (b *BitSet) ToInt() int {

	return int(b.Int64())

}

// Set(i int64) sets the corresponding set of bits by ORing the decimal value with the set
func (b *BitSet) Set(i int64) {

	b.Int = b.Int.Or(b.Int, big.NewInt(i))

}

// SetBit(i int64) sets a single bit in the set
func (b *BitSet) SetBit(i int64) {
	b.Int = b.Int.SetBit(b.Int, int(i), 1)
}

// UnSetBit(i int64) unsets a single bit in the set
func (b *BitSet) UnSetBit(i int64) {
	b.Int = b.Int.SetBit(b.Int, int(i), 0)
}

// Get(i int) uint gets the bit value in the set
func (b *BitSet) Get(i int) uint {

	return b.Int.Bit(i)

}
