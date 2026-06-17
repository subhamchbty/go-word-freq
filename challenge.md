# Exercise: Word Frequency Counter in Go

**Difficulty:** Beginner → Intermediate  
**Estimated time:** 60–90 minutes  
**Concepts covered:** maps, slices, structs, sorting, strings, file I/O, closures, variadic functions

---

## Background

A word frequency counter reads a body of text and reports how often each unique word appears. It's a classic exercise that touches nearly every Go fundamental: string manipulation, maps, slices of structs, custom sort logic, and eventually file I/O and command-line flags.

You will build this incrementally across five parts. **Do not skip ahead** — each part introduces a concept that the next part depends on.

---

## Setup

Create a new directory and initialise a module:

```
mkdir word-freq && cd word-freq
go mod init word-freq
touch main.go
```

All your work goes in `main.go` (for now). Later parts will ask you to split things into files.

---

## Part 1 — Count words in a hard-coded string

### Goal
Write a function that accepts a string and returns a `map[string]int` where each key is a word and each value is how many times it appears.

### Requirements

1. Define a function with the signature:
   ```go
   func countWords(text string) map[string]int
   ```
2. Split the input string into individual words. Use the `strings` package — look up `strings.Fields`.
3. For each word, increment its count in the map. Remember to handle the zero-value behaviour of maps in Go (you should **not** need to check if a key exists before incrementing).
4. In `main`, call `countWords` with the following string and print the resulting map:
   ```
   "the quick brown fox jumps over the lazy dog the fox"
   ```

### Expected output shape
The output does not need to be sorted yet. Any order is fine. You should see `"the"` with a count of `3` and `"fox"` with a count of `2`.

### Questions to think about
- What is the zero value of `int` in Go, and why does it make map incrementing safe without an existence check?
- What does `strings.Fields` do differently from `strings.Split(s, " ")`?

---

## Part 2 — Normalise the input

Real text has punctuation, mixed casing, and extra whitespace. A search for `"The"` and `"the"` should hit the same bucket.

### Requirements

1. Before splitting, convert the entire input to lowercase. Use `strings.ToLower`.
2. After splitting into words, strip leading and trailing punctuation from each word. Use `strings.Trim` with the character set `".,!?;:\"'()-"`.
3. Skip any word that becomes an empty string after trimming.
4. Update `countWords` to apply these steps internally so callers don't need to think about it.

### Test your changes with:
```
"To be, or not to be — that is the question!"
```

You should see `"to"` with a count of `2` and `"be"` with a count of `2`.

### Questions to think about
- Why is it better to handle normalisation inside `countWords` rather than at the call site?
- What happens if you pass an empty string to `countWords`? What does the function return?

---

## Part 3 — Sort and display results

A raw map printout is unordered and hard to read. Your task is to sort the results by frequency (descending) and print them in a clean table.

### Step 3a — Define a struct

Define a named struct to hold a single word and its count:

```go
type WordCount struct {
    Word  string
    Count int
}
```

### Step 3b — Convert the map to a slice

Write a function:

```go
func mapToSlice(freq map[string]int) []WordCount
```

It should iterate over the map and build a `[]WordCount` from it. Use a `range` loop over the map.

### Step 3c — Sort the slice

Use the `sort` package. Sort by `Count` descending. When two words have the same count, sort them alphabetically (ascending) as a tiebreaker — this makes output deterministic.

Look up `sort.Slice`. It takes a slice and a **less** function. You need to write that less function yourself.

### Step 3d — Print a formatted table

Write a function:

```go
func printResults(counts []WordCount)
```

It should print a header and one row per word using `fmt.Printf` with alignment. The word column should be left-aligned and the count column right-aligned. Example format (widths are up to you):

```
WORD              COUNT
────────────────────────
the                   3
fox                   2
brown                 1
...
```

### Questions to think about
- Why can't you sort a `map[string]int` directly in Go?
- In `sort.Slice`'s less function, what should you return when counts are equal and you want alphabetical tiebreaking?
- What does `%-20s` vs `%20s` do in a format string?

---

## Part 4 — Read from a file

Hard-coded strings are fine for testing, but real tools read from files.

### Requirements

1. Create a sample text file called `sample.txt` in your project directory with at least 10 sentences of your choice (copy a paragraph from any article or book).

2. Write a function:
   ```go
   func readFile(path string) (string, error)
   ```
   Use `os.ReadFile` to read the file and return its contents as a string. Return any error to the caller — do not `log.Fatal` inside this function.

3. In `main`, call `readFile`. If it returns an error, print it and exit with a non-zero status code using `os.Exit(1)`.

4. Pass the file contents into your existing `countWords` pipeline.

### Questions to think about
- Why is it better for `readFile` to return an error instead of handling it internally?
- What is the difference between `os.ReadFile` and `os.Open` + manual reading?
- What does `os.Exit(1)` signal to a shell or CI system?

---

## Part 5 — Add a top-N flag

The user should be able to ask for only the top N most frequent words.

### Requirements

1. Use the `flag` package to add an integer flag called `top` with a default of `10` and a usage string explaining what it does.

2. Parse the flags at the start of `main` with `flag.Parse()`.

3. After sorting, slice your results to at most `top` entries before printing. Handle the edge case where the total number of unique words is less than `top`.

4. Also accept the filename as a positional argument using `flag.Args()`. If no filename is provided, print a usage message and exit.

### Run examples

```bash
# Show top 5 words
go run main.go -top=5 sample.txt

# Default: top 10
go run main.go sample.txt
```

### Questions to think about
- What is the difference between a flag (`-top=5`) and a positional argument (`sample.txt`) in Go's `flag` package?
- How do you safely slice a slice when you don't know if it has enough elements? What does `s[:min(n, len(s))]` protect against?

---

## Stretch Goals

If you finish early, try one or more of these:

**A. Exclude stop words**  
Define a set of common English stop words (`"the"`, `"a"`, `"is"`, `"in"`, etc.) as a `map[string]bool` and skip them during counting. Accept an optional `-stop` flag that enables this filter.

**B. Count from stdin**  
If no filename argument is given, read from `os.Stdin` instead of printing an error. This lets you pipe text in: `cat sample.txt | go run main.go`.

**C. Output as JSON**  
Add a `-json` flag. When set, output the results as a JSON array of objects using `encoding/json` instead of the table format.

**D. Multiple files**  
Accept multiple filenames as positional arguments and aggregate word counts across all of them.

**E. Percentage column**  
Add a third column to your table showing each word's count as a percentage of total word occurrences.

---

## Checklist before you consider it done

- [x] `countWords` correctly normalises case and strips punctuation
- [x] Empty strings are never inserted as map keys
- [x] Results are sorted: descending count, then alphabetically on ties
- [x] `readFile` returns an error; `main` handles it
- [x] `-top` flag works and never panics on short word lists
- [x] `go vet ./...` reports no issues
- [x] `gofmt -l .` reports no unformatted files
- [x] Optional: exclude stop words
- [ ] Optional: count from stdin
- [x] Optional: output as JSON
- [ ] Optional: multiple file support
- [ ] Optional: percentage column or JSON field

---

## Packages you are allowed to use

`fmt`, `os`, `strings`, `sort`, `flag`, `encoding/json` (stretch only)  

No third-party libraries. The standard library is more than enough.
