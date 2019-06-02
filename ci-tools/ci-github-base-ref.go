package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"pytest-queries-bot/pytestqueries/ci-tools/utils"
)

func main() {
	data := utils.GqlQueryData{}
	if _, err := flags.Parse(&data); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	hash, err := utils.GetBaseRef(data)

	if err != nil {
		fmt.Printf("Failed to get: %s", err)
		os.Exit(1)
	}

	if _, err := os.Stdout.WriteString(hash); err != nil {
		panic(err)
	}
}
