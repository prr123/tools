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
	"strconv"

    util "github.com/prr123/utility/utilLib"
)


func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg","name","size", "two"}

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
/*
    if dbg {
		fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }
*/

	two:= false
    _, ok = flagMap["two"]
    if ok {two = true}

	filnam :=""
    namval, ok := flagMap["name"]
    if !ok {
        log.Fatalf("error -- cli: no name flag provided in cli!")
    } else {
        if namval.(string) == "none" {log.Fatalf("error -- cli: no file name  provided with /name flag!")}
        filnam = namval.(string)
    }

	sizeStr :=""
    sizval, ok := flagMap["size"]
    if !ok {
        log.Fatalf("error -- cli: no size flag provided in cli!")
    } else {
        if sizval.(string) == "none" {log.Fatalf("error -- clis: no size value provided with /size flag!")}
        sizeStr = sizval.(string)
    }

	if dbg {
		fmt.Println("********** cli flags and their parameters *****")
		fmt.Printf("debug:     %t\n", dbg)
    	fmt.Printf("file name: %s\n", filnam)
    	fmt.Printf("file size: %s\n", sizeStr)
		fmt.Println("******* end cli flags and their parameters ****")
	}

	filsiz, err := CvtSize(sizeStr, two)
	if err != nil { log.Fatalf("error -- file size is not convertible: %v\n",err)} 
	fmt.Printf("file size[%t]:  %d\n", two, filsiz)

	fil, err := CreateFile(filnam)
	if err != nil {log.Fatalf("error -- CreateFile: %v\n", err)}

	// last parameter is option
	// currently only opt val 0 implemented
	err = FillFile(fil, filsiz, 0)
	fil.Close()
	if err != nil {log.Fatalf("error -- CreateFile: %v\n", err)}

	log.Println("success genFile!")
}


func FillFile(fil *os.File, filsiz uint64, opt int) (err error) {

	if fil==nil {return fmt.Errorf("no file pointer!")}

//	bdat := make([]byte, filsiz)
	bdat := []byte{}
	switch opt {
	//alpha numeric
	case 0:
		bdat = GenData(int(filsiz))
	default:
		return fmt.Errorf("unknown option!")
	}

	n, err := fil.Write(bdat)
	if err != nil {return fmt.Errorf("os.Write: %v", err)}
	fmt.Printf("size: %d len: %d n: %d\n", filsiz, len(bdat), n)
	return nil
}

func CvtSize(sizeStr string, two bool) (siz uint64, err error) {

	// check last letter of size
	let := sizeStr[len(sizeStr) -1]
	fmt.Printf("last letter: %q ", let)

	// if last letter is a letter, convert the rest into a number
	var mult uint64 = 1
	switch let {
	case 'K':
		mult = 1000

	case 'M':
		mult = 1000000

	case 'G':
		mult = 1000000000

	default:
		if !util.IsNumeric(let) {
			return 0 , fmt.Errorf("let is not alphaNumeric!")
		}
	}
//	fmt.Printf("size mult: %d\n", mult)

	intStr:=""
    if mult>1 {
        intStr = string(sizeStr[:len(sizeStr)-1])
    } else {
        intStr = string(sizeStr[:len(sizeStr)])
    }

    inum, err := strconv.Atoi(intStr)
    if err !=nil {return 0, fmt.Errorf("error -- cannot convert intStr: %s: %v", intStr, err)}
    num := uint64(inum)*uint64(mult)

    fmt.Printf("res: %d\n", inum)

	if !two {return num, nil}

    num--
    num = num | num>>1
    num = num | num>>2
    num = num | num>>4
    num = num | num>>8
    num = num | num>>16
    num = num | num>>32
    num++

	return num, nil
}

func CreateFile(filnam string)(fn *os.File, err error) {

	// check whether file exists
	if _, err := os.Stat(filnam); err == nil {
		return nil, fmt.Errorf("file <%s> does exist!", filnam)
	}
	// if not create one
	fn, err = os.Create(filnam)
	if err != nil {return nil, fmt.Errorf("osCreate: %v", err)}
	return fn, nil
}



func GenData (siz int) (bdat []byte) {

    var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

    bdat = make([]byte, siz)

    charset := "abcdefghijklmnopqrstuvw0123456789"
    for i := range bdat {
        bdat[i] = charset[seededRand.Intn(len(charset)-1)]
    }
    return bdat
}

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
