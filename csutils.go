package csutils

import (
   // "bufio"	//<-- implements buffered I/O.
   "fmt"	//<-- implements formatted I/O with functions analogous to C's printf and scanf.
   // "os"		//<-- provides a platform-independent interface to operating system functionality.
   "log"
)

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func ReadInteger(message string) int {
	//Prints the message
	fmt.Print(message)
	var number int

	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func ReadName(message string) string {
	//Prints the message
	fmt.Print(message)
	var name string

	_, err := fmt.Scanln(&name)

	if err != nil {
		log.Fatal(err)
	}

	return name
}