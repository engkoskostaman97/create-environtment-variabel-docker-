package main

import "fmt"
import "net/http"
import "os"
import "strings"

func getallenv( w http.ResponseWriter,  r *http.Request){
	for _, element:= range os.Environ(){
		variabel := strings.Split(element,"=")
		fmt.Fprintln(w, variabel[0],"=>",variabel[1])
	}
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, os.Getenv("DB_PASSWORD"))
	})

	http.HandleFunc("/getallenv",getallenv)
	http.ListenAndServe(":8080",nil)
}