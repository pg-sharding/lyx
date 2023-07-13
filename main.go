package main

import (
	"bufio"
	"os"

	"github.com/pg-sharding/lyx/lyx"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	lyx.Run(scanner)
}
