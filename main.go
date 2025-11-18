package main

// #cgo CFLAGS: -g -Wall -I/home/alex/Repos/libnspire/src/api
// #cgo LDFLAGS: -L/usr/lib/libnspire.so -lnspire
// #include <stdlib.h>
// #include "upload.h"
import "C"
import (
	"fmt"
	"os"
	"path"
)

func main() {
	// check if theres enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: Please provide the file to send as an argument")
		return
	}

	// get args
	fsrc := C.CString(os.Args[1])
	fdest := C.CString(path.Base(os.Args[1]))
	// free them when the function ends cause good practice ig
	defer C.free(unsafe.Pointer(fsrc))
	defer C.free(unsafe.Pointer(fdest))

	// upload file (funciton declated in upload.c)
	ret := C.upload(fsrc, fdest)
	
	os.Exit(int(ret))
}
