package formating

import (
	"bufio"
	"os"
	"strings"
)

func ReadWordsFromFile(fileName string) ([]string, error) {
	lines, err := ReadLines(fileName)
	if err != nil {
		return nil, err
	}
	words := LinesToWords(lines)
	return words, nil
}

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

func ReadPhrase() []string {
	var phrase []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phrase = append(phrase, scanner.Text())
	return LinesToWords(phrase)
}
