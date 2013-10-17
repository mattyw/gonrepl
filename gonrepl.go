package main

import (
	bencode "code.google.com/p/bencode-go"
	"fmt"
	"io/ioutil"
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
	if len(args) < 2 {
		return
	}
	nrepl := args[1]
	code := ""

	if len(args) == 2 {
		//code := args[2]
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
		code = string(b)
	} else if len(args) == 3 {
		code = args[2]
	} else {
		fmt.Println(`
        Usage:
        gonrepl host:port code
        gonrepl host:port code from stdin
        `)
		return
	}

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
