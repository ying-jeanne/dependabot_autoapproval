package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
)

func main() {
	color.Green("Hello, World!")

	// Example usage of go-cmp library
	want := []int{1, 2, 3}
	got := []int{1, 2, 4}
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println("Difference found:")
		fmt.Println(diff)
	}
}
