package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Shanghai")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

func main() {
	
	_ = flag.Bool("x", false, "fixed the rule")
	tags := flag.String("t", "", "tags")
	file := flag.String("f", "", "file")
	
	flag.Parse()
	
	var tagsSlice = strings.Split(fmt.Sprintf("%s", *tags), ",")
	var filePath = fmt.Sprintf("%s", *file)
	
	GoFmt(filePath)
	
	var fileLine, fileByte = ReadFileAsSlice(filePath)
	
	var changeObjectSlice = Search(fileLine, fileByte)
	
	Change(filePath, changeObjectSlice, fileLine, fileByte, tagsSlice)
	
	GoFmt(filePath)
	
	println("tags rebuild success")
}