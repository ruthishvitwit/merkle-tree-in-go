package main

import (
	"fmt"
	"math"
)

func main() {
	j := float64(4)
	var n float64
	for n = 0; n < 100; n++ {
		if j <= math.Pow(2, n) {
			k := int(n)
			fmt.Println(k)
			break
		}
	}

	//	t := [][]string{{"a", "b", "c", "d"}}
	// t[0] = append(t[0], "1")
	// new := 0
	// for k := 0; k < 3; k++ {
	// 	for i := 0; i < len(t[k]); i++ {
	// 		// j := t[k][i] + t[k][i+1]

	// 		// fmt.Println(j)
	// 		t[k+1] = append
	// 		// new++
	// 	}
	// 	// new = 0
	// }
	// s := []string{"a"}
	// t = append(t, s)
	// fmt.Println(len(t[0][4]))

	// for i := 0; i < 2; i++ {
	// 	s := []string{}

	// 	for j := 0; j < 4; j++ {
	// 		s = append(s, "a")
	// 	}

	// 	t = append(t, s)
	// }
	// fmt.Println(t)
	// j := float64(64)
	// var n float64
	// var k int
	// for n = 0; n < 100; n++ {
	// 	if j <= math.Pow(2, n) {
	// 		k = int(n)
	// 		break
	// 	}

	// }
	// fmt.Println(k)
}
