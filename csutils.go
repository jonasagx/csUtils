package csutils

import (
   "bufio"	//<-- implements buffered I/O.
   "fmt"	//<-- implements formatted I/O with functions analogous to C's printf and scanf.
   "os"		//<-- provides a platform-independent interface to operating system functionality.
   "log"
   // "io"
)

func ReadFileContent(filePath string) string {
	f, err := os.Open(filePath)
	check(err)

	fInfo, _ := f.Stat()
	if int(fInfo.Size()) > 4000000 {
		log.Fatal("File bigger than 4000000")
	}

	fileBuffer := bufio.NewReader(f)
	fileContent, err := fileBuffer.Peek(int(fInfo.Size()))
	check(err)
	f.Close()
	return string(fileContent)
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func ReadInteger(message string) int {
	//Prints the message
	fmt.Print(message)
	var number int

	_, err := fmt.Scanln(&number)

	check(err)
	return number
}

func ReadText(message string) {
	//Prints the message
	fmt.Print(message)
	var name string

	_, err := fmt.Scanln(&name)

	check(err)

	return name
}

func ReadName(message string) string {
	return ReadText(message)
}