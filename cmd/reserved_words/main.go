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

	temPar := strings.SplitAfter(gramParts[0], `/* SQL */`)
	keyWords := getReservedWords(directory)
	exists := getTypedTokens(bytes)

	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(temPar[0])
	if err != nil {
		panic(err)
	}

	writeTypedTokens(exists, keyWords, f)

	_, err = f.WriteString(temPar[1])
	if err != nil {
		panic(err)
	}

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

	printDiff(old, keyWords, directory)
}

func printDiff(old map[string]struct{}, keyWords map[string]struct{}, directory string) {
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

	lexertokens := getTokensFromLexer(directory)

	fmt.Println("\nNot in lexer.rl: ")
	for k := range keyWords {
		if _, ok := lexertokens[k]; !ok {
			fmt.Println(k)
		}
	}
}

func writeTypedTokens(exists map[string]struct{}, keyWords map[string]struct{}, f *os.File) {
	if len(exists) < len(keyWords) {
		_, err := f.WriteString("\n/* generated */\n%token<str> ")
		if err != nil {
			panic(err)
		}
	}

	for word := range keyWords {
		if _, ok := exists[word]; ok {
			continue
		}
		_, err := f.WriteString(word + " ")
		if err != nil {
			panic(err)
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

		if strings.Compare(word, "$end") == 0 || strings.Compare(word, "error") == 0 || strings.Compare(word, "$unk") == 0 || strings.Compare(word, "copy_generic_opt_list") == 0 {
			continue
		}

		keyWords[word] = struct{}{}
	}

	return keyWords
}

func getTypedTokens(b []byte) map[string]struct{} {
	operations := make(map[string]struct{})

	temp := strings.Split(string(b), "\n")

	for _, el := range temp {
		if !strings.Contains(el, "%token<str>") {
			continue
		}

		parts := strings.Split(strings.ReplaceAll(strings.Split(el, `/*`)[0], "	", " "), " ")

		for _, part := range parts {
			if part == "%token<str>" || part == `` {
				continue
			}
			fmt.Println(part)
			operations[part] = struct{}{}
		}
	}

	return operations
}

func getTokensFromLexer(directory string) map[string]struct{} {
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
