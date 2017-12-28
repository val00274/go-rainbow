package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
)

func selectColor(i int) *color.Color {
	i = i % 24

	switch i / 6 {
	case 0:
		return color.New(color.Attribute(31 + (i % 6)))
	case 1:
		return color.New(color.Attribute(91 + (i % 6)))
	case 2:
		return color.New(color.Underline, color.Attribute(31+(i%6)))
	case 3:
		return color.New(color.Underline, color.Attribute(91+(i%6)))
	default:
		panic(i)
	}
}

func selectComma() rune {
	custom_comma := flag.String("comma", ",", "comma character")
	is_csv := flag.Bool("csv", false, "filetype csv")
	is_tsv := flag.Bool("tsv", false, "filetype tsv")
	flag.Parse()

	switch {
	case *is_csv:
		return ','
	case *is_tsv:
		return '\t'
	default:
		return []rune(*custom_comma)[0]
	}
}

func main() {
	comma := selectComma()

	r := csv.NewReader(os.Stdin)
	r.Comma = comma

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for i, row := range record {
			if i != 0 {
				fmt.Print(string(comma))
			}
			selectColor(i).Print(row)
		}
		fmt.Println()
	}
}
