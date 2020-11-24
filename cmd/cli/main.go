// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterh/liner"

	"github.com/DistAlchemist/Mongongo/service"
)

var (
	hostName  = flag.String("hostname", "localhost", "mongongo server hostname")
	rpcPort   = flag.String("port", "1111", "rpc port to connect to mongongo server")
	prompt    = "mongongo"
	reader    *bufio.Reader
	cc        *rpc.Client
	historyFn = filepath.Join(os.TempDir(), ".liner_example_history")
	names     = []string{"get", "GET", "set", "SET", "select", "SELECT",
		"delete", "DELETE", "explain", "EXPLAIN"}
	line *liner.State
)

func printBanner() {
	fmt.Println("Welcome to Mongongo Command Line Interface!")
	printHelp()
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
	log.Printf("reply.result: %+v\n", reply.Result)
	for k, v := range reply.Result.ResultSet {
		log.Printf("%v: %v\n", k, v)
	}
}

func printHelp() {
	fmt.Printf("Usage: (currently supported)\n")
	fmt.Printf("\tSET tableName.columnFamilyName['rowKey']['column']='value'\n")
	fmt.Printf("keywords(case insensitive): SET, GET, SELECT, DELETE, EXPLAIN\n\n")
	fmt.Printf("press Ctrl-C or type exit to quit\n\n")
}

func processCLISTMT(line string) {
	if strings.HasPrefix(line, "HELP") {
		printHelp()
	} else if strings.HasPrefix(line, "EXIT") {
		quitCli()
		os.Exit(0)
	}
	log.Println("processing CLI statement")
}

func processLine(line string) {
	tokens := strings.Split(line, " ")
	tokens[0] = strings.ToUpper(tokens[0])
	token := tokens[0]
	line = strings.Join(tokens, " ")
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

func quitCli() {
	if f, err := os.Create(historyFn); err != nil {
		log.Print("Err writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	line.Close()
}

func main() {
	// command line tools
	line = liner.NewLiner()
	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, line) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(historyFn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	// parse flags
	flag.Parse()
	printBanner()
	reader = bufio.NewReader(os.Stdin)

	// setup connection to server
	var err error
	cc, err = rpc.DialHTTP("tcp", *hostName+":"+*rpcPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	log.Printf("Connected to %v:%v!\n", *hostName, *rpcPort)

	// start command line interface
	for {
		if name, err := line.Prompt(prompt + "> "); err == nil {
			processLine(name)
			line.AppendHistory(name)
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Err reading line: ", err)
		}
	}
	quitCli()
	return
}
