package csutils

import (
   "bufio"	//<-- implements buffered I/O.
   "fmt"	//<-- implements formatted I/O with functions analogous to C's printf and scanf.
   "os"		//<-- provides a platform-independent interface to operating system functionality.
)

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func ReadInteger(message string) int {
	fmt.Print(message)
	var n int
    fmt.Scanf("%d", &n)
	return n
}

func ReadName(message string) string {
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

func PrintFacts(str string) {
	fmt.Println("The name", str, "has", len(str), "characters")
}