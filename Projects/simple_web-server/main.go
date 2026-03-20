package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"parseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"address = %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"mehtod is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("starting server at port 80\n")
	if err := http.ListenAndServe(":80",nil); err !=nil{
		log.Fatal(err)
	}

}