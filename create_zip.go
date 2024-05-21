package main

import (
	"archive/zip"
	"io"
	"os"
)

func CreateZipFolder(nameOfFolder string,nameOfFile string) error{
	file,err:=os.Create(nameOfFolder+".zip")
	if err !=nil{
		return err
	}
	defer file.Close()
	zipWriter:=zip.NewWriter(file)
	 f,err:=os.Open(nameOfFile)
	 if err !=nil{
		 return err
	 }
	 f.Close()
	 w1,err:=zipWriter.Create("sample.txt")
	 if err !=nil{
		 return err
	 }
	 _,err = io.Copy(w1,f)
	 if err!=nil{
		return err
	}
	zipWriter.Close()
	return nil
}
