package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

type Word struct {
	content                  string
	endsWithComma            bool
	endsWithColon            bool
	endsWithSemicolon        bool
	endsWithPeriod           bool
	endsWithDash             bool
	endsWithQuestionMark     bool
	endsWithExclaimationMark bool
}

func stripPunctuation(input string) string {
	out := input
	for {
		if out[len(out)-1:] == "," {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == ":" {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == ";" {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == "." {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == "-" {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == "?" {
			out = out[:len(out)-1]
			continue
		}

		if out[len(out)-1:] == "!" {
			out = out[:len(out)-1]
			continue
		}
		break
	}
	return out
}

func ToAtinLerb(word string, capitalize bool) string {
	if len(word) < 3 {
		if capitalize {
			return strings.ToUpper(word[:1]) + word[1:]
		}
		return word
	}
	out := word[1:]
	erb := strings.ToLower(word[:1]) + "erb"

	if capitalize {
		out = strings.ToUpper(out[:1]) + out[1:]
	}
	return out + "-" + erb
}

func WordsToAtinLerb(input string) string {
	words := strings.Split(input, " ")
	var parsedWords []Word

	var out []string

	for _, w := range words {
		word := Word{stripPunctuation(w), w[len(w)-1:] == ",", w[len(w)-1:] == ":", w[len(w)-1:] == ";", w[len(w)-1:] == ".", w[len(w)-1:] == "-", w[len(w)-1:] == "?", w[len(w)-1:] == "!"}
		parsedWords = append(parsedWords, word)
	}

	for i, w := range parsedWords {
		o := ToAtinLerb(w.content, i == 0)

		if w.endsWithComma {
			o += ","
		}

		if w.endsWithColon {
			o += ":"
		}

		if w.endsWithSemicolon {
			o += ";"
		}

		if w.endsWithPeriod {
			o += "."
		}

		if w.endsWithDash {
			o += "-"
		}

		if w.endsWithQuestionMark {
			o += "?"
		}

		if w.endsWithExclaimationMark {
			o += "!"
		}

		out = append(out, o)
	}

	return strings.Join(out, " ")
}

func main() {
	// Create new parser object
	parser := argparse.NewParser("Erb-ferb Atin-lerb Translator", "translates the given text to Ferb Latin")
	// Create string flag
	gottentext := parser.String("t", "text", &argparse.Options{Required: true, Help: "Text to translate into Ferb Latin"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	/*
		a := &gottentext

		b := &a

		c := &b

		d := &c

		e := &d

		f := &e

		fmt.Println(*******f)
	*/

	fmt.Println(WordsToAtinLerb(*gottentext))
}
