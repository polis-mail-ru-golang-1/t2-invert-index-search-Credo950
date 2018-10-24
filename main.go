package main

import (
	"fmt"
	"log"
	"os"

	format "github.com/polis-mail-ru-golang-1/t2-invert-index-search-Credo950/formating"
	inv "github.com/polis-mail-ru-golang-1/t2-invert-index-search-Credo950/invertIndex"
)

func main() {
	dict := make(inv.Dictionary)
	args := os.Args[1:]
	for _, fileName := range args {
		words, err := format.ReadWordsFromFile(fileName)
		if err != nil {
			log.Fatalf("%s", err)
		}
		inv.FillDict(words, dict, fileName)
	}

	fmt.Println(inv.DictPrintToString(dict))

	phrase := format.ReadPhrase()

	counter := inv.CountWords(dict, phrase)

	PrintResult(counter)
}

func PrintResult(counter map[string]int) {
	if len(counter) == 0 {
		return
	}
	var temp string
	var max int
	for i, item := range counter {
		if item > max {
			temp = i
			max = item
		}
	}
	fmt.Printf("- %s; совпадений - %d\n", temp, max)
	delete(counter, temp)
	PrintResult(counter)
}
