package main

import (
	"archive/zip"
	"fmt"
)

func ListContentsOfZipFile(nameOfZipFile string) error {
	zipReader,err:=zip.OpenReader(nameOfZipFile)
	if err!=nil{
		return err
	}
	defer zipReader.Close()

	for _,file:= range zipReader.File {
		fmt.Println(file.Name)
	}
	return nil
}