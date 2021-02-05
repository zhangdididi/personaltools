package main

import (
	"fmt"
	"testing"
)

func TestTrimString(t *testing.T) {
	strs := []string{" srgjsojrdglsrg ajkdljglkaj  ", "fasklefjalkjgia ", "54252" + "w2342 ", "sfag "}
	for _, tt := range strs {
		t.Run("字符串", func(t *testing.T) {
			fmt.Println(TrimSpace(tt))
		})
	}
}
