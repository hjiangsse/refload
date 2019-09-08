// Copyright 2019 The hjiangsse. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents NGST reference data structure using native Go struct types
package refload

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// convert a string to another type
func convString(s string, t reflect.Type) (interface{}, error) {
	//create a new element, which type is v
	velm := reflect.New(t).Elem()

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:

		//if s is an empty string, then set it to "0"
		if len(s) == 0 {
			s = "0"
		}

		if i, err := strconv.ParseInt(s, 10, 64); err != nil {
			return nil, err
		} else {
			velm.SetInt(i)
		}
		return velm.Interface(), nil
	case reflect.Uint, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:

		//if s is an empty string, then set it to "0"
		if len(s) == 0 {
			s = "0"
		}

		if i, err := strconv.ParseUint(s, 10, 64); err != nil {
			return nil, err
		} else {
			velm.SetUint(i)
		}
		return velm.Interface(), nil
	//process load one character per field
	case reflect.Uint8:
		//if s is an empty string, then set it to "0"
		if len(s) == 0 {
			s = "\000"
		}

		sli := []uint8(s)
		velm.SetUint(uint64(sli[0]))
		return velm.Interface(), nil
	case reflect.String:
		velm.SetString(s)
		return velm.Interface(), nil
	default:
		return nil, fmt.Errorf("%s can not convert to type %v", s, t.String())
	}
}

// load reference data from file(have some varible fields in tail of the line)
// parameters:
// path -- the pathname of reference file
// sep  -- sperator character of reference file line
// ctrlFieldInx -- the control field index of reference file line
// s -- the destination collector of all reference line
func LoadRefDatVarTail(path, sep string, ctrlFiledInx int, s interface{}) (int, error) {
	var v reflect.Value
	var err error
	if v = reflect.ValueOf(s); v.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("the type of dest container type: %v, not a pointer", v.Type())
	}

	if v.Elem().Kind() != reflect.Slice {
		return 0, fmt.Errorf("dest container not a pointer to slice, but %v", v.Elem().Kind())
	}

	e := reflect.TypeOf(v.Elem().Interface()).Elem()
	elem := reflect.New(e).Elem() //elem is a line element(the struct is defined by user)

	//read reference file line by line
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return 0, err
	}
	defer file.Close()

	lineNum, err := lineCounter(file)
	if err != nil {
		return 0, err
	}

	//make a slice, the size of this slice should be lineNum
	v.Elem().Set(reflect.MakeSlice(v.Elem().Type(), lineNum, lineNum))

	file.Seek(0, 0)
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trimedLine := strings.TrimSpace(scanner.Text())
		if len(trimedLine) == 0 && index < lineNum-1 {
			//current line empty and is not the last line
			continue
		}

		lineSegs := strings.Split(trimedLine, sep)

		//parse normal fields(before the control field)
		for i := 0; i < ctrlFiledInx; i++ {
			v, err := convString(strings.TrimSpace(lineSegs[i]), elem.Field(i).Type())
			if err != nil {
				return index + 1, err
			}
			elem.Field(i).Set(reflect.ValueOf(v))
		}

		tailElemFieldNum, _ := strconv.Atoi(lineSegs[ctrlFiledInx])          //the number of fields in tail element struct
		tailElemNum := (len(lineSegs) - ctrlFiledInx - 1) / tailElemFieldNum //tail elem slice size(tail fields number in a line)

		tailLineSegs := lineSegs[ctrlFiledInx+1:]                                                                  //get taling line segments
		elem.Field(ctrlFiledInx).Set(reflect.MakeSlice(elem.Field(ctrlFiledInx).Type(), tailElemNum, tailElemNum)) //allocate tailing space

		//parse tailing fields(after the control field)
		for i := 0; i < tailElemNum; i++ {
			for j := 0; j < tailElemFieldNum; j++ {
				v, err := convString(strings.TrimSpace(tailLineSegs[i*tailElemFieldNum+j]), elem.Field(ctrlFiledInx).Index(i).Field(j).Type())
				if err != nil {
					return index + 1, err
				}
				elem.Field(ctrlFiledInx).Index(i).Field(j).Set(reflect.ValueOf(v))
			}
		}

		//append the new elem to dest collector slice
		v.Elem().Index(index).Set(elem)
		index++
	}

	return index, nil
}

// load reference data from file
func LoadRefDat(path string, sep string, s interface{}) (int, error) {
	index := 0

	//check if s is a pointer?
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("non-pointer %v\n", v.Type())
	}

	//check if s point to a slice
	ve := v.Elem()
	if ve.Kind() != reflect.Slice {
		return 0, fmt.Errorf("pointed value non-slice %v\n", ve.Type())
	}

	//get the slice element type, and create a new element
	e := reflect.TypeOf(ve.Interface()).Elem()
	elem := reflect.New(e).Elem() //now elem can be setted
	numFields := elem.NumField()

	//read the file line by line
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	linenum, err := lineCounter(file)
	if err != nil {
		return 0, err
	}

	//make slice, the size is linenum
	ve.Set(reflect.MakeSlice(ve.Type(), linenum, linenum))
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//split current line with separator
		trimline := strings.TrimSpace(scanner.Text())
		//if current line is empty and it is not the last line
		if len(trimline) == 0 && index < linenum-1 {
			continue
		}

		segs := strings.Split(trimline, sep)
		//line and slice type not match, and this line is not an empty line
		if len(segs) != numFields {
			return index + 1, fmt.Errorf("[line:%d] line seg num %d != struct field num %d", index+1, len(segs), numFields)
		}

		for i := 0; i < numFields; i++ {
			v, err := convString(strings.TrimSpace(segs[i]), elem.Field(i).Type())
			if err != nil {
				return index + 1, err
			}
			elem.Field(i).Set(reflect.ValueOf(v))
		}
		//append new elem to the slice
		ve.Index(index).Set(elem)
		index++
	}
	return index, nil
}

// get the line count of a file
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
