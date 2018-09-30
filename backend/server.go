package main

import (
	"encoding/json"
	"log"
	"net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
	"strconv"
	"time"
)
type Person struct {
	Name     string
	AgeYears int
}
type list struct {
	title	 string
}
type Blog struct {
	Title	 string
	Content	 string
}
type Jsonx struct {
	Name	 int `json:"Name"`
	Description	 string `json:"Description"`
}
func paramHandler(w http.ResponseWriter, r *http.Request) {

}
func listProducts(w http.ResponseWriter, r *http.Request) {
	// list all products
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var s []string
    s = list1()
    w.Header().Set("Content-Type", "application/json")
    enc := json.NewEncoder(w)
	err := enc.Encode(s)
	if err != nil {
		// if encoding fails, create an error page with code 500.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func insert (title string, content string, date string) {
    database, _ := sql.Open("sqlite3", "./blog.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, date TEXT, title TEXT, content TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO people (date, title, content) VALUES (?, ?, ?)")
    statement.Exec(date, title, content)
    defer database.Close()
}
func list1() []string{
    database, _ := sql.Open("sqlite3", "./blog.db")
    // statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, date TEXT, title TEXT, content TEXT)")
	var id int
    var title string
    rows, _ := database.Query("SELECT id, title FROM people")
    
    // var lastname string
    i :=0
    for rows.Next() {
    	i=i+1
    }
    s := make([]string,i)
    i = 0
    rows, _ = database.Query("SELECT id, title FROM people")
    for rows.Next() {
        rows.Scan(&id, &title)
        s[i] = title
        i=i+1
        // fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
    }
    // defer database.Close()
    return s
}
func addProduct(w http.ResponseWriter, r *http.Request) {
	// add a product
	//w.header.send
	log.Printf( "Hello,in add product")
	// w.WriteHeader(310)
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" {
		title = "friend"
	}
	if content == "" {
		return
	}
	current_time := time.Now().Local()
	date := current_time.Format("2006-01-02")
	log.Printf( "Hello, %s!", title)
	insert(title,content,date)
	p := Person{"gopher", 5}
	w.Header().Set("Content-Type", "application/json")
	// encode p to the output.
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}
func checkErr(err error) {
	if err !=nil {
		log.Printf("error occurd,%s",err)
	}
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["blogid"]
	// str1 := fmt.Sprintf("<h1>fetching product with ID %q</h1>", id)
    database, _ := sql.Open("sqlite3", "./blog.db")
    rows1, err1 := database.Prepare("SELECT title,content FROM people WHERE id=?")
    checkErr(err1)
    // rows, err := rows1.Exec(id)
    var title string
    var content string
    time, _ := strconv.Atoi(id)
	err := rows1.QueryRow(time+1).Scan(&title,&content)
    checkErr(err)
    // id1 := strconv.Atoi(id)
    time, err = strconv.Atoi(id)
	log.Printf( "id, %d", time)
    
    // err=rows.Scan(&title, &content)
    checkErr(err)
	log.Printf( "title,content, %s,%s!", title,content)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	p := Blog {title,content}
	enc := json.NewEncoder(w)
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.FPrintln(w,str1)
	// fmt.Fprintln(w, str1)
	// get a specific product
}
func addlike(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf( "entered")
	var p Jsonx
	dec := json.NewDecoder(r.Body)
	log.Printf( "entered144")
	err := dec.Decode(&p)
	id :=p.Name
	time := id
	// time, _ := strconv.Atoi(id)
	// log.Printf( "id, %d", time)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf( "entered150")
	idx :=p.Description
	// id := mux.Vars(r)["name"]
    // timex, _:= strconv.Atoi(idx)
	log.Printf( "idx, %d", idx)
	database, _ := sql.Open("sqlite3", "./blog.db")
    rows1, _ := database.Prepare("UPDATE people SET like = like + 1 WHERE id = ?")
	_,err3 := rows1.Exec(time+1)
    checkErr(err3)	

}
func main() {
	r := mux.NewRouter()
	// match only GET requests on /product/
	r.HandleFunc("/", listProducts).Methods("GET")
	r.HandleFunc("/like", addlike).Methods("POST")

	// match only POST requests on /product/
	r.HandleFunc("/blogger/", addProduct).Methods("POST")

	// match GET regardless of productID
	r.HandleFunc("/blog/{blogid}", getProduct).Methods("GET")

	// handle all requests with the Gorilla router.
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}