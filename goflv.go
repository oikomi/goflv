package main

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
	"github.com/oikomi/goflv/flv"
)


/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

const char* build_time(void) {
	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

func BuildTime() string {
	return buildTime
}

const VERSION string = "0.10"

func version() {
	fmt.Printf("goflv version %s Copyright (c) 2014 Harold Miao (miaohonghit@gmail.com)  \n", VERSION)
}

func help() {
	fmt.Printf("goflv flvfilename \n");
}

func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}
	
	fs := flv.NewFlvFileSpec(os.Args[1])
	fh := flv.NewFlvFileHandle()
	
	err := fh.FlvOpen(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	err = fh.FlvFileStat(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	err = fh.FlvReadHeader(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	err = fh.FlvReadBody(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	//res, _ := json.Marshal(fs)
	json.Marshal(fs)
	//log.Println(string(res))

}