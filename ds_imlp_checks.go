package main

import (
	"fmt"

	"github.com/kcharymyrat/my_go_dsa_lib/data_structures"
	"github.com/kcharymyrat/my_go_dsa_lib/ds_interfaces"
)

func doubleDynamicArrayDSImplementation() {
	var doubleDynArrFloat data_structures.DoubleDynamicArray[float64]
	var doubleDynArrStr data_structures.DoubleDynamicArray[string]
	fmt.Println(&doubleDynArrFloat)
	fmt.Println(&doubleDynArrStr)

	(&doubleDynArrFloat).Build(6, 10, 20, 30, 40, 50, 60, 70, 80)
	(&doubleDynArrStr).Build("a", "b", "c", "d", "e")

	fmt.Println(&doubleDynArrFloat, doubleDynArrFloat.GetData())
	fmt.Println(&doubleDynArrStr, doubleDynArrStr.GetData())

	if dynFloatSeq, ok := interface{}(&doubleDynArrFloat).(ds_interfaces.DynamicSequenceInterface[float64]); ok {
		fmt.Println("doubleDynArrFloat implements DynamicSequenceInterface[float64]:", dynFloatSeq)
	} else {
		fmt.Println("doubleDynArrFloat does NOT implement DynamicSequenceInterface[float64]")
	}

	if dynStrSeq, ok := interface{}(&doubleDynArrStr).(ds_interfaces.DynamicSequenceInterface[string]); ok {
		fmt.Println("doubleDynArrStr implements DynamicSequenceInterface[string]:", dynStrSeq)
	} else {
		fmt.Println("doubleDynArrStr does NOT implement DynamicSequenceInterface[string]")
	}

	doubleDynArr := doubleDynArrFloat.GetData()
	init7Ptr := &doubleDynArr[7]

	fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(doubleDynArr), init7Ptr)

	doubleDynArrFloat.SetAt(2, 77)
	condition := init7Ptr == &doubleDynArr[7]
	fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(doubleDynArr), init7Ptr, condition)
	againDoubleDynArr := doubleDynArrFloat.GetData()

	for i := 0; i < 25; i++ {
		doubleDynArrFloat.InsertLast(float64(i + 1))
		againDoubleDynArr = doubleDynArrFloat.GetData()
		condition = init7Ptr == &againDoubleDynArr[7]
		fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)
	}

	doubleDynArr = doubleDynArrFloat.GetData()
	init7Ptr = &doubleDynArr[7]
	for i := 0; i < 25; i++ {
		doubleDynArrFloat.DeleteLast()
		againDoubleDynArr = doubleDynArrFloat.GetData()
		condition = init7Ptr == &againDoubleDynArr[7]
		fmt.Println(doubleDynArrFloat.String(), doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)
	}

	doubleDynArrFloat.InsertFirst(-1)
	againDoubleDynArr = doubleDynArrFloat.GetData()
	condition = init7Ptr == &againDoubleDynArr[7]
	fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	doubleDynArrFloat.DeleteFirst()
	againDoubleDynArr = doubleDynArrFloat.GetData()
	condition = init7Ptr == &againDoubleDynArr[7]
	fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	doubleDynArrFloat.InsertAt(1, 11)
	againDoubleDynArr = doubleDynArrFloat.GetData()
	condition = init7Ptr == &againDoubleDynArr[7]
	fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	doubleDynArrFloat.DeleteAt(1)
	againDoubleDynArr = doubleDynArrFloat.GetData()
	condition = init7Ptr == &againDoubleDynArr[7]
	fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	iterator := doubleDynArrFloat.Iterator()
	for iterator.HasNext() {
		next, _ := iterator.Next()
		fmt.Println(next)
	}

}

func dynamicArrayDSImplementation() {
	var doubleDynArrFloat data_structures.DynamicArray[float64]
	var doubleDynArrStr data_structures.DynamicArray[string]

	doubleDynArrFloat.Build(6, 10, 20, 30, 40, 50, 60, 70, 80)
	doubleDynArrStr.Build("a", "b", "c", "d", "e")

	if dynFloatSeq, ok := interface{}(&doubleDynArrFloat).(ds_interfaces.DynamicSequenceInterface[float64]); ok {
		fmt.Println("doubleDynArrFloat implements DynamicSequenceInterface[float64]:", dynFloatSeq)
	} else {
		fmt.Println("doubleDynArrFloat does NOT implement DynamicSequenceInterface[float64]")
	}

	if dynStrSeq, ok := interface{}(&doubleDynArrStr).(ds_interfaces.DynamicSequenceInterface[string]); ok {
		fmt.Println("doubleDynArrStr implements DynamicSequenceInterface[string]:", dynStrSeq)
	} else {
		fmt.Println("doubleDynArrStr does NOT implement DynamicSequenceInterface[string]")
	}

	doubleDynArr := doubleDynArrFloat.GetData()
	init7Ptr := &doubleDynArr[7]

	fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(doubleDynArr), init7Ptr)

	doubleDynArrFloat.SetAt(2, 77)
	condition := init7Ptr == &doubleDynArr[7]
	fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(doubleDynArr), init7Ptr, condition)
	againDoubleDynArr := doubleDynArrFloat.GetData()

	for i := 0; i < 25; i++ {
		doubleDynArrFloat.InsertLast(float64(i + 1))
		againDoubleDynArr = doubleDynArrFloat.GetData()
		condition = init7Ptr == &againDoubleDynArr[7]
		fmt.Println(doubleDynArrFloat, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)
	}

	doubleDynArr = doubleDynArrFloat.GetData()
	init7Ptr = &doubleDynArr[7]
	for i := 0; i < 25; i++ {
		doubleDynArrFloat.DeleteLast()
		againDoubleDynArr = doubleDynArrFloat.GetData()
		condition = init7Ptr == &againDoubleDynArr[7]
		fmt.Println(doubleDynArrFloat.String(), doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)
	}

	// doubleDynArrFloat.InsertFirst(-1)
	// againDoubleDynArr = doubleDynArrFloat.GetData()
	// condition = init7Ptr == &againDoubleDynArr[7]
	// fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	// doubleDynArrFloat.DeleteFirst()
	// againDoubleDynArr = doubleDynArrFloat.GetData()
	// condition = init7Ptr == &againDoubleDynArr[7]
	// fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	// doubleDynArrFloat.InsertAt(1, 11)
	// againDoubleDynArr = doubleDynArrFloat.GetData()
	// condition = init7Ptr == &againDoubleDynArr[7]
	// fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

	// doubleDynArrFloat.DeleteAt(1)
	// againDoubleDynArr = doubleDynArrFloat.GetData()
	// condition = init7Ptr == &againDoubleDynArr[7]
	// fmt.Println(againDoubleDynArr, doubleDynArrFloat.Len(), cap(againDoubleDynArr), init7Ptr, condition)

}

func arrayDSImplementation() {
	var arrInt data_structures.Array[int]
	var arrStr data_structures.Array[string]

	arrInt.Build(1, 2, 3, 4, 5)
	arrStr.Build("a", "b", "c", "d", "e")

	fmt.Println(arrInt, arrStr)
	arrInt.SetAt(1, 10)
	arrStr.SetAt(1, "yo")
	intVal, errInt := arrInt.GetAt(10)
	strVal, errStr := arrStr.GetAt(-1)
	fmt.Println(arrInt.Len(), intVal, errInt)
	fmt.Println(arrStr.Len(), strVal, errStr)

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
