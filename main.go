package main

import (
	"errors"
	"fmt"
	"log"
)

//ERROR EXAMPLE w/o type assertion:
//Rry to convert interface{} to certain type. ex string:

/** func convToString(data interface{}) string {
*** 	return string(data)
*** }
**/

func convToString(data interface{}) string {
	return string(data.(string))
}

func convToStringWithCheck(data interface{}) (string, error) {
	value, ok := data.(string)
	if !ok {
		err := errors.New("Step 2: Assertion fail: argument passed is not in string type")
		return "", err
	}
	return string(value), nil
}

func assertWithSwitch(data interface{}) {
	switch data.(type) {
	case string:
		fmt.Println("Step 3: Switch case in string type")
		// Do somthing with string type data
	case int:
		fmt.Println(("Step 3: Switch case in integer type"))
		// Do somthing with int type data
	default:
		fmt.Printf("Step 3: Switch case in unkown type")
	}
}

func main() {
	/*********************
	** Direct assertion **
	**********************/

	// Case 1: pass string type
	data1 := convToString("Step 1: Directly assert data to string type")
	fmt.Println(data1)

	// Case 2: pass int type
	// _ := convToString(123) // ERROR: panic: interface conversion: interface {} is int, not string

	/*************************
	** Assertion with check **
	*************************/

	// Case 1: pass string type
	data2, err := convToStringWithCheck("Step 2: Check if data is string type before assert it")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data2)

	// Case 2: pass int type
	_, err = convToStringWithCheck(123)
	if err != nil {
		fmt.Println(err)
	}

	/*************************
	** Assertion with switch **
	*************************/

	// Case 1: pass string type
	assertWithSwitch("step3")
	// Case 2: pass int type
	assertWithSwitch(123)
	// Case 3: pass error type
	assertWithSwitch(errors.New("An error"))
}
