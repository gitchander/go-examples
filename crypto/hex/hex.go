package hex

import (
	"bytes"
)

const (
	hexLower = "0123456789abcdef"
	hexUpper = "0123456789ABCDEF"
)

func byteToNibbles(b byte) (hi, lo byte) {
	hi = b >> 4
	lo = b & 0xF
	return
}

func nibblesToByte(hi byte, lo byte) (b byte) {
	b |= hi << 4
	b |= lo & 0x0F
	return
}

func quoRem(x, y int) (quo, rem int) {
	quo = x / y
	rem = x % y
	return
}

func writeByteHex(b byte, bs []byte) {
	hi, lo := byteToNibbles(b)
	bs[0] = hexUpper[hi]
	bs[1] = hexUpper[lo]
}

func HexQuad(bs []byte) string {

	//example return value: "D7A8FBB3 07D78094 69CA9ABC B0082E4F 8D5651E4 6D3CDB76 2D02D0BF 37C9E592"

	quo, rem := quoRem(len(bs), 4)

	buffer := new(bytes.Buffer)

	const spaceChar = ' ' // Space
	k := 0

	p := make([]byte, 9) // format - " AABBCCDD"
	p[0] = spaceChar

	if quo > 0 {
		fill := func(src []byte, dest []byte) {
			writeByteHex(src[0], dest[0:])
			writeByteHex(src[1], dest[2:])
			writeByteHex(src[2], dest[4:])
			writeByteHex(src[3], dest[6:])
		}

		fill(bs[k:k+4], p[1:])
		k += 4
		buffer.Write(p[1:])

		for i := 1; i < quo; i++ {
			fill(bs[k:k+4], p[1:])
			k += 4
			buffer.Write(p)
		}
	}

	if rem > 0 {
		if k > 0 {
			buffer.WriteByte(spaceChar)
		}
		for i := 0; i < rem; i++ {
			writeByteHex(bs[k], p[1:])
			buffer.Write(p[1:3])
			k++
		}
	}

	return string(buffer.Bytes())
}
