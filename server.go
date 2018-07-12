package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Content string `json:"content"`
}

type Articles []Article

func main(){
fmt.Println("Heloo")
	HandleRequest()
}
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page hit. ")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request){
	 	articles:=Articles{{Title:"About Fitness", Description:"Fitness is mainly requires for sedentary lifestyle.", Content:"dfsghfdgsfdgsf   hellllloooo world!!!"},
			{Title:"About Calm", Description:"State of mindis really important while riding any kind of vehicles.", Content:"Chapters about states of mind."}}
fmt.Println("EndPoint hit returning articles...")
	json.NewEncoder(w).Encode(articles)

}

func HandleRequest(){
	Router:=mux.
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", ReturnAllArticles)
	log.Fatal(http.ListenAndServe(":8082", nil))
}