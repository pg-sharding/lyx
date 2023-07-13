package main

import (
	"bufio"
	"lx/lx"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	lx.Run(scanner)
}
