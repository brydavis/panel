package main

import (
	"fmt"
)

// TODO
// - Transform Go structures into frames

func main() {

	{
		df := New(map[string][]string{
			"Q": []string{"l", "o", "r", "q"},
			"N": []string{"5", "2", "9", "7"},
			"M": []string{"false", "true", "true", "false"},
		})

		df.Dtypes()
		fmt.Println(df.Sum("N"))

	}

	// {
	// 	df1 := Read("./data/gota.csv", true)

	// 	fil := df1.
	// 		Filter("a", func(a string) bool {
	// 			return a == "a"
	// 		}).
	// 		Filter("b", func(b float64) bool {
	// 			return b >= 4.0

	// 		}).
	// 		Filter("d", func(d bool) bool {
	// 			return d == true
	// 		})

	// 		// df1.PrintRowTypes()
	// 	fmt.Println(fil)

	// }

	// { // Read a URL
	// 	// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
	// 	// link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
	// 	link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

	// 	df := Read(link, true)
	// 	fmt.Println(df["gpa"])
	// }

	// { // Read a CSV
	// 	df := Read("data/iris.csv", true)
	// 	df = df.Map("species", func(s string) string {
	// 		return fmt.Sprintf("__%s", s)
	// 	})

	// 	fmt.Println(df["species"])

	// }

	// {
	// 	df := Read("hello.json", false)
	// 	if df == nil {
	// 		fmt.Println("nothing there")
	// 	}
	// }
}
