package main

import (
	"fmt"
	"os"

	"github.com/chaosaffe/dolphin/pkg/nzbget"
)

func main() {

	// TODO: read address from config file

	c := nzbget.NewClient("http://10.0.60.101:30411")

	v, err := c.Version()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("NZBGet Version: %s\n", v)

	nzbs, err := c.ListNZBs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%d Active NZB's\n", len(nzbs))

}
