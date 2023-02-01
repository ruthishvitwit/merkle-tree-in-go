package merkle1

import (
	"fmt"
	"strconv"
)

// buffers
var h []string = []string{
	"6a09e667", "bb67ae85", "3c6ef372", "a54ff53a",
	"510e527f", "9b05688c", "1f83d9ab", "5be0cd19",
}

// these are the constants
var k []string = []string{
	"428a2f98", "71374491", "b5c0fbcf", "e9b5dba5",
	"3956c25b", "59f111f1", "923f82a4", "ab1c5ed5",
	"d807aa98", "12835b01", "243185be", "550c7dc3",
	"72be5d74", "80deb1fe", "9bdc06a7", "c19bf174",
	"e49b69c1", "efbe4786", "0fc19dc6", "240ca1cc",
	"2de92c6f", "4a7484aa", "5cb0a9dc", "76f988da",
	"983e5152", "a831c66d", "b00327c8", "bf597fc7",
	"c6e00bf3", "d5a79147", "06ca6351", "14292967",
	"27b70a85", "2e1b2138", "4d2c6dfc", "53380d13",
	"650a7354", "766a0abb", "81c2c92e", "92722c85",
	"a2bfe8a1", "a81a664b", "c24b8b70", "c76c51a3",
	"d192e819", "d6990624", "f40e3585", "106aa070",
	"19a4c116", "1e376c08", "2748774c", "34b0bcb5",
	"391c0cb3", "4ed8aa4a", "5b9cca4f", "682e6ff3",
	"748f82ee", "78a5636f", "84c87814", "8cc70208",
	"90befffa", "a4506ceb", "bef9a3f7", "c67178f2",
}

// convert const to binary
func conv_ktb() [64]string {
	var g [64]string
	for i := 0; i < 64; i++ {
		n := bin_32bit(dec_bin_hex(k[i]))
		g[i] = n
	}
	return g
}

// conv string to 8digit binary
func conv_ato8b(b string) string {
	a := [8]string{"0", "00", "000", "0000", "00000", "000000", "0000000", "00000000"}
	if len(b) < 8 {
		b = a[7-len(b)] + b

	}
	return b
}

// filling of binary to 448
func padding(s string) string {
	s = s + "1"
	for len(s) < 448 {
		s = s + "0"
	}
	return s
}

// adding prefix zeros to binary
func sizeadd(d int) string {
	var b string
	b = fmt.Sprintf("%b", d)
	for len(b) < 64 {
		b = "0" + b
	}
	return b
}

// dividing binary string into 32 spaced seperated strings
func divide(s string) [64]string {
	var a [64]string
	for i := 0; i < 16; i++ {
		a[i] = s[(i * 32):(i*32 + 32)]

	}
	for i := 16; i < 64; i++ {
		a[i] = "00000000000000000000000000000000"
	}
	return a
}

// xor on binary-strings
func xor(s string, a string) string {
	var b string
	var i int
	for i = 0; i < 32; i++ {
		if s[i] == a[i] {
			b = b + "0"
		} else {
			b = b + "1"
		}
	}
	//fmt.Println(b)
	return b
}

// rightrotate on binary-string
func rr(a string, i int) string {
	b := a[32-i:32] + a[:32-i]
	//fmt.Println(b)
	return b
}

// rightshift on binary-string
func rs(a string, i int) string {
	var b string
	for c := 0; c < i; c++ {
		b = "0" + b
	}
	b = b + a[:32-i]
	return b
}

// addition of binary-strings
func plus(s string, a string) string {
	var b string
	k := len(a) - 1
	j := "0"
	for i := len(s) - 1; i > -1; i-- {
		if k < 0 || i < 0 {
			break
		}
		if s[i] == '1' && a[k] == '1' {
			if j == "1" {
				b = "1" + b
				j = "1"
			} else {
				b = "0" + b
				j = "1"
			}
		} else if s[i] == '0' && a[k] == '0' {
			if j == "1" {
				b = "1" + b
				j = "0"
			} else {
				b = "0" + b
				j = "0"
			}
		} else {
			if j == "1" {
				b = "0" + b
				j = "1"
			} else if j == "0" {
				b = "1" + b
				j = "0"
			}
		}
		if (i == 0 || k == 0) && j == "1" {
			b = "1" + b
			j = "1"

		}
		k--
	}
	c := b[len(b)-32 : len(b)]
	return c
}

// decimal number convert binary with required number of bits length
func f(d int64, g int) string {
	b := strconv.FormatInt(d, 2)
	z := g - len(b)
	for i := 0; i < z; i++ {
		b = "0" + b
	}
	return b
}

