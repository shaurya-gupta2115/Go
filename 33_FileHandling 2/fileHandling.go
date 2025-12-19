package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){

	// // create a file
	// f, err := os.Create("./newfile.txt")

	// //handling errror
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// f.WriteString("HEY! MY NAME IS SHAURYA GUPTA")
	
	// reading
	// readF, err := os.ReadFile("./newfile.txt")
	// fmt.Println(string(readF))

	sourceFile, err := os.Open("./newfile.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("copiesNewFile.txt")
	if err!= nil{
		panic(err)
	}
	defer destFile.Close()

	// if we want to use streaming fashion => bufio

	reader := bufio.NewReader(sourceFile) // buffered wrapper around the source file
	writer := bufio.NewWriter(destFile)

	//infinite loop
	for{
		b, err := reader.ReadByte()
		if err != nil{
			if err == io.EOF {
				break
			}
			panic(err)
		}

		e := writer.WriteByte(b)
		if e != nil{
			panic(e)
		}
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("File ", sourceFile.Name(), " is copied to " , destFile.Name())
	


	// deleting the file
	// toBeDeletionFile, err := os.Open("./copiesNewFile.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer toBeDeletionFile.Close()

	// e := os.Remove(toBeDeletionFile.Name())
	// if e != nil{
	// 	panic(e)
	// }

	

}