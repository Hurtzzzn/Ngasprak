package main

import "fmt"

func palindrom(s string, kiri, kanan int) bool {
	if kiri >= kanan {
		return true
	}
	if s[kiri] != s[kanan] {
		return false
	}
	return palindrom(s, kiri+1, kanan-1)
}

func main() {
	var st string
	fmt.Scan(&st)
	if palindrom(st, 0, len(st)-1) {
		fmt.Println("YA")
	} else {
		fmt.Println("TIDAK")
	}
}
