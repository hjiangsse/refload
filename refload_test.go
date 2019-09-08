//This is the test file to test refload
package refload

import (
	"fmt"
	"reflect"
	"testing"
)

//convString test1: normal test
func Test_convString_int(t *testing.T) {
	//[int] test
	var x int = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

//userdefined type test
type ElemType uint32

func Test_convString_userdefined_type(t *testing.T) {
	var x ElemType
	if v, err := convString("1000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [ElemType] to [int32] fail: %s", err)
	} else {
		t.Log("conv [ElemType] to [int32] ok! test pass!")
		fmt.Println(reflect.TypeOf(v))
	}
}

func Test_convString_int8(t *testing.T) {
	//[int] test
	var x int8 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_int16(t *testing.T) {
	//[int] test
	var x int16 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_int32(t *testing.T) {
	//[int] test
	var x int32 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_int64(t *testing.T) {
	//[int] test
	var x int64 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_uint(t *testing.T) {
	//[int] test
	var x uint = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("900000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_uint8(t *testing.T) {
	//[int] test
	var x uint8 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_uint16(t *testing.T) {
	//[int] test
	var x uint16 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_uint32(t *testing.T) {
	//[int] test
	var x uint32 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_uint64(t *testing.T) {
	//[int] test
	var x uint64 = 10
	if v, err := convString("100", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == 100 {
		t.Log("conv [string] to [int] ok! test pass!")
	}

	if v, err := convString("100000000000000000", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else {
		fmt.Println(v)
	}
}

func Test_convString_string(t *testing.T) {
	//[int] test
	var x string = "hjiang"
	if v, err := convString("heng", reflect.TypeOf(x)); err != nil {
		t.Errorf("test conv [string] to [int] fail: %s", err)
	} else if v == "heng" {
		t.Log("conv [string] to [int] ok! test pass!")
	}
}

//if a string can't be convert to unsigned int?
func Test_convString_convError(t *testing.T) {
	var x uint32 = 100
	if _, err := convString("hjiang", reflect.TypeOf(x)); err != nil {
		t.Logf("conv [string] %s to [uint32] fail: %s, OK!", "hjiang", err)
	}
}

//if convert a string to a invalid type
func TestConvStrInvType(t *testing.T) {
	var x float32 = 3.14
	str := "3.56"
	if _, err := convString(str, reflect.TypeOf(x)); err != nil {
		t.Logf("conv [string] %s to [uint32] fail: %s, OK!", str, err)
	}
}

func TestLoadRefDat_NoPtr(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice testElem
	if _, err := LoadRefDat("./testfile/testfile1.dat", "|", elemSlice); err != nil {
		t.Logf("valid struct, valid file, but: %s", err)
	} else {
		fmt.Println(elemSlice)
	}
}

func TestLoadRefDat_NoSlice(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice testElem
	if _, err := LoadRefDat("./testfile/testfile1.dat", "|", &elemSlice); err != nil {
		t.Logf("valid struct, valid file, but: %s", err)
	} else {
		fmt.Println(elemSlice)
	}
}

//field of struct is valid, line in file valid
func Test_LoadRefDat_PathInv(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice []testElem
	if _, err := LoadRefDat("./testfile/testfile1_inv.dat", "|", &elemSlice); err != nil {
		t.Logf("valid struct, valid file, but: %s", err)
	} else {
		fmt.Println(elemSlice)
	}
}

//field of struct is valid, line in file valid
func Test_LoadRefDat_1(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice []testElem
	if _, err := LoadRefDat("./testfile/testfile1.dat", "|", &elemSlice); err != nil {
		t.Errorf("valid struct, valid file, but: %s", err)
	} else {
		fmt.Println(elemSlice)
	}
}

//field of struct is valid, line in file invalid
//"0123T" conv to int err?
func Test_LoadRefDat_2(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice []testElem
	if _, err := LoadRefDat("./testfile/testfile2.dat", "|", &elemSlice); err != nil {
		str := fmt.Sprintf("valid struct, valid file, but: %s", err)
		t.Log(str)
	} else {
		fmt.Println(elemSlice)
	}
}

//field of struct is valid, line in file invalid
//line segs num less than struct field number
func Test_LoadRefDat_3(t *testing.T) {
	type testElem struct {
		Field1 int
		Field2 uint
		Field3 string
	}

	var elemSlice []testElem
	if _, err := LoadRefDat("./testfile/testfile3.dat", "|", &elemSlice); err != nil {
		str := fmt.Sprintf("valid struct, invalid file, %s", err)
		t.Log(str)
	} else {
		fmt.Println(elemSlice)
	}
}

//some fields of the struct is user defined type
/*
func Test_LoadRefDat_4(t *testing.T) {
	type NameT string
	type AgeT uint32

	type Person struct {
		Name NameT
		Age  AgeT
	}

	var elemSlice []Person
	if _, err := LoadRefDat("./testfile/testfile4.dat", "|", &elemSlice); err != nil {
		str := fmt.Sprintf("Struct have user defined type, %s\n", err)
		t.Errorf("%s", str)
	} else {
		fmt.Println(elemSlice)
	}
}
*/

//some of fields of the struct have space leading of tailing
func Test_LoadRefDat_Space(t *testing.T) {
	type testElem struct {
		F1 int
		F2 uint
		F3 string
	}

	var elemSlice []testElem
	if _, err := LoadRefDat("./testfile/testfile5.dat", "|", &elemSlice); err != nil {
		str := fmt.Sprintf("Test Error!")
		t.Errorf("%s", str)
	} else {
		fmt.Println(elemSlice)
	}
}

type TailStruct struct {
	F1 uint32
	F2 uint64
}

type FileLine struct {
	F1   string
	F2   string
	F3   string
	F4   string
	Tail []TailStruct
}

func Test_LoadRefDatVarTail(t *testing.T) {
	refFilePath := "./refdat/test.txt"
	destColl := make([]FileLine, 0)
	_, err := LoadRefDatVarTail(refFilePath, "|", 4, &destColl)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(destColl)
}
