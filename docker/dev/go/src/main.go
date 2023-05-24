package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getFileSize() int64 {
	filePath := "/root/docker.img"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the file information
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Get the file size in bytes
	fileSize := fileInfo.Size()
	return fileSize
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)

	size := getFileSize()
	fmt.Fprintf(w, "Size: %d", size)
	fmt.Fprintf(w, `
	
	#     # ####### #        #####  ####### #     # ####### 
	#  #  # #       #       #     # #     # ##   ## #       
	#  #  # #       #       #       #     # # # # # #       
	#  #  # #####   #       #       #     # #  #  # #####   
	#  #  # #       #       #       #     # #     # #       
	#  #  # #       #       #     # #     # #     # #       
	 ## ##  ####### #######  #####  ####### #     # #######

`)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
