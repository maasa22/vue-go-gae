package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Todo ... This is a todo
type Todo struct {
	id       int
	todo     string
	priority string
}

// 初期化
var initialTodos []Todo

func main() {

	initialTodos = append(initialTodos,
		Todo{
			id:       1,
			todo:     "Consider Adam as a candidate",
			priority: "life changing",
		},
		Todo{
			id:       2,
			todo:     "Take out the garbage",
			priority: "life changing",
		},
		Todo{
			id:       3,
			todo:     "Paint the house",
			priority: "important",
		},
	)

	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //本来はallowするURLを指定とかすべきそう。
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// fmt.Fprint(w, "Hello, World!")

	if r.URL.Path == "/api" { //なぜか、こいつだけ動かない。直下はダメ？
		fmt.Fprint(w, "api")
	} else if r.URL.Path == "/api/hoge" { //URL上での最後のスラッシュ、goでの最後のスラッシュ、vueでの最後のスラッシュの有無で動いたり動かなかったりする。
		fmt.Fprint(w, "api/hogeee")
	} else if r.URL.Path == "/api/fuga" {
		fmt.Fprint(w, "api/fuga")
	} else if r.URL.Path == "/api/todos" {
		for i, s := range initialTodos {
			textTmp := ""
			textTmp = fmt.Sprintf("index: %d, name: %s\n", i, s.todo)
			fmt.Fprint(w, textTmp)
		}
		// fmt.Fprint(w, initialTodos[0])
		// fmt.Fprint(w, initialTodos[0].id)
	}

}
