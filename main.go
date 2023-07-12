package main

import (
	"fmt"
	"lx/lx"
	"os"
)

func main() {
	tt := lx.NewTokenizer(os.Stdin)
	for {
		v := tt.Tok()
		if v == 0 {
			fmt.Println("end")
		} else {
			fmt.Printf("token type %d\n", v)
		}
	}
}
