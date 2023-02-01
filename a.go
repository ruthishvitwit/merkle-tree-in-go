package main

import (
	"fmt"
	a "merkle/merkle"
)

func addnode(i string) {
	t[0] = append(t[0], i)
	t = append(t[:1])
	build()
}

func deletenode(data string) {
	str := t[0]
	for i := 0; i < len(str); i++ {
		if str[i] == data {
			str = append(str[:i], str[i+1:]...)
			t = [][]string{}
			t[0] = append(t[0], str[:]...)
			build()
		}
	}
}

func find(data string) bool {
	str := t[0]
	for i := 0; i < len(str); i++ {
		if str[i] == data {
			return true
		}
	}
	return false
}

func build() {
	//n := height(len(t[0]))
	var hash []string
	for i := 0; i < len(t[0]); i++ {
		hash = append(hash, a.Sha(t[0][i]))
	}
	t = append(t, hash)
	k := 1
	// fmt.Println(n)
	for len(t[k]) > 1 {
		s := []string{}
		for i := 0; i < len(t[k]); i = i + 2 {
			var j string
			if i+1 < len(t[k]) {
				j = t[k][i]
			} else {
				j = t[k][i] + t[k][i+1]
			}
			// fmt.Println(len(t[k]))
			// fmt.Println(j)
			s = append(s, a.Sha(j))
			fmt.Println(s)
		}
		k++
		fmt.Println("sucess")
		t = append(t, s)
	}
	//fmt.Println(len(t[2][0]))
}

var t [][]string

func main() {
	t = [][]string{{"hi", "in"}}
	build()
	fmt.Println(find("hi"))

}
