// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// [END import]
// [START main_func]

func main() {
	// http.HandleFunc("/", indexHandler)
	http.HandleFunc("/", indexHandler2)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// // indexHandler responds to requests with our greeting.
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprint(w, "Hello, World!")
// }

func indexHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.URL.Path == "/api" { //なぜか、こいつだけ動かない。直下はダメ？
		fmt.Fprint(w, "api")
	} else if r.URL.Path == "/api/hoge" { //URL上での最後のスラッシュ、goでの最後のスラッシュ、vueでの最後のスラッシュの有無で動いたり動かなかったりする。
		fmt.Fprint(w, "api/hogeeeeeeeeeeeee")
	} else if r.URL.Path == "/api/fuga" {
		fmt.Fprint(w, "api/fuga")
	}
}
