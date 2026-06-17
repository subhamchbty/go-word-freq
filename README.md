# go-word-freq

A command-line tool that counts word frequencies in a text file. Built as a learning exercise while going through Go fundamentals.

## What it does

Give it a text file and it prints the most frequent words in descending order. You can control how many results you want, skip common filler words, or get the output as JSON.

## Usage

```bash
go run main.go [flags] <file>
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-top` | `10` | Show only the top N most frequent words |
| `-stop` | `false` | Filter out common English stop words (the, a, is, in, etc.) |
| `-json` | `false` | Output results as JSON instead of a table |

**Examples**

```bash
# Top 10 words (default)
go run main.go sample.txt

# Top 5 words
go run main.go -top=5 sample.txt

# Skip common words like "the", "a", "and"
go run main.go -stop sample.txt

# Get JSON output
go run main.go -json sample.txt

# Combine flags
go run main.go -top=3 -stop -json sample.txt
```

**Sample output**

```
WORD              COUNT
-------------------------------
go                     12
func                    9
error                   7
```

```json
[{"word":"go","count":12},{"word":"func","count":9},{"word":"error","count":7}]
```

## What I learned building this

This was my first real Go project beyond hello world. Some things that clicked for me:

- **Maps and zero values** — you don't need to check if a key exists before incrementing `freq[word]++`. Go initialises missing int keys to `0` automatically, which feels weird coming from other languages (JS) but is actually really handy.
- **`strings.FieldsSeq`** — splits on any whitespace and skips empty tokens. Much better than `strings.Split(s, " ")` which leaves empty strings if there are double spaces.
- **`sort.Slice` with a closure** — I had to write a custom comparator to sort by count descending and break ties alphabetically. The closure captures the slice, which felt like magic at first.
- **Returning errors vs. handling them** — I got the feedback early that functions like `readFile` should return errors up to the caller instead of calling `log.Fatal` internally. Makes the function reusable and testable.
- **Struct tags for JSON** — adding `` `json:"word"` `` to struct fields was new to me. Without it, `encoding/json` just uses the field name as-is (capitalised), which isn't great for JSON consumers.

## What's still rough

- The stop words list is tiny. A real implementation would have hundreds.
- No stdin support yet — you have to pass a file path.
- Tests would be good to add, especially for `countWords` since the normalisation logic has a few edge cases.

## Running

Requires Go 1.23+ (uses `strings.FieldsSeq` which was added in 1.23).

```bash
go run main.go sample.txt
```
