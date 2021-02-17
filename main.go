package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b []byte
	for _, s := range os.Args[1:] {
		for _, w := range strings.Split(s, " ") {
			v, err := strconv.ParseUint(w, 10, 8)
			if err != nil {
				panic(err)
			}
			b = append(b, uint8(v))
		}
	}
	fmt.Printf("%s", b)
}
