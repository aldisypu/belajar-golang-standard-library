package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aldi Syah", "Aldi"))
	fmt.Println(strings.Split("Aldi Syah", " "))
	fmt.Println(strings.ToLower("Aldi Syah"))
	fmt.Println(strings.ToUpper("Aldi Syah"))
	fmt.Println(strings.Trim("      Aldi Syah      ", " "))
	fmt.Println(strings.ReplaceAll("Aldi Syah Aldi Syah", "Aldi", "Budi"))
}
