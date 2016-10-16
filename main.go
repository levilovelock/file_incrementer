package main

import (
	"fmt"
    "io/ioutil"
  	"net/http"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func incrementer(w http.ResponseWriter, r *http.Request) {
	file := "C:/Users/levi/stream/mycount.txt"
	result := "0"

	fileData, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		fileData = nil
	}

	stringRep := string(fileData)
	fmt.Println("File contains: ",stringRep)

	if stringRep == "0" || stringRep == "" {
		result = "1"
	} else {
		i, ierr := strconv.Atoi(stringRep)
		if ierr != nil {
			result = "1"
		} else {
			result = strconv.Itoa(i+1)
		}
	}

    d1 := []byte(result)
    writeErr := ioutil.WriteFile("C:/Users/levi/stream/mycount.txt", d1, 0644)
    check(writeErr)
}

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  http.HandleFunc("/inc", incrementer)

  fmt.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}