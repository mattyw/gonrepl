package main

import (
	bencode "code.google.com/p/bencode-go"
	"fmt"
	"net"
	"os"
)

type Response struct {
	Ex     string
	Value  string
	Status []string
}

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	nrepl := args[1]
	code := args[2]

	conn, err := net.Dial("tcp", nrepl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//fmt.Println("marhsalling")
	instruction := map[string]interface{}{
		"op":   "eval",
		"code": code,
	}
	err = bencode.Marshal(conn, instruction)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		result := Response{}
		//fmt.Println("UNmarhsalling")
		err = bencode.Unmarshal(conn, &result)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(result)
		if result.Ex != "" {
			fmt.Println(result.Ex)
		}

		if result.Value != "" {
			fmt.Println(result.Value)
		}

		if len(result.Status) > 0 && result.Status[0] == "done" {
			return
		}

	}
}
