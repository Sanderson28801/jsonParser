package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Token struct{ Kind, Value string }

func tokenize(src string) ([]Token, error) {
	var tokens []Token
	i := 0
	for i < len(src) {
		c := src[i]
		switch {
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			i++
		case strings.ContainsRune("{}[],:", rune(c)):
			tokens = append(tokens, Token{"PUNCT", string(c)}); i++
		case c == '"':
			// TODO: handle \" \\ \n \t \uXXXX escapes.
			j := i + 1
			for j < len(src) && src[j] != '"' { j++ }
			tokens = append(tokens, Token{"STRING", src[i+1 : j]}); i = j + 1
		case c == '-' || (c >= '0' && c <= '9'):
			j := i; if c == '-' { j++ }
			for j < len(src) && src[j] >= '0' && src[j] <= '9' { j++ }
			if j < len(src) && src[j] == '.' {
				j++
				for j < len(src) && src[j] >= '0' && src[j] <= '9' { j++ }
			}
			// TODO: handle scientific notation.
			tokens = append(tokens, Token{"NUMBER", src[i:j]}); i = j
		case strings.HasPrefix(src[i:], "true"):  tokens = append(tokens, Token{"TRUE",  "true"});  i += 4
		case strings.HasPrefix(src[i:], "false"): tokens = append(tokens, Token{"FALSE", "false"}); i += 5
		case strings.HasPrefix(src[i:], "null"):  tokens = append(tokens, Token{"NULL",  "null"});  i += 4
		default:
			return nil, fmt.Errorf("unexpected character %q at position %d", c, i)
		}
	}
	tokens = append(tokens, Token{"EOF", ""})
	return tokens, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if line == "" { continue }
		toks, err := tokenize(line)
		if err != nil { fmt.Println("ERR", err); continue }
		for _, t := range toks { fmt.Println(t.Kind, t.Value) }
		fmt.Println()
	}
}
