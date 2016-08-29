package main

import "fmt"

type Analysis map[string]interface{}

func (p Panel) OLS(output string, inputs ...string) {
	return OLS(p, output, inputs)
}

// Regression analysis in a box
func OLS(p Panel, output string, inputs ...string) {
	fmt.Printf("output: %s\n", output)
	for i := range inputs {
		fmt.Printf("input: %s\n", inputs[i])
	}
}
