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
    flags:=[]string{"dbg","dir"}

    // default file
    dirFilnam := ""

    useStr := "./fileTreeTest [/dir=dirname] [/dbg]"
    helpStr := "program that lists all subdirs of a dir\n"

    if numarg > len(flags) + 1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: ./fileTreeTest /dir=dirname [/dbg]\n", useStr)
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

    val, ok := flagMap["dir"]
    if !ok {
		log.Fatalf("error -- need dir flag!\n")
    } else {
        if val.(string) == "none" {log.Fatalf("error -- no dirname provided with /dir flag!")}
        dirFilnam = val.(string)
        log.Printf("dir name: %s\n", dirFilnam)
    }

	info, err:=os.Stat(dirFilnam)
	if err !=nil {log.Fatalf("error -- Stat: %v\n",err)}
	if !info.IsDir() {log.Fatalf("error -- provided file is not a directory!\n")}

	log.Printf("success: %s is a directory\n", dirFilnam)

	startDir,err := os.Open(dirFilnam)
	if err !=nil {log.Fatalf("error -- open: %v\n",err)}

	names, err := startDir.Readdirnames(-1)
	for i:=0; i<len(names); i++ {
		fmt.Printf("  %d: %s\n",i+1, names[i])
	}
	dirNames := make([]string,len(names))
	fmt.Println(" *** directories ****")
	count:=-1
	for i:=0; i<len(names); i++ {
		filnam := dirFilnam + "/" + names[i]

		info, err:=os.Stat(filnam)
		if err !=nil {log.Fatalf("error -- Stat: %s: %v",filnam, err) }
		if info.IsDir() {count++; dirNames[count] = names[i];}
	}

	if count == -1 {fmt.Println("  no directories\n")} else {
		for i:=0; i< count+1; i++ {
			fmt.Printf("  %d: %s\n",i+1, dirNames[i])
		}
	}

	for i:=0; i< count+1; i++ {
		fmt.Printf("  %d: %s\n",i+1, dirNames[i])

	}


	log.Println("success end template!")
}