// convert binary into 32 bit
func bin_32bit(d int64) string {
	return f(d, 32)
}

// convert the hexadecimal string into decimal format
func dec_bin_hex(h string) int64 {
	s, _ := strconv.ParseInt(h, 16, 64)
	return s
}

// modifying the string into array and applying changes to elements
func modify(a [64]string) [64]string {
	var s0, s1 string
	for i := 16; i < 64; i++ {
		s0 = xor(xor(rr(a[i-15], 7), rr(a[i-15], 18)), rs(a[i-15], 3))
		s1 = xor(xor(rr(a[i-2], 17), rr(a[i-2], 19)), rs(a[i-2], 10))
		a[i] = plus(plus(plus(a[i-16], s0), a[i-7]), s1)
		s0 = ""
		s1 = ""
	}

	return a
}

// and function for binary-string
func and(s string, a string) string {
	var b string
	var i int
	for i = 0; i < 32; i++ {
		if s[i] == '1' && a[i] == '1' {
			b = b + "1"
		} else {
			b = b + "0"
		}
	}
	//fmt.Println(b)
	return b
}

// not function for binary-string
func not(s string) string {
	var b string
	var i int
	for i = 0; i < 32; i++ {
		if s[i] == '1' {
			b = b + "0"
		} else {
			b = b + "1"
		}
	}
	//fmt.Println(b)
	return b
}

// div of elements into 8-bit binary
func div(b string) [4]string {
	var a [4]string
	for j := 0; j < 4; j++ {
		a[j] = b[(j * 8):(j*8 + 8)]
	}

	return a
}

// sub func to conv string to hexa
func conv_stoh(a [32]string) string {
	var b uint64
	var d string = ""
	for i := 0; i < 32; i++ {
		b, _ = strconv.ParseUint(a[i], 2, 64)
		c := fmt.Sprintf("%x", b)
		//fmt.Println(b, c)
		d = d + c
	}

	return d
}

// final conversion of binary-string to hexa
func finalconv(b [8]string) string {
	var a [32]string
	var d [4]string
	var j int = 0
	for i := 0; i <= 7; i++ {
		d = div(b[i])
		for k := 0; k < 4; k++ {
			a[j] = d[k]
			//fmt.Println(a[j])
			j++
		}
	}
	e := conv_stoh(a)
	return e
}

// compression of binary
func comprs(a [8]string, ar [64]string, c [64]string) string {
	b := a
	for i := 0; i < 64; i++ {
		s1 := xor(xor(rr(a[4], 6), rr(a[4], 11)), rr(a[4], 25))
		ch := xor(and(a[4], a[5]), and(not(a[4]), a[6]))
		temp1 := plus(plus(plus(plus(a[7], s1), ch), ar[i]), c[i])
		s0 := xor(xor(rr(a[0], 2), rr(a[0], 13)), rr(a[0], 22))
		m := xor(xor(and(a[0], a[1]), and(a[0], a[2])), and(a[1], a[2]))
		temp2 := plus(s0, m)
		a[7] = a[6]
		a[6] = a[5]
		a[5] = a[4]
		a[4] = plus(a[3], temp1)
		a[3] = a[2]
		a[2] = a[1]
		a[1] = a[0]
		a[0] = plus(temp1, temp2)
		//fmt.Println(s1, ch, temp1, s0, m, temp2, a[0])
	}
	for i := 0; i < 8; i++ {
		//fmt.Println(a[i])
		b[i] = plus(b[i], a[i])
		//fmt.Println(b[i])
	}
	//d := "hi"
	d := finalconv(b)
	return d

}
func Sha256(i string) string {

	var a1 [8]string
	for k := 0; k < 8; k++ {
		a1[k] = bin_32bit(dec_bin_hex(h[k]))
		//fmt.Println(a1[k])
	}
	k1 := conv_ktb()
	//i = i << 2
	var s string
	var b string
	var a [64]string
	for _, c := range i {
		b = fmt.Sprintf("%s%b", b, c)
		s = s + conv_ato8b(b)
		b = ""

	}
	d := len(s)
	//fmt.Println(s)
	b = sizeadd(d)
	s = padding(s)
	s = s + b
	a = divide(s)
	a = modify(a)
	e := comprs(a1, k1, a)

	return e

}

func Sha(j string) string {
	//var i string
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//i := scanner.Text()
	e := Sha256(j)
	//fmt.Println(len(b))
	//fmt.Println(e)
	return e
}
