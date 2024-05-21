package main

import "log"

func main(){

	// err:=CreateZipFolder("sample","sample.txt")
	// if err !=nil{
	// 	log.Fatal(err)
	// }

	err:=ListContentsOfZipFile("sample.zip")
	if err !=nil{
		log.Fatal(err)
	}
	
}