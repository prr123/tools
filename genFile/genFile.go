// genFile.go
// program that generates a random file
//
// author: prr azul software
// date: 3/11/2023
// copyright (c) 2023 prr, azulsoftware
//

package main

import (
    "log"
    "fmt"
    "os"
	"math/rand"
    "time"

    util "github.com/prr123/utility/utilLib"
)

func GenRanData (rangeStart, rangeEnd int) (bdat []byte) {

    var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

    offset := rangeEnd - rangeStart

    randLength := seededRand.Intn(offset) + rangeStart
    bdat = make([]byte, randLength)

    charset := "abcdefghijklmnopqrstuvw0123456789"
    for i := range bdat {
        bdat[i] = charset[seededRand.Intn(len(charset)-1)]
    }
    return bdat
}


func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg","name","size"}

    useStr := "./genFile /name=filnam /size=n [/dbg]"
    helpStr := "program that generates files with random content of size n\n"

    if numarg > len(flags)+1 {
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

	dbg:= false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
		fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

	filnam :=""
    namval, ok := flagMap["name"]
    if !ok {
        log.Fatalf("no name flag provided in cli!")
    } else {
        if namval.(string) == "none" {log.Fatalf("-- error: no file name  provided with /name flag!")}
        filnam = namval.(string)
    }

	sizeStr :=""
    sizval, ok := flagMap["size"]
    if !ok {
        log.Fatalf("no size flag provided in cli!")
    } else {
        if sizval.(string) == "none" {log.Fatalf("-- error: size value provided with /size flag!")}
        sizeStr = sizval.(string)
    }

	if dbg {
		fmt.Println("********** cli flags and their parameters *****")
		fmt.Printf("debug:     %t\n", dbg)
    	fmt.Printf("file name: %s\n", filnam)
    	fmt.Printf("file size: %s\n", sizeStr)
		fmt.Println("******* end cli flags and their parameters ****")
	}

	// check whether file exists


	// if not create one

	filsiz, err := CvtSize(sizeStr)
	if err != nil { log.Fatalf("file size is not convertible: %v\n",err)} 
	fmt.Printf("file size: %d\n", filsiz)
	log.Println("success genFile!")
}


func CvtSize(sizeStr string) (siz int64, err error) {

	// check last letter of size
	let := sizeStr[len(sizeStr) -1]
	fmt.Printf("last letter: %q ", let)

	// if last letter is a letter, convert the rest into a number
	var mult int64 = 1
	switch let {
	case 'K':
		mult = 1000

	case 'M':
		mult = 1000000

	case 'G':
		mult = 1000000000

//	case 
//		fmt.Printf("number")

	default:
		if !util.IsNumeric(let) {
			return -1 , fmt.Errorf("let is not alphaNumeric!")
		}
	}
	fmt.Printf("size mult: %d\n", mult)

	return siz, nil
}


