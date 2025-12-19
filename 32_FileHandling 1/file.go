package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){

	/*
	//OPENING 

	f, err := os.Open("./example.txt")
	if err != nil{
		//log the error
		panic(err)
	}
	
	fileInfo, err := f.Stat() //return fileInfo and err 
	if err != nil{
		//error
		panic(err)
	}

	fmt.Println("File name is : ", fileInfo.Name())
	fmt.Println("Is this a file or folder : ", fileInfo.IsDir())
	fmt.Println("File size : ", fileInfo.Size())
	fmt.Println("Mode of file : " , fileInfo.Mode())
	fmt.Println("Last Modified at : ", fileInfo.ModTime())

	*/


	f, err := os.Open("./example.txt")
	if err != nil{
		//error log
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 15)

	d, err := f.Read(buf)
	if err != nil{
		panic(err)
	}

	// Only print the bytes that were actually read.
	fmt.Println("data : ", string(buf), "Size is : ", d)


	// Reading the named file and returns the content 
	f1, err := os.ReadFile("./example.txt") // files ka content ek baar me hi load kr deti haii to ye dikkat hai => use stream things 
	if err != nil{
		//error log
		panic(err)
	}
	fmt.Println(string(f1))


	// read folders
	dir, err := os.Open("../")
	if err !=nil{
		panic(err)
	}
	defer dir.Close()

	dirInfo , err := dir.ReadDir(-1) // it will return slice 1, 2 -1

	for _, fn := range dirInfo{ // fn -> foldername
		fmt.Println(fn.Name(), fn.IsDir()) // agar koi file hui to wo false return krdega but folders honge to isDir use true return krega
	}


	// creating a file 
	nf, err := os.Create("example2.txt")
	if err != nil{
		panic(err)
	}

	defer f.Close()

	f.WriteString("hi go , ")
	f.WriteString("i first started thinking that you are good language :) ")
	fmt.Println("File Created :" ,nf)

	 // again reading 
	fi, err := os.ReadFile("./example2.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(fi))


	// things if we want to replace

	data, err := os.ReadFile("./example2.txt") //ioutil is deprecated
	if err != nil{
		log.Fatal(err)
	}

	//now replacing the text
	content := string(data)
	newData := strings.ReplaceAll(content, "fine" , "good")
	fmt.Println("Changes has been done : ", newData)

	


}