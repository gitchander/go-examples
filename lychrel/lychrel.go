package lychrel

import "bytes"

const digs = "0123456789abcdefghijklmnopqrstuvwxyz"

type carryTemp struct {
	carry byte
	temp  byte
}

var TableBase10 = NewTable(10)

type ctTable struct {
	base int
	cts  []*carryTemp
}

func NewTable(base int) *ctTable {

	if (base <= 0) || (base > len(digs)) {
		return nil
	}

	cts := make([]*carryTemp, base*2)
	var quotient, remainder uint64

	for i, _ := range cts {
		quotient, remainder = divmod(uint64(i), uint64(base))
		cts[i] = &carryTemp{
			carry: byte(quotient),
			temp:  byte(remainder),
		}
	}

	return &ctTable{
		base: base,
		cts:  cts,
	}
}

func divmod(dividend, divisor uint64) (quotient, remainder uint64) {

	quotient = dividend / divisor
	remainder = dividend - quotient*divisor

	return
}

type Number interface {
	IsPalindrome() bool
	ReverseThenAdd()
	String() string
	CountDigits() int
}

type privNumber struct {
	b []byte
	t *ctTable
}

func NewNumber(val uint64) Number {
	return NewNumberTable(TableBase10, val)
}

func NewNumberTable(t *ctTable, val uint64) Number {

	var (
		bs        []byte
		remainder uint64
	)

	ubase := uint64(t.base)

	for val > 0 {
		val, remainder = divmod(val, ubase)
		bs = append(bs, byte(remainder))
	}

	return &privNumber{
		b: bs,
		t: t,
	}
}

func (this *privNumber) IsPalindrome() bool {

	b := this.b

	i, j := 0, len(b)-1
	for i < j {

		if b[i] != b[j] {
			return false
		}

		i, j = i+1, j-1
	}

	return true
}

func (this *privNumber) reverseThenAdd_1() {

	b := this.b
	n := len(b)
	cts := this.t.cts

	var (
		temp, carry byte
		ct          *carryTemp
	)

	i, j := 0, n-1
	for i < j {
		temp = b[i] + b[j]
		b[i], b[j] = temp, temp
		i, j = i+1, j-1
	}

	if i == j {
		b[i] += b[i]
	}

	carry = 0
	for i = range b {
		temp = b[i] + carry
		ct = cts[temp]
		b[i] = ct.temp
		carry = ct.carry
	}

	if carry > 0 {
		c := cap(this.b)
		if n < c {
			this.b = this.b[:n+1]
		} else {
			a := make([]byte, n+1, c*2)
			copy(a[:n], this.b)
			this.b = a
		}
		this.b[n] = carry
	}
}

func (this *privNumber) reverseThenAdd_2() {

	b := this.b
	n := len(b)
	cts := this.t.cts

	var (
		temp, carry byte
		ct          *carryTemp
	)

	carry = 0

	i, j := 0, n-1
	for i < j {

		d := b[i] + b[j]
		b[i], b[j] = d, d

		temp = b[i] + carry
		ct = cts[temp]
		b[i] = ct.temp
		carry = ct.carry

		i, j = i+1, j-1
	}

	if i == j {
		b[i] += b[i]
	}

	for i < n {

		temp = b[i] + carry
		ct = cts[temp]
		b[i] = ct.temp
		carry = ct.carry

		i++
	}

	if carry > 0 {
		c := cap(this.b)
		if n < c {
			this.b = this.b[:n+1]
		} else {
			a := make([]byte, n+1, c*2)
			copy(a[:n], this.b)
			this.b = a
		}
		this.b[n] = carry
	}
}

func (this *privNumber) ReverseThenAdd() {

	//this.reverseThenAdd_1()
	this.reverseThenAdd_2()
}

func (this *privNumber) String() string {

	b := this.b
	buffer := new(bytes.Buffer)

	i := len(b)
	for i > 0 {
		i--
		buffer.WriteByte(digs[b[i]])
	}

	return buffer.String()
}

func (this *privNumber) CountDigits() (n int) {

	if this != nil {
		n = len(this.b)
	}

	return
}

func LychrelTest(n Number, count int) int {

	for i := 0; i < count; i++ {

		if n.IsPalindrome() {
			return i
		}

		n.ReverseThenAdd()
	}

	return count
}

func FindLychrelNumbers(max uint64, base int) []uint64 {

	var ns []uint64
	const count = 1000

	t := NewTable(base)
	for i := uint64(1); i < max; i++ {

		n := NewNumberTable(t, i)
		number := LychrelTest(n, count)
		if number == count {
			ns = append(ns, i)
		}
	}

	return ns
}
