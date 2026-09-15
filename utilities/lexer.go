package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct{ Kind, Value string }

func Tokenize(src string) ([]Token, error) {
	var tokens []Token
	i := 0
	for i < len(src) {
		c := src[i]
		switch {
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			i++
		case strings.ContainsRune("{}[],:", rune(c)):
			tokens = append(tokens, Token{"PUNCT", string(c)})
			i++
		case c == '"':
			j := i + 1

			for j < len(src) && src[j] != '"' {
				j++
			}
			decoded, err := strconv.Unquote(src[i : j+1])

			if err != nil {
				fmt.Println("Unexpected Token")
			}
			tokens = append(tokens, Token{"STRING", "\"" + decoded + "\""})
			i = j + 1
		case c == '-' || (c >= '0' && c <= '9'):
			j := i
			if c == '-' {
				j++
			}
			for j < len(src) && src[j] >= '0' && src[j] <= '9' {
				j++
			}

			if j < len(src) && src[j] == '.' {
				j++
				for j < len(src) && ((src[j] >= '0' && src[j] <= '9') || src[j] == 'e') {
					j++
				}
			}

			tokens = append(tokens, Token{"NUMBER", src[i:j]})
			i = j
		case strings.HasPrefix(src[i:], "true"):
			tokens = append(tokens, Token{"TRUE", "true"})
			i += 4
		case strings.HasPrefix(src[i:], "false"):
			tokens = append(tokens, Token{"FALSE", "false"})
			i += 5
		case strings.HasPrefix(src[i:], "null"):
			tokens = append(tokens, Token{"NULL", "null"})
			i += 4
		default:
			return nil, fmt.Errorf("unexpected character %q at position %d", c, i)
		}
	}
	tokens = append(tokens, Token{"EOF", ""})
	return tokens, nil
}
