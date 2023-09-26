package main

import (
	"fmt"
)

func main2() {
	m := map[string]string{
		"Яндекс":"+74957397000",
		"Музей Яндекса":"+74991101133",
	}
	
	fmt.Println(SwapKeysAndValues(m))
}

func SwapKeysAndValues(m map[string]string) map[string]string {
	n := make(map[string]string)
	for key, val := range m {
		n[val] = key
	}
	return n
}