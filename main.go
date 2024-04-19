package main

import (
	"fmt"

	"github.com/kcharymyrat/my_go_dsa_lib/data_structures"
	"github.com/kcharymyrat/my_go_dsa_lib/ds_interfaces"
)

func main() {
	var arrInt data_structures.Array[int]
	var arrStr data_structures.Array[string]

	arrInt.Build(1, 2, 3, 4, 5)
	arrStr.Build("a", "b", "c", "d", "e")

	fmt.Println(arrInt, arrStr)
	arrInt.SetAt(1, 10)
	arrStr.SetAt(1, "yo")
	fmt.Println(arrInt.Len(), arrInt.GetAt(1))
	fmt.Println(arrStr.Len(), arrStr.GetAt(1))

	// Check if arrInt implements StaticSequenceInterface
	if intSeq, ok := interface{}(&arrInt).(ds_interfaces.StaticSequenceInterface[int]); ok {
		fmt.Println("arrInt implements StaticSequenceInterface[int]:", intSeq)
	} else {
		fmt.Println("arrInt does not implement StaticSequenceInterface[int]")
	}

	// Check if arrStr implements StaticSequenceInterface
	if strSeq, ok := interface{}(&arrStr).(ds_interfaces.StaticSequenceInterface[string]); ok {
		fmt.Println("arrStr implements StaticSequenceInterface[string]:", strSeq)
	} else {
		fmt.Println("arrStr does not implement StaticSequenceInterface[string]")
	}

}
