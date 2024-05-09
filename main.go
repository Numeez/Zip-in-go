package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	
	
)

func main(){
	
	// file_in_num:=[]int{}
	word_count:=0
	dictionaryMap:= map[string]int{}
	f,err:= os.Stat("sample.txt")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(f.Size())
	file,err:= os.Open("sample.txt")
	
	if err != nil{
		log.Fatal(err)
	}
	
	Scanner:= bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	

	for Scanner.Scan() {
		word:= Scanner.Text()
		
		
		_,exists:= dictionaryMap[word]
		if !exists{
			dictionaryMap[word] = word_count
			word_count++
			
		}
	}
	
	err = Scanner.Err()
	if err !=nil{
		log.Fatal(err)
	}


	

	


}