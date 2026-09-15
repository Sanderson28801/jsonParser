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
	"jsonParser/utilities"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		tokens, err := utilities.Tokenize(line)

		if err != nil {
			fmt.Println("ERR")
		}
		parser := utilities.Parser{
			Tokens: tokens,
			Cursor: 0}
		val, err := utilities.ParseToken(&parser)
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
