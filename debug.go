package main

import (
	"fmt"
	"reflect"
)

func (p Panel) PrintRowTypes() {
	PrintRowTypes(p)

}

func PrintRowTypes(p Panel) {
	for head, col := range p {
		for row, val := range col {
			switch t := val.(type) {
			case float64:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case int:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case bool:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case string:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			default:
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, reflect.TypeOf(val))
			}
		}
	}
}
