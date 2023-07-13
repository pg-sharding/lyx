package main

import (
	"bufio"
	lx "lx/lyx"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	lx.Run(scanner)
}
