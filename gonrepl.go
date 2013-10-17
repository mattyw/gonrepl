package main

import (
	bencode "code.google.com/p/bencode-go"
	"fmt"
	"net"
	"os"
)

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

	//We can read multiple values, but just read the first
	result := map[string]string{}
	//fmt.Println("UNmarhsalling")
	err = bencode.Unmarshal(conn, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(result)
	ex, ok := result["ex"]
	if ok {
		fmt.Println(ex)
	}
	value, ok := result["value"]
	if ok {
		fmt.Println(value)
	}
}
