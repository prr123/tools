// filetreeRev.go
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
	"strings"
    "time"

    util "github.com/prr123/utility/utilLib"
)

type nestTree struct {
	nest []dirList
}

type dirList struct {
	dirlist []dirObj
}

type filInfo struct {
	Name string
	Size int64
	Modified time.Time
}

type dirObj struct {
	Name string
	SubDirList []string
	FilList []filInfo
}

func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","nest", "dir"}

    // default file
    dirFilnam := ""

    useStr := "./filetreeRev [help] |[/dir=dirname] [/nest=nestnum] [/dbg]"
    helpStr := "program that lists all subdirs of a dir\n"

    if numarg > len(flags) + 1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
		fmt.Printf("help: %s\n", helpStr)
		fmt.Printf("usage: %s\n", useStr)
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

    nval, ok := flagMap["nest"]
    if !ok {
		log.Fatalf("error -- need nest flag!\n")
    }
	if nval.(string) == "none" {log.Fatalf("error -- no max nest level  provided with /nest flag!")}
	nestStr := nval.(string)
	if dbg {log.Printf("dbg -- nest Level str: %s\n", nestStr)}
	maxNest := 0
	_, err = fmt.Sscanf(nestStr,"%d",&maxNest)
	if err != nil {log.Fatalf("error -- nest %s is not an integer! %v", nestStr, err)}
	log.Printf("max nest level: %d\n", maxNest)



	maxNest = 3
	fileTree:= nestTree{}
	fileTree.nest = make([]dirList, maxNest)
//	fmt.Printf("nestList: %v\n", filTree)


	dir, err := getDir(dirFilnam)
	if err !=nil {log.Fatalf("error -- getDir: %v", err)}

	list := make([]dirObj, 1)
	list[0] = *dir
	fileTree.nest[0].dirlist = list


	PrintFileTree(&fileTree)

//	nextLevel := true
	for inest:=1; inest< maxNest; inest++ {

		// let's determine how many sirs there are in the new nesting level
		// For that we need to count the sub dir of the previous level
		dirs := fileTree.nest[inest-1]
		subTot :=0
		for i:=0; i< len(dirs.dirlist); i++ {
			dir := dirs.dirlist[i]
			subTot += len(dir.SubDirList)
		}
//
//		if subTot == 0 {nextLevel = false; break}
		fileTree.nest[inest].dirlist = make([]dirObj, subTot)

		ic := -1
		fmt.Printf("nest: %d dir list: %d\n", inest, len(dirs.dirlist))
		for i:=0; i< len(dirs.dirlist); i++ {
			numSubDir := len(fileTree.nest[inest-1].dirlist[i].SubDirList)
			fmt.Printf("subdir list: %d\n", numSubDir)
			for j:=0; j< numSubDir; j++ {
				ic++
				dirNam:=fileTree.nest[inest-1].dirlist[i].SubDirList[j]
				dir, err := getDir(dirNam)
				fileTree.nest[inest].dirlist[ic]=*dir
				if err != nil {log.Fatalf("error -- getDir nest: %d, subdir: %i - %v", inest, i, err)}
			}
		}
		PrintFileTree(&fileTree)
	}
	PrintFileTree(&fileTree)
}

func getDir(dirFilnam string)(dir *dirObj, err error) {

	parDir := dirObj{
		Name: dirFilnam,
	}
//	fmt.Printf("************* directory %s ****************\n", dirFilnam)
	startDirFil,err := os.Open(dirFilnam)
	if err !=nil {return nil, fmt.Errorf("error -- open: %v\n",err)}

	names, err := startDirFil.Readdirnames(-1)
	if err!=nil {return nil, fmt.Errorf("Readdirnames: %v",err)}
	dirList := make([]string, len(names))
	filList := make([]filInfo, len(names))

	dirCount:=-1
	filCount:=-1
	for i:=0; i<len(names); i++ {
		filnam := dirFilnam + "/" + names[i]
		info, err:=os.Lstat(filnam)
		if err !=nil {
			errStr:=err.Error()
			idx := strings.Index(errStr, "no such file")
			if idx == -1 {return nil, fmt.Errorf("Lstat: %s: %v",filnam, err) }
			continue
		}
		switch mode := info.Mode(); {
		case mode.IsRegular():
//			fmt.Println("regular file")
			filCount++
			filList[filCount].Name = names[i]
			filList[filCount].Size = info.Size()
			filList[filCount].Modified = info.ModTime()
		case mode.IsDir():
//			fmt.Println("directory")
			dirCount++
			dirList[dirCount] = dirFilnam + "/" + names[i]
		default:
		}
/*
	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&fs.ModeNamedPipe != 0:
		fmt.Println("named pipe")

*/
	}

	slFil := filList[:filCount+1]
	parDir.FilList = slFil
	slDir :=dirList[:dirCount+1]
	parDir.SubDirList = slDir

	return &parDir, nil
}

func PrintDirObj(dir dirObj) {

	dirList := dir.SubDirList
	numDirs := len(dirList)
	fmt.Printf("********** Dir: %s *********************\n", dir.Name)
	if numDirs == 0 {
		fmt.Printf("***** No Sub-Directories *******\n")
	} else {
		fmt.Printf("***** Sub-Directories: %d *******\n", numDirs)
		for i:=0; i< numDirs; i++ {
			dirnam := dirList[i]
			fmt.Printf("  %d: %s\n", i+1, dirnam)
		}
	}
	filList := dir.FilList
	numFiles := len(filList)
	if numFiles == 0 {
		fmt.Printf("********** No Files *********\n")
	} else {
		fmt.Printf("********** Files: %d *********\n", numFiles)
		for i:=0; i< numFiles; i++ {
			filnam := filList[i].Name
			size := filList[i].Size
			mod := filList[i].Modified
			fmt.Printf("  %d: %-25s %-20s %dK\n", i+1, filnam, mod.Format(time.RFC1123), size/1000)
		}
	}
	fmt.Printf("***************************\n")
}

func PrintFileTree (fileTree *nestTree) {

	fmt.Println("****************** fileTree *****************")
	maxnest := len(fileTree.nest)
	fmt.Printf("nest levels: %d\n", maxnest)

	for inest:=0; inest < maxnest; inest++ {
		numDirs := len((*fileTree).nest[inest].dirlist)
		fmt.Printf("nest level: %d dirs: %d\n", inest, numDirs)
		if numDirs == 0 {break}
		for i:=0; i< numDirs; i++ {
			dirObj := (*fileTree).nest[inest].dirlist[i]
			fmt.Printf("== dir[%d]: ", i)
			if len(dirObj.Name) == 0 {
				fmt.Printf("no dirObj\n")
			} else {
				numSubDir := len(dirObj.SubDirList)
				numFiles := len(dirObj.FilList)
				fmt.Printf("name: %s, sub dirs: %d, files: %d ===\n", dirObj.Name,numSubDir,numFiles) 
				PrintDirObj(dirObj)
			}
		}
	}

	fmt.Println("**************** end fileTree ***************")

}
