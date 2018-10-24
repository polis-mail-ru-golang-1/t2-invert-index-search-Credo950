package invertIndex

import "fmt"

type Dictionary = map[string]map[string]int

func FillDict(words []string, dict Dictionary, fileName string) {
	for _, word := range words {
		_, ok := dict[word]
		if ok {
			dict[word][fileName]++
		} else {
			dict[word] = map[string]int{}
			dict[word][fileName]++
		}
	}
	return
}

func CountWords(dict Dictionary, phrase []string) map[string]int {
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
	return counter
}

func DictPrintToString(dict Dictionary) []string {
	var out []string
	out = append(out, "-------------------------------------------------------------\n")
	for word, item := range dict {
		out = append(out, fmt.Sprintf("%10s %+v\n", word, item))
	}
	out = append(out, "-------------------------------------------------------------")
	return out
}
