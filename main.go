package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	freq := countWords("the quick brown fox jumps over the lazy dog the fox")

	printResults(mapToSlice(freq))
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

func mapToSlice(freq map[string]int) []WordCount {
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

	return slice
}

func printResults(counts []WordCount) {
	fmt.Printf("%-15s %15s\n", "WORD", "COUNT")
	fmt.Printf("-------------------------------\n")
	for _, freq := range counts {
		fmt.Printf("%-15s %15d\n", freq.Word, freq.Count)
	}
}
