package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"github.com/DistAlchemist/Mongongo/service"
)

var (
	hostName = flag.String("hostname", "localhost", "mongongo server hostname")
	rpcPort  = flag.String("port", "1111", "rpc port to connect to mongongo server")
	prompt   = "mongongo"
	reader   *bufio.Reader
	cc       *rpc.Client
)

func printBanner() {
	fmt.Println("Welcome to Mongongo Command Line Interface!")
}

func processServerQuery(line string) {
	//
	args := service.ExecuteArgs{}
	reply := service.ExecuteReply{}
	args.Line = line
	err := cc.Call("Mongongo.ExecuteQueryOnServer", &args, &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	fmt.Printf("reply.result: %+v\n", reply.Result)
}

func processCLISTMT(line string) {
	fmt.Println("processing CLI statement")
}

func processLine(line string) {
	tokens := strings.Split(line, " ")
	token := strings.ToUpper(tokens[0])
	if strings.HasPrefix(token, "GET") ||
		strings.HasPrefix(token, "SELECT") ||
		strings.HasPrefix(token, "SET") ||
		strings.HasPrefix(token, "DELETE") ||
		strings.HasPrefix(token, "EXPLAIN") {
		processServerQuery(line)
	} else {
		processCLISTMT(line)
	}
}

func main() {
	flag.Parse()
	printBanner()
	reader = bufio.NewReader(os.Stdin)
	var err error
	cc, err = rpc.DialHTTP("tcp", *hostName+":"+*rpcPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Printf("Connected to %v:%v!\n", *hostName, *rpcPort)
	for {
		fmt.Print(prompt + "> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(line)
		processLine(line)
	}
	return
}
