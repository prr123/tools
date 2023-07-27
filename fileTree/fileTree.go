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

type dir struct {
	Name string
	SubDirList *[]dir
	NumSubDir int
}

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

	startDir := dir {
		Name: dirFilnam,
		SubDirList: nil,
		NumSubDir: -1,
	}

	nstartdir, err := getDir(startDir)
	if err !=nil {log.Fatalf("error -- getDir: %v", err)}

//	fmt.Printf("******* nest: %d **********\n", nest)
	fmt.Printf("ndirs: %s %d\n", nstartdir.Name, nstartdir.NumSubDir)
	for nest:=1; nest< 4; nest++ {
		if nest > 1 {
log.Printf("nest %d\n", nest)
			for j:=0; j< nstartdir.NumSubDir; j++ {
				jstartdir:= (*nstartdir.SubDirList)[j]
				nstartdir = jstartdir

				for i:=0; i< nstartdir.NumSubDir; i++ {
					nStartDir := (*nstartdir.SubDirList)[i]
					fmt.Printf(" %d: %s\n", i+1, nStartDir.Name)
					mstartDir, err := getDir(nStartDir)
					if err !=nil {log.Fatalf("error -- getDir: %v", err)}
					(*nstartdir.SubDirList)[i] = mstartDir
				}
			}
		} else {
log.Printf("nest 1\n")
		for i:=0; i< nstartdir.NumSubDir; i++ {
			nStartDir := (*nstartdir.SubDirList)[i]
			fmt.Printf(" %d: %s\n", i+1, nStartDir.Name)
			mstartDir, err := getDir(nStartDir)
			if err !=nil {log.Fatalf("error -- getDir: %v", err)}
			(*nstartdir.SubDirList)[i] = mstartDir
		}


		}
	}
	log.Println("success end template!")
}

func getDir(startDir dir)(endDir dir, err error) {

	dirFilnam := startDir.Name
	startDirFil,err := os.Open(dirFilnam)
	if err !=nil {return startDir, fmt.Errorf("error -- open: %v\n",err)}

	names, err := startDirFil.Readdirnames(-1)
	if err!=nil {return startDir, fmt.Errorf("Readdirnames: %v",err)}
	dirList := make([]dir, len(names))

	for i:=0; i<len(names); i++ {
		fmt.Printf("  %d: %s\n",i+1, names[i])
	}

	fmt.Println(" *** directories ****")
	count:=-1
	for i:=0; i<len(names); i++ {
		filnam := dirFilnam + "/" + names[i]

		info, err:=os.Stat(filnam)
		if err !=nil {log.Fatalf("error -- Stat: %s: %v",filnam, err) }
		if info.IsDir() {count++; dirList[count].Name = dirFilnam + "/" + names[i];}
	}
	sl :=dirList[:count+1]
	startDir.SubDirList = &sl
	startDir.NumSubDir = count

	return startDir, nil

}
