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
	//"strconv"
	//"strings"
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

func (b *BitSet) ToHexString() string {

	return fmt.Sprintf("%x", b.Bytes())

}

func (b *BitSet) Set(i int64) {

	b.Int = b.Int.Or(b.Int, big.NewInt(i))

}

// return 0, errors.New("Unmarshalling error: type mismatch")