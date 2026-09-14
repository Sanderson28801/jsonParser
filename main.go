package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO (scalar-values): implement per the lesson description.
func ParseToken(line string) string {
	switch line {
	case "true":
		return "True"
	case "false":
		return "False"
	case "null":
		return "None"
	}
	if _, err := strconv.ParseInt(line, 10, 64); err == nil {
		return line
	} else if f, err := strconv.ParseFloat(line, 64); err == nil {
		float_str := strconv.FormatFloat(f, 'f', -1, 64)
		if !strings.Contains(float_str, ".") {
			return float_str + ".0"
		}
		return float_str
	} else if line[0] == '"' && line[len(line)-1] == '"' {
		return "'" + line[1:len(line)-1] + "'"
	}

	return fmt.Sprintf("ERR not a JSON literal: '%s'", line)

}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		fmt.Println(ParseToken(line))
	}
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

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Token struct{ Kind, Value string }

// func tokenize(src string) ([]Token, error) {
// 	var tokens []Token
// 	i := 0
// 	for i < len(src) {
// 		c := src[i]
// 		switch {
// 		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
// 			i++
// 		case strings.ContainsRune("{}[],:", rune(c)):
// 			tokens = append(tokens, Token{"PUNCT", string(c)})
// 			i++
// 		case c == '"':
// 			j := i + 1

// 			for j < len(src) && src[j] != '"' {
// 				j++
// 			}
// 			decoded, err := strconv.Unquote(src[i : j+1])

// 			if err != nil {
// 				fmt.Println("Unexpected Token")
// 			}
// 			tokens = append(tokens, Token{"STRING", decoded})
// 			i = j + 1
// 		case c == '-' || (c >= '0' && c <= '9'):
// 			j := i
// 			if c == '-' {
// 				j++
// 			}
// 			for j < len(src) && src[j] >= '0' && src[j] <= '9' {
// 				j++
// 			}

// 			if j < len(src) && src[j] == '.' {
// 				j++
// 				for j < len(src) && ((src[j] >= '0' && src[j] <= '9') || src[j] == 'e') {
// 					j++
// 				}
// 			}

// 			tokens = append(tokens, Token{"NUMBER", src[i:j]})
// 			i = j
// 		case strings.HasPrefix(src[i:], "true"):
// 			tokens = append(tokens, Token{"TRUE", "true"})
// 			i += 4
// 		case strings.HasPrefix(src[i:], "false"):
// 			tokens = append(tokens, Token{"FALSE", "false"})
// 			i += 5
// 		case strings.HasPrefix(src[i:], "null"):
// 			tokens = append(tokens, Token{"NULL", "null"})
// 			i += 4
// 		default:
// 			return nil, fmt.Errorf("unexpected character %q at position %d", c, i)
// 		}
// 	}
// 	tokens = append(tokens, Token{"EOF", ""})
// 	return tokens, nil
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
