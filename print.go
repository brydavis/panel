package main

// May modify with delimit options: func (p Panel) String(delim string) string
// For example, p.String(",")... How would you pass delim in fmt.Println(p)???
//

import "fmt"

func (p Panel) Print() {
	// Pretty way????
	fmt.Println(p)
}

func (p Panel) String() string {
	// width := p.Size().Width
	length := p.Size().Length

	var res string
	for row := 0; row < length; row++ {
		s := ""
		for head, _ := range p {
			s += fmt.Sprintf("%v\t", p[head][row])
		}
		res += fmt.Sprintf("%s\n", s)
	}

	return res
}
