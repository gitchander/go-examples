package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	testRSA()
}

func testRSA() {

	var (
		r = newRandNow()
		p = randPrime(r)
		q = randPrime(r)
	)

	// var (
	// 	p = 3557
	// 	q = 2579
	// )

	fmt.Println("p:", p)
	fmt.Println("q:", q)

	n := p * q

	fmt.Println("n:", n)

	phi := (p - 1) * (q - 1)
	fmt.Println("phi:", phi)

	e := pickPE(phi)
	fmt.Println("e:", e)

	d := modularInverse(e, phi)
	fmt.Println("d:", d)

	publicKey := PublicKey{E: e, N: n}
	fmt.Println("public key:", publicKey)

	privateKey := PrivateKey{D: d, N: n}
	fmt.Println("private key:", privateKey)

	//---------------------------------------------

	plaintext := 111111

	// Encrypt
	ciphertext := fastPowerMod(plaintext, publicKey.E, publicKey.N)
	fmt.Println("ciphertext:", ciphertext)

	// Decrypt
	plaintextDec := fastPowerMod(ciphertext, privateKey.D, privateKey.N)
	fmt.Println("plaintextDec:", plaintextDec)

	if plaintext != plaintextDec {
		log.Fatalf("%d != %d", plaintext, plaintextDec)
	}
}

func testPrime() {
	x := math.MaxInt32 - 18
	x = 9

	// fmt.Println(x)
	// fmt.Println(math.MaxInt64)
	// return

	ok := IsPrime(x)
	if ok {
		fmt.Printf("%d is prime\n", x)
	} else {
		fmt.Printf("%d is not prime\n", x)
	}

	fmt.Println(gcd(54, 24))
}

type PublicKey struct {
	N int // modulus
	E int // public exponent
}

type PrivateKey struct {
	N int // modulus
	D int // private exponent
}

func IsPrime(a int) bool {
	if a < 2 {
		return false
	}
	d := 2
	for d*d <= a {
		if (a % d) == 0 {
			return false
		}
		d++
	}
	return true
}

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randByte(r *rand.Rand) byte {
	return byte(r.Uint32())
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randIntBits(r *rand.Rand, n int) int {
	var d int
	for i := 1; i < n-1; i++ {
		if randBool(r) {
			d |= 1 << i
		}
	}
	d |= 1 << 0       // set first bit
	d |= 1 << (n - 1) // set last bit
	return d
}

func randPrime(r *rand.Rand) int {
	for {
		x := randIntBits(r, 14)

		//x := r.Int31() & 0xFF
		//x |= 1 // set 0-th bit
		//fmt.Println(x)

		//p := r.Int() & 0xFFFFFFFF

		p := int(x)

		if IsPrime(p) {
			return p
		}
	}
}

// (a ^ n) mod m
func fastPowerMod(a, n, m int) int {
	if m == 1 {
		return 0
	}
	b := 1
	a = a % m
	for n > 0 {
		if (n & 1) == 1 { // n is odd
			b = (b * a) % m
		}
		n >>= 1 // n = n / 2
		a = (a * a) % m
	}
	return b
}

// GCD - Greatest Common Denominator: largest number that can devide two numbers.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// power returns a^n
func power(a int, n int) int {
	b := 1
	for i := 0; i < n; i++ {
		b *= a
	}
	return b
}

// fastPower returns a^n
func fastPower(a int, n int) int {
	b := 1
	for n > 0 {
		if (n & 1) == 1 { // n is odd
			b *= a // b = b * a
		}
		n >>= 1 // n = n / 2
		a *= a  // a = a ^ 2
	}
	return b
}

var globalPrimes []int

func init() {
	//globalPrimes = calcPrimes(math.MaxInt32)
	globalPrimes = calcPrimes(100000)
}

func calcPrimes(max int) []int {
	var primes []int
	for a := 2; a < max; a++ {
		if IsPrime(a) {
			primes = append(primes, a)
		}
	}
	return primes
}

// public exponent
// pick public exponent
func pickPE(fi int) int {

	for _, prime := range globalPrimes {

		// This will simply pick the first, not a great idea,
		// but simple to understand algo.

		if ((prime % fi) != 0) && (gcd(prime, fi) == 1) {
			return prime
		}
	}

	// Value must be greater than 1, so this is sufficent to indicate failure.
	return 0
}

// Khan Academy:
// A number multiplied by its inverse = 1
// In modulo arithmetic, the modular inverse exists
// A % C == A**-1
// Only numbers coprime to C (share no prime factors with C)
// have a modular inverse (mod C)
// Can be optimized by using the Euclidean Algorithm.
func modularInverse_(a, c int) int {
	a %= c
	for i := 1; i < c; i++ {
		if ((a * i) % c) == 1 {
			return i
		}
	}
	return 0
}

func modularInverse(e, phi int) int {
	return invMod(e, phi)
}

// Расширенный алгоритм Евклида.
func gcdex(a, b int, x, y *int) int {
	if a == 0 {
		*x = 0
		*y = 1
		return b
	}
	var x1, y1 int
	d := gcdex(b%a, a, &x1, &y1)
	*x = y1 - (b/a)*x1
	*y = x1
	return d
}

// Обратное по модулю.
func invMod(a, m int) int {
	var x, y int
	gcdex(a, m, &x, &y)
	x = (x%m + m) % m
	return x
}
