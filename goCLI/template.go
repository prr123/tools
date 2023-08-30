// template.go
// program that has a cli template
// author: prr azul software
// date: 5 July 2023
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
    flags:=[]string{"dbg","csr"}

    // default file
    csrFilnam := "csrTest.yaml"

    useStr := "./template [/csr=csrfile] [/dbg]"
    helpStr := "template for main program with a cli flag parser\n"

    if numarg > 4 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: ./template [/flag1=] [/flag2]\n", useStr)
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

    val, ok := flagMap["csr"]
    if !ok {
        log.Printf("default csrList: %s\n", csrFilnam)
    } else {
        if val.(string) == "none" {log.Fatalf("error: no yaml file provided with /csr flag!")}
        csrFilnam = val.(string)
        log.Printf("csrList: %s\n", csrFilnam)
    }

	log.Printf("debug: %t\n", dbg)
    log.Printf("Using csr file: %s\n", csrFilnam)

	log.Println("success end template!")
}
