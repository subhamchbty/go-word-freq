package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	top := flag.Int("top", 10, "Specify top-N integer. Default is 10.")

	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		log.Fatal("Error: File path is required")
		os.Exit(1)
	}

	data, err := readFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	freq := countWords(data)
	printResults(mapToSlice(freq, *top))

}

func countWords(text string) map[string]int {
	freq := map[string]int{}

	lowered := strings.ToLower(text)
	words := strings.FieldsSeq(lowered)

	for value := range words {
		word := strings.Trim(value, ".,!?;:\"'()-")
		if word == "" {
			continue
		}

		count := freq[word]

		freq[word] = count + 1
	}

	return freq
}

func mapToSlice(freq map[string]int, top int) []WordCount {
	slice := []WordCount{}

	for key, value := range freq {
		newCount := []WordCount{
			{Word: key, Count: value},
		}
		slice = append(slice, newCount...)
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].Count == slice[j].Count {
			return slice[i].Word < slice[j].Word
		}
		return slice[i].Count > slice[j].Count
	})

	return slice[:top]
}

func printResults(counts []WordCount) {
	fmt.Printf("%-15s %15s\n", "WORD", "COUNT")
	fmt.Printf("-------------------------------\n")
	for _, freq := range counts {
		fmt.Printf("%-15s %15d\n", freq.Word, freq.Count)
	}
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
