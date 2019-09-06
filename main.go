package main

import (
	"fmt"
	"os"
)

func print_list(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: sslexpirystatsd <command> [argument]")
	fmt.Fprintln(os.Stderr, "Command: check <host:port> (Example: github.com:443)")
	fmt.Fprintln(os.Stderr, "Command: checkstatsd <host:port> (Example: github.com:443)")
	os.Exit(2)
}

func fatal_error(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exactly_arguments(arguments int) {
	if len(os.Args) != arguments {
		usage()
	}
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}
	switch os.Args[1] {
	case "check":
		exactly_arguments(3)
		seconds, err := hostCheck(os.Args[2])
		fatal_error(err)
		fmt.Printf("%s expires in %d seconds.\n", os.Args[2], seconds)
	case "checkstatsd":
		exactly_arguments(3)
		seconds, err := hostCheckStatsd(os.Args[2])
		fatal_error(err)
		fmt.Printf("%s expires in %d seconds.\n", os.Args[2], seconds)
	default:
		usage()
	}
}
