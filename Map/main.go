package main

import "fmt"

func main() {
	colours := map[string]string{
		"Red":  "#ff000",
		"Blue": "#ff001",
	}
	colours["Green"] = "#0fffc"

	printMap(colours)
	delete(colours, "Red")
	printMap(colours)

	colours["White"] = "#3322"
	printMap(colours)
}

func printMap(colours map[string]string) {
	fmt.Println("---------")
	for k, v := range colours {
		fmt.Println(k, v)
	}
	fmt.Println("---------")
}
