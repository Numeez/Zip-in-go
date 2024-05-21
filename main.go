package main

import "log"

func main(){
	err:=ZipFile("sample.txt")
	if err !=nil{
		log.Fatal(err)
	}

	err=CreateZipFolder("sample","sample.txt.gz")
	if err !=nil{
		log.Fatal(err)
	}
	
}