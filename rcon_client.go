package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	rcon "github.com/TF2Stadium/TF2RconWrapper"
	isatty "github.com/mattn/go-isatty"
)

func main() {
	var addr, pwd string

	flag.StringVar(&addr, "addr", "", "Server Address")
	flag.StringVar(&pwd, "pwd", "", "Server RCON Password")

	flag.Parse()

	if addr == "" {
		log.Fatal("Address not specified")
	}
	if pwd == "" {
		log.Fatal("RCON password not specified")
	}

	conn, err := rcon.NewTF2RconConnection(addr, pwd)
	if err != nil {
		log.Fatal(err.Error())
	}

	reader := bufio.NewReader(os.Stdin)
	isCLI := isatty.IsTerminal(os.Stdin.Fd())

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
