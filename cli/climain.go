package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	hostName = flag.String("hostname", "localhost", "mongongo server hostname")
	rpcPort  = flag.Int("rpc port", 1111, "rpc port to connect to mongongo server")
	prompt   = "mongongo"
	reader   *bufio.Reader
)

func printBanner() {
	fmt.Println("Welcome to Mongongo Command Line Interface!")
}

func main() {
	flag.Parse()
	printBanner()
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + "> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(line)
	}
	return
}
