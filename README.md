# bitset
a bitset in Go to work with bitstrings represented as hexadecimal values

the motivation was a simple bitstring that could also be used on the frontend javascript side 

to keep track of binary flag values

Example:

b := bitset.New(0)

// set bit 15

b.Set(15)

// returns "0f"

hexstring := b.ToHexString()

// returns 1  note things start at zero so this is the 4th bit

bit := b.Get(3)

a := bitset.New(0)

err := a.FromHexString("0f")

if err != nil {

}



