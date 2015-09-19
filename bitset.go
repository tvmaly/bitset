/*
Package bitset implements a simple bitset using math/big
the motivation was that I needed a simple way to store
bitstrings as hex values that could easily work on the frontend
side of a web app with javascript

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

type BitSet struct {
	*big.Int
}

func New(x int64) *BitSet {
	return &BitSet{big.NewInt(x)}
}

func (b *BitSet) FromHexString(hexstr string) error {

	var ok bool

	b.Int, ok = b.SetString(hexstr, 16)

	if !ok {
		return errors.New("cannot convert " + hexstr + " to hex value to big int")
	}

	return nil
}

func (b *BitSet) String() string {

	return b.ToHexString()

}

func (b *BitSet) ToHexString() string {

	return fmt.Sprintf("%x", b.Bytes())

}

func (b *BitSet) Set(i int64) {

	b.Int = b.Int.Or(b.Int, big.NewInt(i))

}

func (b *BitSet) SetBit(i int64) {
	b.Int = b.Int.SetBit(b.Int, int(i), 1)
}

func (b *BitSet) UnSetBit(i int64) {
	b.Int = b.Int.SetBit(b.Int, int(i), 0)
}

func (b *BitSet) Get(i int) uint {

	return b.Int.Bit(i)

}
