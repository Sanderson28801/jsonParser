// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"regexp"
// )

// // TODO (parse-array): implement per the lesson description.

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
// 	for sc.Scan() {
// 		line := sc.Text()
// 		if line == "" {
// 			continue
// 		}
// 		if line[0] == '[' {
// 			fmt.Println(ParseArray(line))
// 		}
// 	}
// }

// func ParseArray(line string) []string {

// 	h := []string{"hi"}
// 	return h
// }

// func ParseValue(line string) string {
// 	switch line {
// 	case "true":
// 		return "True"
// 	case "false":
// 		return "False"
// 	case "null":
// 		return "None"
// 	}
// 	// if b, err := strconv.ParseInt(line, 10, 64); err == nil {
// 	// 	if b == -0 {
// 	// 		return "0"
// 	// 	}
// 	// 	return line
// 	// } else if f, err := strconv.ParseFloat(line, 64); err == nil {

// 	// 	if !math.IsNaN(f) && !math.IsInf(f, 0) {
// 	// 		float_str := strconv.FormatFloat(f, 'f', -1, 64)
// 	// 		if !strings.Contains(float_str, ".") {
// 	// 			return float_str + ".0"
// 	// 		}
// 	// 		return float_str
// 	// 	}
// 	if line[0] == '"' && line[len(line)-1] == '"' {
// 		return "'" + line[1:len(line)-1] + "'"
// 	} else {
// 		pattern := regexp.MustCompile(`\A-?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?\z`)
// 	}

// 	return fmt.Sprintf("ERR not a JSON literal: '%s'", line)

// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"regexp"
// )

// // TODO (number-grammar): implement per the lesson description.

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

// 	pattern := regexp.MustCompile(`\A-?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?\z`)

// 	for sc.Scan() {
// 		line := sc.Text()
// 		if line == "" {
// 			continue
// 		}
// 		if pattern.MatchString(line) {
// 			fmt.Println("OK")
// 		} else {
// 			fmt.Println("ERR invalid number")
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		tokens, err := Tokenize(line)

		if err != nil {
			fmt.Println("ERR")
		}
		parser := Parser{
			Tokens: tokens,
			Cursor: 0}
		val, err := ParseToken(&parser)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
		// if parser.peek().Kind == "PUNCT" && tokens[0].Value == "[" {
		// 	parsed_line := utilities.ParseArray(tokens[1:])

		// 	fmt.Println("[" + strings.Join(parsed_line, ", ") + "]")
		// }

		// fmt.Println(ParseToken(line))
	}
}

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

type Parser struct {
	Tokens []Token
	Cursor int // This is your incrementer
}

// Peak at the current character without moving forward
func (p *Parser) peek() Token {
	if p.Cursor >= len(p.Tokens) {
		return Token{"EOF", "EOF"} // End of input
	}
	return p.Tokens[p.Cursor]
}

// Consume the current character and move the incrementer forward
func (p *Parser) advance() Token {
	ch := p.peek()
	if ch.Kind != "EOF" {
		p.Cursor++
	}
	return ch
}

// TODO (scalar-values): implement per the lesson description.
func ParseToken(parser *Parser) (string, error) {
	// fmt.Println(tokens[0].Value)
	// fmt.Println(line)
	switch parser.peek().Value {
	case "true":
		return "True", nil
	case "false":
		return "False", nil
	case "null":
		return "None", nil
	case "[":
		parser.advance()
		result, err := ParseArray(parser)
		if err != nil {
			return "", err
		}
		return "[" + strings.Join(result, ", ") + "]", nil
	}
	// if b, err := strconv.ParseInt(line, 10, 64); err == nil {
	// 	if b == -0 {
	// 		return "0"
	// 	}
	// 	return line
	// } else if f, err := strconv.ParseFloat(line, 64); err == nil {

	// 	if !math.IsNaN(f) && !math.IsInf(f, 0) {
	// 		float_str := strconv.FormatFloat(f, 'f', -1, 64)
	// 		if !strings.Contains(float_str, ".") {
	// 			return float_str + ".0"
	// 		}
	// 		return float_str
	// 	}
	// } else
	if parser.peek().Value[0] == '"' && parser.peek().Value[len(parser.peek().Value)-1] == '"' {
		return "'" + parser.peek().Value[1:len(parser.peek().Value)-1] + "'", nil
	} else {
		pattern := regexp.MustCompile(`\A-?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?\z`)
		if f, err := strconv.ParseFloat(parser.peek().Value, 64); pattern.MatchString(parser.peek().Value) && err == nil {

			float_str := strconv.FormatFloat(f, 'f', -1, 64)

			return float_str, nil
		}
	}

	return "", fmt.Errorf("Invalid Input")
}

