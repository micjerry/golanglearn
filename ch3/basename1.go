package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for  _,args := range os.Args[1:] {
		fmt.Println(basename(args))
		fmt.Println(basenamestr(args))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	
	return s
}

func basenamestr(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}