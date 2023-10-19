package permLib



import (
	"testing"
)


func TestPrint(t *testing.T) {

	var perm PermObj

	perm = 5

	permobj := &perm

	permobj.PrintPerm()
}

func TestSet(t *testing.T) {

	var perm PermObj
	perm = 0
	permobj := &perm

	err := permobj.SetPerm(0, true)
	if err != nil {t.Errorf("SetPerm: %v\n", err)}
	if perm != 1 {t.Errorf("perm should be 1!\n")}
}

func TestGet(t *testing.T) {

	var perm PermObj
	perm = 4
	permobj := &perm

	val, err := permobj.GetPerm(2)
	if err != nil {t.Errorf("GetPerm: %v\n", err)}

	if !val {t.Errorf("vak should be true!\n")}

}

func TestList(t *testing.T) {

	var perm PermObj
	perm = 6
	permobj := &perm
	res := permobj.ListPerm()
//	if err != nil {t.Errorf("ListPerm: %v\n", err)}

	for i:=0; i< 8; i++ {
		switch i {
		case 1:
			if !res[i] {t.Errorf("pos 1 should true!")}
		case 2:
			if !res[i] {t.Errorf("pos 2 should true!")}
		default:
			if res[i] {t.Errorf("pos %d should false!", i)}
		}
	}
}
