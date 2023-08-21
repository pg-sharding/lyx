package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile(directory + "/lyx/gram.y")
	if err != nil {
		panic(err)
	}

	gramParts := strings.SplitAfter(string(bytes), "reserved_keyword:\n")

	f, err := os.OpenFile(directory+"/lyx/gram.y", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(gramParts[0])
	if err != nil {
		panic(err)
	}

	keyWords := getReservedWords(directory)
	num := 0
	for word := range keyWords {
		if num == 0 {
			_, err = f.WriteString("	" + word + " {$$=$1}\n")
			if err != nil {
				panic(err)
			}
		} else {
			_, err = f.WriteString("	" + "| " + word + " {$$=$1}\n")
			if err != nil {
				panic(err)
			}
		}
		num++
	}

	last := strings.SplitAfter(gramParts[1], "\n")
	num = 1
	for strings.Contains(last[num], "|") {
		num++
	}

	for num < len(last) {
		_, err = f.WriteString(last[num])
		if err != nil {
			panic(err)
		}
		num++
	}
}

func getReservedWords(directory string) map[string]struct{} {
	b, err := ioutil.ReadFile(directory + "/lyx/lexer.rl")
	if err != nil {
		panic(err)
	}

	keyWords := make(map[string]struct{})

	for i, el := range strings.Split(string(b), "\n") {
		if i == 0 || strings.Contains(el, "#") {
			continue
		}

		tem := strings.Split(el, "tok = ")

		if len(tem) == 1 {
			continue
		} else if len(tem) > 2 {
			panic(fmt.Errorf("more than one token in string"))
		}

		tok := strings.Split(tem[1], ";")
		str := tok[0]
		if strings.Contains(tok[0], "(") {
			str = strings.ReplaceAll(strings.Split(tok[0], "(")[1], ")", "")
		}
		keyWords[str] = struct{}{}
	}

	return keyWords
}
