package hex

import (
	"bytes"
)

const (
	hexLower = "0123456789abcdef"
	hexUpper = "0123456789ABCDEF"
)

func HiNibble(b byte) byte {
	return b >> 4
}

func LoNibble(b byte) byte {
	return b & 0x0F
}

func NibblesToByte(hiNibble byte, loNibble byte) byte {
	return (hiNibble << 4) | (loNibble & 0x0F)
}

func QuoRem(x, y int) (q, r int) {
	q = x / y
	r = x - q*y
	return
}

func HexQuad(bs []byte) string {

	//example return value: "D7A8FBB3 07D78094 69CA9ABC B0082E4F 8D5651E4 6D3CDB76 2D02D0BF 37C9E592"

	q, r := QuoRem(len(bs), 4)

	buffer := new(bytes.Buffer)

	const spaceChar = ' ' // Space
	k := 0

	if q > 0 {

		p := make([]byte, 9) // format - " AABBCCDD"
		p[0] = spaceChar

		fill := func(src []byte, dest []byte) {
			dest[0] = hexUpper[HiNibble(src[0])]
			dest[1] = hexUpper[LoNibble(src[0])]

			dest[2] = hexUpper[HiNibble(src[1])]
			dest[3] = hexUpper[LoNibble(src[1])]

			dest[4] = hexUpper[HiNibble(src[2])]
			dest[5] = hexUpper[LoNibble(src[2])]

			dest[6] = hexUpper[HiNibble(src[3])]
			dest[7] = hexUpper[LoNibble(src[3])]
		}

		fill(bs[k:k+4], p[1:])
		k += 4
		buffer.Write(p[1:])

		for i := 1; i < q; i++ {

			fill(bs[k:k+4], p[1:])
			k += 4
			buffer.Write(p)
		}
	}

	if r > 0 {
		if k > 0 {
			buffer.WriteByte(spaceChar)
		}

		for i := 0; i < r; i++ {
			buffer.WriteByte(hexUpper[HiNibble(bs[k])])
			buffer.WriteByte(hexUpper[LoNibble(bs[k])])
			k++
		}
	}

	return string(buffer.Bytes())
}
