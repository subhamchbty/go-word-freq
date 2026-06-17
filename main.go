package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

var stopWords = map[string]bool{
	"the": true,
	"i":   true,
	"to":  true,
	"a":   true,
	"for": true,
	"is":  true,
	"in":  true,
	"and": true,
	"at":  true,
}

func main() {
	top := flag.Int("top", 10, "Specify top-N integer. Default is 10.")
	stop := flag.Bool("stop", false, "Specify if common english words are stopped from counting. Default is false.")
	jsonOut := flag.Bool("json", false, "Specify if output should be in JSON format. Default is false.")

	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		log.Fatal("Error: File path is required")
		os.Exit(1)
	}

	data, err := readFile(filePath)
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	freq := countWords(data, *stop)
	printResults(mapToSlice(freq, *top), *jsonOut)

}

func countWords(text string, stop bool) map[string]int {
	freq := map[string]int{}

	lowered := strings.ToLower(text)
	words := strings.FieldsSeq(lowered)

	for value := range words {
		word := strings.Trim(value, ".,!?;:\"'()-")
		if word == "" {
			continue
		}

		if stop && stopWords[word] == true {
			continue
		}

		freq[word] = freq[word] + 1
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

	return slice[:min(top, len(slice))]
}

func printResults(counts []WordCount, jsonOut bool) {
	if !jsonOut {
		fmt.Printf("%-15s %15s\n", "WORD", "COUNT")
		fmt.Printf("-------------------------------\n")
		for _, freq := range counts {
			fmt.Printf("%-15s %15d\n", freq.Word, freq.Count)
		}
	} else {
		jsonData, err := json.Marshal(counts)
		// jsonData, err := json.MarshalIndent(counts, "", "    ")
		if err != nil {
			log.Fatal("Error: cannot format to json")
		}

		fmt.Println(string(jsonData))
	}
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
