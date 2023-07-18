// xplInterface.go
// program that explores the properties of the golang inteface type
// author: prr azul software
// date: 16 July 2023
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

type structA struct {
	Numb int
	Name string
}

type structB struct {
	Num int
	Nam string
}



func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg"}

    // default file
//    csrFilnam := "csrTest.yaml"

    useStr := "./xplInterface [/dbg]"
    helpStr := "program that explores the interface type\n"

    if numarg > 3 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s\n", useStr)
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

/*
    val, ok := flagMap["csr"]
    if !ok {
        log.Printf("default csrList: %s\n", csrFilnam)
    } else {
        if val.(string) == "none" {log.Fatalf("error: no yaml file provided with /csr flag!")}
        csrFilnam = val.(string)
        log.Printf("csrList: %s\n", csrFilnam)
    }
*/


	log.Printf("debug: %t\n", dbg)
//    log.Printf("Using csr file: %s\n", csrFilnam)

	var test interface{}

	testStructA := structA{
		Numb: 1,
		Name: "hello A",
	}

	testStructB := structB{
		Num: 1,
		Nam: "hello B",
	}

	PrintStructA(testStructA)
	PrintStructB(testStructB)
	fmt.Printf("test: %v\n", test)

	// we can assign any type to an interface

	fmt.Println("\n assigning interface test to a structA\n")
	test = testStructA
	fmt.Printf("test: %v\n", test)
	fmt.Printf("\nprinting test with type assertion!\n")
	PrintStructA(test.(structA))
	fmt.Printf("\n *** test assigned to structA **\n")
	PrintStruct(test)

	fmt.Printf("\n *** test assign to  structB  **\n")
	test = testStructB
	PrintStruct(test)

	log.Println("success xplInterface!")
}

func PrintStructA (stA structA) {
	fmt.Println("****** struct A *******")
	fmt.Printf("Numb: %d\n",stA.Numb)
	fmt.Printf("Name: %s\n",stA.Name)
	fmt.Println("**** end struct A *****")
}

func PrintStructB (stB structB) {
	fmt.Println("****** struct B *******")
	fmt.Printf("Num: %d\n",stB.Num)
	fmt.Printf("Nam: %s\n",stB.Nam)
	fmt.Println("**** end struct B *****")
}

func PrintStruct (test interface{}) {

	switch test.(type) {
	case structA:
		PrintStructA(test.(structA))
	case structB:
		PrintStructB(test.(structB))

	default:
		fmt.Println("unknown structure!")
	}

}
