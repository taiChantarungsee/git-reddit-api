package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/turnage/graw/reddit"
)

//Router
func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/reddit-comments", HomeHandler)
	http.Handle("/", router)
	return router
}

//Putting the code in the handler function? Will this be quick?
//Do this with Python APIs? Also make it do some processing...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Done!")
	bot, err := reddit.NewBotFromAgentFile("", 0)
	if err != nil{
		fmt.Println("Failed to start bot.", err)
		return
	}

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil{
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	} 
}

func main() {
	r := NewRouter()
	fmt.Print("Server starting!")
	log.Fatal(http.ListenAndServe(":8080", r))
}