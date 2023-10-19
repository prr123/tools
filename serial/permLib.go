// permLib a library to set and retrieve boolean permissions
// Author: prr, azul software
// Date: 19 Oct 2023
// copyright (c) 2023 prr, azul software
//
package permLib

import (
//	"log"
	"fmt"
)


type PermObj byte

func (perm *PermObj) SetPerm(pos int, val bool) (err error) {

	var mask PermObj = 0
	if pos> 8|| pos < 0 {return fmt.Errorf("invalid pos!")}
	if val {mask = 1}

	mask = mask << pos
//	fmt.Printf("mask: %08b\n", mask)

	*perm = (*perm)^mask
//	fmt.Printf("perm: %08b\n", *perm)
	return nil
}

func (perm *PermObj) GetPerm(pos int) (val bool, err error) {

	var mask PermObj = 0
	if pos> 8|| pos < 0 {return false, fmt.Errorf("invalid pos!")}

	val = false

	mask = mask << pos
//	fmt.Printf("mask: %08b\n", mask)
	if (*perm)^mask > 0 {val = true}
//	fmt.Printf("perm: %08b\n", *perm)

	return val, nil
}


func (perm *PermObj) ListPerm () (res [8]bool) {

	var mask PermObj = 1
	for i:=0; i< 8; i++ {
		res[i] = false
		if *perm & mask > 0 {res[i] = true}
		mask = mask << 1
	}
	return res
}

func (perm *PermObj) PrintPerm() {
	fmt.Println("******* Perm ******")
//	fmt.Printf("perm: %08b\n", *perm)
	var mask PermObj = 1
	for i:=0; i< 8; i++ {
		res := false
//	fmt.Printf("mask: %08b ", mask)
//	fmt.Printf("perm: %08b\n", *perm)
		if *perm & mask >0 {res= true}
		mask = mask << 1
		fmt.Printf(" %d: %t\n", i, res)
	}
	fmt.Println("***** End Perm ****")

}
