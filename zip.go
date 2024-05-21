package main

import (
	"compress/gzip"
	"io"
	"os"
)

func ZipFile(pathOfFile string) error{
	file,err:=os.Open(pathOfFile)
	if err!=nil{
		return err
	}
	defer file.Close()
	gzWriter,err:=os.Create(pathOfFile+".gz")
	if err !=nil{
		return err
	}
	defer  gzWriter.Close()

	zipWriter:=gzip.NewWriter(gzWriter)
	defer zipWriter.Close()

	_,err =io.Copy(zipWriter,file)
	if err !=nil{
		return err
	}
	zipWriter.Close()
	return nil

}