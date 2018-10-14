package main

import funct "./functional"
import "log"

func main() {
	dict := make(map[string][]string)

	args, err := funct.FillDict(dict)
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	funct.PrintDict(dict)

	phrase := funct.ReadPhrase()

	counter := make(map[string]int)
	for _, fileName := range args {
		counter[fileName] = 0
	}

	for i, fileSlice := range dict {
		for _, word := range phrase {
			if i == word {
				for _, fileName := range fileSlice {
					counter[fileName]++
				}
			}
		}
	}

	funct.PrintResult(counter)
}
