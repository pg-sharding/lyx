package main

import (
	"bufio"
	"lyx/lyx"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	lyx.Run(scanner)
}
