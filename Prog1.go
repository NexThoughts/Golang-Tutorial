//package main

//import (
//	"net/http"
//	"fmt"
//	"log"
//)


//
//import (
//	"fmt"
//	"database/sql"
//	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
//)
//
//func init() {
//	//sql.Register("postgres", new(PostgresDriver))
//}
//
//func main() {
//	var a string="kkkkkkkk"
//	fmt.Println(a)
//	sql.Open("postgres", "mydb")
//
//
//}


package main

import (
"fmt"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print(http.ListenAndServe(":8081", nil))
}