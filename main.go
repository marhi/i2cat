package main

import (
	"github.com/marhi/i2cat/i2catlib"
	"os"
	"fmt"
)

func main()  {
	if len(os.Args) < 2 {
		i2catlib.PrintImg(os.Stdin)
		os.Exit(0)
	}

	files := os.Args[1:]
	for _, v := range files {
		fh, err := os.Open(v)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", os.Args[1])
			os.Exit(1)
		}

		i2catlib.PrintImg(fh)

		fh.Close()
	}
}