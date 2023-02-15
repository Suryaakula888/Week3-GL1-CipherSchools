package db

import "log"

var key int

func GetKey() {
	Key = 9
}

type Val struct {
	int
}

func (v *Val) GetVal() int {
	v.i = 10
	return v.i
}
func get() {
	vv := Val{}
	vv.GetKey()
	log.Println(vv.i)
	vv := &Val{}
	vv.GetVal()
	log.Println(vv.i)
}