func ParseArray(parser *Parser) ([]string, error) {
	// var output []string
	// // var stack []string
	// if parser.peek().Kind == "PUNCT" && parser.peek().Value == "]" {
	// 	parser.advance()
	// 	return output
	// }
	// val, err := ParseToken(parser)
	// if err != nil {
	// 	return []string{"Error 2"}
	// }
	// output = append(output, val)
	// parser.advance()
	// // fmt.Println(tokens)
	// for parser.advance().Value == "," {

	// 	// print(parser.Tokens)
	// 	if parser.peek().Kind == "EOF" || parser.peek().Kind == "PUNCT" {
	// 		fmt.Println("ERR trailing comma")
	// 	}
	// 	// fmt.Println(parser.Cursor)
	// 	val, err = ParseToken(parser)
	// 	// fmt.Println(parser.Cursor)
	// 	if err != nil {
	// 		return []string{"Error 1"}
	// 	}
	// 	output = append(output, val)
	// 	parser.advance()

	// }
	// // if parser.peek().Value != "]" {
	// // 	fmt.Println("ERR no ending ]")
	// // }
	// parser.advance()
	// return output

	var output []string
	if parser.peek().Kind == "PUNCT" && parser.peek().Value == "]" {
		parser.advance()
		return output, nil
	}
	val, err := ParseToken(parser)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	output = append(output, val)
	parser.advance()
	for parser.peek().Kind == "PUNCT" && parser.peek().Value == "," {
		parser.advance()
		val, err := ParseToken(parser)
		if err != nil {
			return nil, fmt.Errorf("ERR trailing comma")
		}
		output = append(output, val)
		parser.advance()
	}
	// fmt.Println(parser.peek().Value)
	if parser.peek().Value != "]" {
		return nil, fmt.Errorf("Error 2")
	}

	return output, nil

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"unicode/utf16"
// )

// // TODO (string-escapes): implement per the lesson description.

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

// 	escape_map := map[rune]rune{
// 		'"':  '"',
// 		'\\': '\\',
// 		'/':  '/',
// 		'b':  '\u0008',
// 		'f':  '\u000c',
// 		'n':  '\u000a',
// 		'r':  '\u000d',
// 		't':  '\u0009',
// 	}
// 	for sc.Scan() {
// 		terminated := false
// 		decoded := ""
// 		err := false
// 		line := sc.Text()
// 		if line == "" {
// 			continue
// 		}

// 		for index := 0; index < len(line); {
// 			if index == len(line)-1 && line[index] == '"' {
// 				terminated = true
// 			}
// 			if line[index] == '\\' {
// 				if index+1 >= len(line) {
// 					fmt.Println("ERR")
// 					err = true
// 					return
// 				}
// 				next_val := line[index+1]
// 				val, exists := escape_map[rune(next_val)]

// 				if exists {
// 					decoded += string(val)
// 				} else if next_val == 'u' {
// 					if index+5 < len(line) {
// 						high := line[index+2 : index+6]
// 						high_val, _ := strconv.ParseUint(high, 16, 16)
// 						if high_val >= 0xD800 && high_val <= 0xDBFF {
// 							if index+11 < len(line) {
// 								low := line[index+8 : index+12]
// 								low_val, _ := strconv.ParseUint(low, 16, 16)
// 								result := string(utf16.Decode([]uint16{uint16(high_val), uint16(low_val)}))
// 								decoded += result
// 								index += 10
// 							}
// 						} else {
// 							// bytes, _ := hex.DecodeString(high)
// 							decoded += string(rune(high_val))
// 							index += 4
// 						}

// 					}
// 					// if next_val >= 0xD800 && next_val <= 0xDBFF {
// 					// 	fmt.Println("Surrogate Pair")
// 					// }
// 				} else {
// 					fmt.Println("ERR")
// 					err = true
// 					return
// 				}
// 				index++
// 				// } else if next_val == 'u'{

// 				// }
// 			} else {
// 				decoded += string(line[index])
// 			}
// 			index++

// 		}
// 		if !terminated {
// 			fmt.Println("ERR")
// 			return
// 		}

// 		if !err {
// 			fmt.Println(decoded[1 : len(decoded)-1])
// 		}

// 		// decoded, err := strconv.Unquote(line)
// 		// if err != nil {
// 		// 	fmt.Println("ERR")
// 		// }
// 		// fmt.Println(decoded)
// 	}
// }

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	for sc.Scan() {
// 		line := sc.Text()
// 		if line == "" {
// 			continue
// 		}
// 		toks, err := tokenize(line)
// 		if err != nil {
// 			fmt.Println("ERR", err)
// 			continue
// 		}
// 		for _, t := range toks {
// 			if t.Kind == "EOF" {
// 				fmt.Println("EOF")
// 			} else {
// 				fmt.Println(t.Kind, t.Value)
// 			}
// 		}
// 	}
// }
