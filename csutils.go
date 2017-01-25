package csutils

import (
   "bufio"	//<-- implements buffered I/O.
   "fmt"	//<-- implements formatted I/O with functions analogous to C's printf and scanf.
   "os"		//<-- provides a platform-independent interface to operating system functionality.
)

func Print(str string) {
	fmt.Print(str)	
}

func ReadInput(message string) {
	//Prints the message
	fmt.Print("Write your name then press ENTER, please: ")

	//Reads the text input
	var scan = bufio.NewScanner(os.Stdin)
	
	//Check if anything was passed as input and extract it
	scan.Scan()
	//Put the input in a variable of type string
	var name string = scan.Text()
	return name
}