package bigInt

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigInt struct {
	*big.Int
}

func NewBigInt(i int64) BigInt {
	return BigInt{big.NewInt(0).SetInt64(i)}
}

func NewIntUnsigned(i uint64) BigInt {
	return BigInt{big.NewInt(0).SetUint64(i)}
}

func NewFromGo(i *big.Int) BigInt {
	return BigInt{big.NewInt(0).Set(i)}
}

func Zero() BigInt {
	return NewBigInt(0)
}

// PositiveFromUnsignedBytes interprets b as the bytes of a big-endian unsigned
// integer and returns a positive Int with this absolute value.
func PositiveFromUnsignedBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{i}
}

// MustFromString convers dec string into big integer and panics if conversion
// is not sucessful.
func MustFromString(s string) BigInt {
	v, err := FromString(s)
	if err != nil {
		panic(err)
	}
	return v
}

func FromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{v}, nil
}

func (bi BigInt) Copy() BigInt {
	return BigInt{Int: new(big.Int).Set(bi.Int)}
}

func Product(ints ...BigInt) BigInt {
	p := NewBigInt(1)
	for _, i := range ints {
		p = Mul(p, i)
	}
	return p
}

func Mul(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Mul(a.Int, b.Int)}
}

func Div(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Div(a.Int, b.Int)}
}

func Mod(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Mod(a.Int, b.Int)}
}

func Add(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Add(a.Int, b.Int)}
}

func Sum(ints ...BigInt) BigInt {
	sum := Zero()
	for _, i := range ints {
		sum = Add(sum, i)
	}
	return sum
}

func Subtract(num1 BigInt, ints ...BigInt) BigInt {
	sub := num1
	for _, i := range ints {
		sub = Sub(sub, i)
	}
	return sub
}

func Sub(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Sub(a.Int, b.Int)}
}

//  Returns a**e unless e <= 0 (in which case returns 1).
func Exp(a BigInt, e BigInt) BigInt {
	return BigInt{big.NewInt(0).Exp(a.Int, e.Int, nil)}
}

// Returns x << n
func Lsh(a BigInt, n uint) BigInt {
	return BigInt{big.NewInt(0).Lsh(a.Int, n)}
}

// Returns x >> n
func Rsh(a BigInt, n uint) BigInt {
	return BigInt{big.NewInt(0).Rsh(a.Int, n)}
}

func BitLen(a BigInt) uint {
	return uint(a.Int.BitLen())
}

func Min(x, y BigInt) BigInt {
	// taken from max.Min()
	if x.Equals(Zero()) && x.Equals(y) {
		if x.Sign() != 0 {
			return x
		}
		return y
	}
	if x.LessThan(y) {
		return x
	}
	return y
}

func Cmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

// LessThan returns true if bi < o
func (bi BigInt) LessThan(o BigInt) bool {
	return Cmp(bi, o) < 0
}

// LessThanEqual returns true if bi <= o
func (bi BigInt) LessThanEqual(o BigInt) bool {
	return bi.LessThan(o) || bi.Equals(o)
}

// GreaterThan returns true if bi > o
func (bi BigInt) GreaterThan(o BigInt) bool {
	return Cmp(bi, o) > 0
}

// GreaterThanEqual returns true if bi >= o
func (bi BigInt) GreaterThanEqual(o BigInt) bool {
	return bi.GreaterThan(o) || bi.Equals(o)
}

// Neg returns the negative of bi.
func (bi BigInt) Neg() BigInt {
	return BigInt{big.NewInt(0).Neg(bi.Int)}
}

// Abs returns the absolute value of bi.
func (bi BigInt) Abs() BigInt {
	if bi.GreaterThanEqual(Zero()) {
		return bi.Copy()
	}
	return bi.Neg()
}

// Equals returns true if bi == o
func (bi BigInt) Equals(o BigInt) bool {
	return Cmp(bi, o) == 0
}

func (bi *BigInt) MarshalJSON() ([]byte, error) {
	if bi.Int == nil {
		zero := Zero()
		return json.Marshal(zero)
	}
	return json.Marshal(bi.String())
}

func (bi *BigInt) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return fmt.Errorf("failed to parse big string: '%s'", string(b))
	}

	bi.Int = i
	return nil
}

func (bi *BigInt) Bytes() ([]byte, error) {
	if bi.Int == nil {
		return []byte{}, fmt.Errorf("failed to convert to bytes, big is nil")
	}

	switch {
	case bi.Sign() > 0:
		return append([]byte{0}, bi.Int.Bytes()...), nil
	case bi.Sign() < 0:
		return append([]byte{1}, bi.Int.Bytes()...), nil
	default: //  bi.Sign() == 0:
		return []byte{}, nil
	}
}

func FromBytes(buf []byte) (BigInt, error) {
	if len(buf) == 0 {
		return NewBigInt(0), nil
	}

	var negative bool
	switch buf[0] {
	case 0:
		negative = false
	case 1:
		negative = true
	default:
		return Zero(), fmt.Errorf("big int prefix should be either 0 or 1, got %d", buf[0])
	}

	i := big.NewInt(0).SetBytes(buf[1:])
	if negative {
		i.Neg(i)
	}

	return BigInt{i}, nil
}

func (bi *BigInt) MarshalBinary() ([]byte, error) {
	if bi.Int == nil {
		zero := Zero()
		return zero.Bytes()
	}
	return bi.Bytes()
}

func (bi *BigInt) UnmarshalBinary(buf []byte) error {
	i, err := FromBytes(buf)
	if err != nil {
		return err
	}

	*bi = i

	return nil
}

func (bi *BigInt) IsZero() bool {
	return bi.Int.Sign() == 0
}

func (bi *BigInt) Nil() bool {
	return bi.Int == nil
}

func (bi *BigInt) NilOrZero() bool {
	return bi.Int == nil || bi.Int.Sign() == 0
}
