package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errHandler(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {

		}
	}
}

func handleFile(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		return err1
	}
	_, err2 := writer.Write(all)
	if err2 != nil {
		return err2
	}
	return nil
}

func main() {
	http.HandleFunc("/list/", errHandler(handleFile))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
