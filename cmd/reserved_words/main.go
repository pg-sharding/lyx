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
	old := make(map[string]struct{})
	old[strings.Split(strings.Trim(last[0], "	 "), " {$$=$1}")[0]] = struct{}{}
	for strings.Contains(last[num], "|") {
		old[strings.Split(strings.Split(last[num], "| ")[1], " {$$=$1}")[0]] = struct{}{}
		num++
	}

	for num < len(last) {
		_, err = f.WriteString(last[num])
		if err != nil {
			panic(err)
		}
		num++
	}

	fmt.Println("Removed: ")
	for k := range old {
		if _, ok := keyWords[k]; !ok {
			fmt.Println(k)
		}
	}
	fmt.Println("\nAdded: ")
	for k := range keyWords {
		if _, ok := old[k]; !ok {
			fmt.Println(k)
		}
	}
}

func getReservedWords(directory string) map[string]struct{} {
	b, err := ioutil.ReadFile(directory + "/lyx/gram.go")
	if err != nil {
		panic(err)
	}

	keyWords := make(map[string]struct{})

	temp := strings.Split(strings.Split(strings.Split(string(b), "var yyToknames =")[1], "{")[1], "}")[0]

	for _, el := range strings.Split(temp, "\n") {
		if strings.Contains(el, `//`) || len(el) < 3 {
			continue
		}

		word := strings.Split(el, `"`)[1]

		if strings.Compare(word, "$end") == 0 || strings.Compare(word, "error") == 0 || strings.Compare(word, "$unk") == 0 {
			continue
		}

		keyWords[word] = struct{}{}
	}

	return keyWords
}
