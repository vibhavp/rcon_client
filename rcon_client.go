package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	rcon "github.com/TF2Stadium/TF2RconWrapper"
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

	for {
		fmt.Print("RCON> ")
		query, _ := reader.ReadString('\n')
		reply, err := conn.Query(query)
		if err != nil {
			if err == rcon.UnknownCommandError {
				log.Println("Unknown Command")
				continue
			}

			log.Fatal(err.Error())
		}
		fmt.Printf("%s\n", reply)
	}
}
