package main

import funct "./functional"
import "log"

func main() {
	dict := make(map[string]map[string]int)

	_, err := funct.FillDict(dict)
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	funct.PrintDict(dict)

	phrase := funct.ReadPhrase()

	counter := make(map[string]int)

	for i, item := range dict {
		for fileName, count := range item {
			for _, word := range phrase {
				if i == word {
					counter[fileName] += count
				}
			}
		}
	}

	funct.PrintResult(counter)
}
