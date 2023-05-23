package main

import "fmt"
import "net/http"

func welcome (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "welcome to ENGKOS")
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Selamat Datang Bro")
	})

	http.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":8080",nil)
}