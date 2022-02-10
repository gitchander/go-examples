package main

import (
	"fmt"
	"math"
)

// https://golang.org/pkg/fmt/

func main() {
	fmt.Printf("%d\n", 9876543210) // 9876543210
	fmt.Printf("five=% d\n", 5)    // five= 5

	fmt.Printf("%f\n", math.Pi)       // 3.141593
	fmt.Printf("%.2f\n", math.Pi)     // 3.14
	fmt.Printf("|%8.2f|\n", math.Pi)  // |    3.14|
	fmt.Printf("|%-8.2f|\n", math.Pi) // |3.14    |
	fmt.Printf("%g\n", math.Pi)       // 3.141592653589793
	fmt.Printf("%e\n", 3.23546e-7)    // 3.235460e-07

	// %% - literal percent sign (no operand)
	fmt.Printf("%d %%\n", 100) // 100 %

	// right-justify
	fmt.Printf("|%4d|%4d|%4d|%4d|\n", 1, 21, 321, 4321) // |   1|  21| 321|4321|

	// left-justify
	fmt.Printf("|%-4d|%-4d|%-4d|%-4d|\n", 1, 21, 321, 4321) // |1   |21  |321 |4321|

	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	// byte hex
	fmt.Printf("%#x\n", 0x07)    // 0x7
	fmt.Printf("%#X\n", 0x07)    // 0X7
	fmt.Printf("%#X\n", 0xA7)    // 0XA7
	fmt.Printf("%#02x\n", 0x07)  // 0x07
	fmt.Printf("0x%02x\n", 0x07) // 0x07
	fmt.Printf("0x%02x\n", 0xA7) // 0xa7
	fmt.Printf("0x%02X\n", 0xA7) // 0xA7

	// bytes hex
	fmt.Printf("%x\n", 123456789) // 75bcd15
	fmt.Printf("%X\n", 123456789) // 75BCD15
	fmt.Printf("%x\n", data)      // 0123456789abcdef
	fmt.Printf("%X\n", data)      // 0123456789ABCDEF
	fmt.Printf("% x\n", data)     // 01 23 45 67 89 ab cd ef
	fmt.Printf("% X\n", data)     // 01 23 45 67 89 AB CD EF
	fmt.Printf("%x\n", -0.25)     // -0x1p-02

	fmt.Printf("%o, %o\n", 8, 64) // 10, 100

	fmt.Printf("%b\n", 123456) // 11110001001000000

	// %t - boolean: true or false
	fmt.Printf("%t\n", true)  // true
	fmt.Printf("%t\n", false) // false

	fmt.Printf("%s\n", "this is a text") // this is a text

	// %q - quoted string "abc" or rune 'c'
	fmt.Printf("%q\n", "Это буква") // "Это буква"
	fmt.Printf("%q\n", 'Щ')         // 'Щ'

	// Unicode
	// %c - rune (Unicode code point)
	fmt.Printf("%c\n", 0x0410)  // А
	fmt.Printf("%c\n", 0x042F)  // Я
	fmt.Printf("U+%04X\n", 'А') // U+0410
	fmt.Printf("%U\n", 'А')     // U+0410
	fmt.Printf("%U\n", 'Я')     // U+042F
	fmt.Printf("%#U\n", 'Ї')    // U+0407 'Ї'

	// %T
	fmt.Printf("%T\n", true)          // bool
	fmt.Printf("%T\n", 0)             // int
	fmt.Printf("%T\n", uint8(0))      // uint8
	fmt.Printf("%T\n", int64(0))      // int64
	fmt.Printf("%T\n", 0.0)           // float64
	fmt.Printf("%T\n", complex(1, 2)) // complex128

	// %v
	u := user{name: "Jack", age: 1000}
	fmt.Printf("%v\n", u)  // {Jack 1000}
	fmt.Printf("%+v\n", u) // {name:Jack age:1000}
	fmt.Printf("%#v\n", u) // main.user{name:"Jack", age:1000}

	// Explicit argument indexes:
	fmt.Printf("%[3]d %[2]d %[1]d\n", 1, 2, 3) // 3 2 1
	fmt.Printf("%[2]d %[2]d %[2]d\n", 1, 2, 3) // 2 2 2

	fmt.Printf("%[1]*[2]d%[1]*[3]d%[1]*[4]d%[1]*[5]d\n", 5, 1, 2, 3, 4) //    1    2    3    4
	fmt.Printf("%5[1]d%5[2]d%5[3]d%5[4]d\n", 1, 2, 3, 4)                //    1    2    3    4
	fmt.Printf("%5d%5d%5d%5d\n", 1, 2, 3, 4)                            //    1    2    3    4

	fmt.Printf("%[1]*d\n", 5, 123) //  123
	fmt.Printf("%5d\n", 123)       //  123

	fmt.Printf("%[1]*.[2]*[3]f\n", 8, 4, 3.1415) //  3.1415
	fmt.Printf("%[1]*.[2]*f\n", 8, 4, 3.1415)    //  3.1415
	fmt.Printf("%8.4f\n", 3.1415)                //  3.1415
}

type user struct {
	name string
	age  int
}
