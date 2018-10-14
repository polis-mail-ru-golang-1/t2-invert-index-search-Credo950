package functional

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LinesToWords(lines []string) []string {
	var result []string
	for _, item := range lines {
		arr := strings.Split(item, " ")
		for _, item := range arr {
			result = append(result, item)
		}
	}
	return result
}

func FillDict(dict map[string][]string) ([]string, error) {
	args := os.Args[1:]
	for _, fileName := range args {
		lines, err := ReadLines(fileName)
		if err != nil {
			return nil, err
		}
		words := LinesToWords(lines)
		for _, word := range words {
			if !Include(dict[word], fileName) {
				dict[word] = append(dict[word], fileName)
			}
		}
	}
	return args, nil
}

func PrintDict(dict map[string][]string) {
	fmt.Println("-------------------------------------------------------------")
	for i, item := range dict {
		fmt.Printf("%10s %s\n", i, item)
	}
	fmt.Println("-------------------------------------------------------------")
}

func ReadPhrase() []string {
	var phrase []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phrase = append(phrase, scanner.Text())
	return LinesToWords(phrase)
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
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