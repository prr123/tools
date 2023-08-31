// kvServer.go
// kv store api
// author: prr azul software
// date: 31 Aug 2023
// copyright 2023 prr, azulsoftware
//

package main

import (
    "log"
    "fmt"
    "os"
//    "time"

    util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","port"}

    // default file
    useStr := "./kvServer [/port=14000] [/dbg]"
    helpStr := "server for kv api\n"
	helpStr += "API is:\n"
	helpStr += " -- add: adr:port/kv/add\n"
	helpStr += " -- upd: adr:port/kv/upd\n"
	helpStr += " -- get: adr:port/kv/get\n"
	helpStr += " -- del: adr:port/kv/del\n"

    if numarg > 3 {
        fmt.Println("too many arguments in cl!")
		fmt.Printf("usage is: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
		fmt.Printf("help: %s\n", helpStr)
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)
	}

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
		fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

	portStr := "14000"
    val, ok := flagMap["port"]
    if ok {
        portStr = val.(string)
    }

	log.Printf("Debug: %t\n", dbg)
    log.Printf("kvServer listening on port: %s\n", portStr)

	log.Println("success kvServer!")
}
