// language-identifier identifies the languages in a given text.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/BluntSporks/lexicon"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Parse flags.
	lexDir := flag.String("lex", lex.DefaultDataPath(), "Directory of language files to use for identification")
	textFile := flag.String("text", "", "Name of text file to identify")
	flag.Parse()

	// Check flags.
	if *textFile == "" {
		log.Fatal("Missing -text parameter")
	}

	// Load the languages.
	langWords := lex.LoadAllLangs(*lexDir)

	// Open file.
	hdl, err := os.Open(*textFile)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Match words of at least two in length.
	wordRegExp := regexp.MustCompile(`\pL{2,}`)

	// Scan file line by line.
	langCnts := make(map[string]int)
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Process line
		lowLine := strings.TrimRight(line, " \t")
		if len(lowLine) > 0 {
			lowLine = strings.ToLower(line)
			matches := wordRegExp.FindAllString(line, -1)
			n := len(matches)
			if n < 3 {
				continue
			}

			// Check each position in each line.
			for i := 0; i < n-3; i++ {
				for lang := range langWords {
					// Assume that the result is found.
					found := true

					// Include exactly three words in each check.
					for j := 0; j < 3; j++ {
						match := matches[i+j]

						// All three matches must be found in the language, or else it is not found.
						if !langWords[lang][match] {
							found = false
							break
						}
					}

					// Increment the counter for this three-word match.
					if found {
						langCnts[lang]++
					}
				}
			}
		}
	}

	// Print the results.
	for lang, cnt := range langCnts {
		fmt.Printf("%s,%d\n", lang, cnt)
	}
}
