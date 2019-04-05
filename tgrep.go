package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	p := flag.String("fpath", "testfile", "file path")
	flag.Parse()
	file, err := os.Open(*p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Let's find stuff, what word do you want to find?")
	var searchWord string
	fmt.Scanln(&searchWord)
	var searchWordLen int
	searchWordLen = len(searchWord)
	s := bufio.NewScanner(file)
	for s.Scan() {
		if strings.Contains(s.Text(), searchWord) {
			var start int
			start = strings.Index(s.Text(), searchWord)
			var text string
			text = s.Text()
			var output []string

			output = append(output, text[0:start])
			red := color.New(color.FgRed).SprintFunc()
			redText := text[start:(start + searchWordLen)]
			output = append(output, red(redText))
			text = text[(start + searchWordLen):len(text)]

			for strings.Index(text, searchWord) > -1 {
				start = strings.Index(text, searchWord)
				output = append(output, text[0:start])
				redText = text[start:(start + searchWordLen)]
				output = append(output, red(redText))
				text = text[(start + searchWordLen):len(text)]
			}

			output = append(output, text)
			fmt.Println(strings.Join(output, ""))
		}
	}

}
