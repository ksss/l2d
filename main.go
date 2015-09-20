package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n = flag.Int("n", 1, "print every `n' lines")
	d = flag.String("d", ".", "print this every lines")
)

func main() {
	flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	i := 0
	for s.Scan() {
		if (i % *n) == 0 {
			fmt.Print(*d)
			i = 0
		}
		i++
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Print("\n")
}
