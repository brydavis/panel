package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

// 1. Loading data from file or memory
// 2. Perform data cleaning and standardization

func Load(fname string, head bool) Panel {
	p := make(Panel)

	switch path.Ext(fname) {
	case ".csv":
		h := []string{}
		f, _ := os.Open(fname)
		headless := false
		if !head {
			headless = true
		}

		r := csv.NewReader(bufio.NewReader(f))
		for {
			if head {
				line, _ := r.Read()
				for _, val := range line {
					h = append(h, strings.ToLower(val))
				}
				head = false
			} else {
				line, err := r.Read()
				if err == io.EOF {
					break
				}
				for key, val := range line {
					if headless {
						k := strconv.Itoa(key)
						p[k] = append(p[k], val)
					} else {
						p[h[key]] = append(p[h[key]], val)
					}
				}
			}
		}

	case ".xml":

	case ".tsv":

	case ".json":

	case ".xlsx":
		file, err := xlsx.OpenFile(fname)
		if err != nil {
			fmt.Println(err)
		}
		for _, sheet := range file.Sheets {
			for _, row := range sheet.Rows {
				for _, cell := range row.Cells {
					c, _ := cell.String()
					fmt.Printf("%s\n", c)
				}
			}
		}
	}

	return p.Clean()
}
