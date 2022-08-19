package main

import (
	"github.com/phial3/accumulate/choke"
	"github.com/phial3/accumulate/tcp/socks5"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/22
 */

func main() {
	svr := socks5.NewServer("127.0.0.1:10200")
	go svr.Run()
	choke.Choke()
}
