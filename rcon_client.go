package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	rcon "github.com/TF2Stadium/TF2RconWrapper"
	speakeasy "github.com/bgentry/speakeasy"
	isatty "github.com/mattn/go-isatty"
)

func stdinIsatty() bool {
	return isatty.IsTerminal(os.Stdin.Fd())
}

func readParameter(reader *bufio.Reader, prompt string) string {
	var str string

	if stdinIsatty() {
		str, _ = speakeasy.Ask(prompt)
	} else {
		str, _ = reader.ReadString('\n')
	}

	return strings.TrimSpace(str)
}

func main() {
	var addr, pwd string

	flag.StringVar(&addr, "addr", "", "Server Address")
	flag.StringVar(&pwd, "pwd", "", "Server RCON Password")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	isCLI := stdinIsatty()

	if addr == "" {
		addr = readParameter(reader, "Please enter server address: ")
	}

	if pwd == "" {
		pwd = readParameter(reader, "Please enter RCON password: ")
	}

	conn, err := rcon.NewTF2RconConnection(addr, pwd)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		if isCLI {
			fmt.Print("RCON> ")
		}

		query, readErr := reader.ReadString('\n')
		if readErr != nil {
			break
		}

		reply, err := conn.Query(query)
		if err != nil {
			if err == rcon.ErrUnknownCommand {
				log.Println("Unknown Command")
				continue
			}

			log.Fatal(err.Error())
		}
		fmt.Printf("%s\n", reply)
	}
}
