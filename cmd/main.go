package main

import (
	"bufio"
	"encoding/json"
	"github.com/tangtj/gorediscli/cli"
	"log"
	"net"
	"os"
)

func main() {

	c, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		print("> ")
		line, _, _ := reader.ReadLine()
		line = cli.Convert2Command(line)
		c.Write(line)
		r, _ := cli.Resp(c)
		str, _ := json.Marshal(r)
		println(string(str))
	}
}
