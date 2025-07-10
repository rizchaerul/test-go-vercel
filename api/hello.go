package handler

import (
	"fmt"
	"net/http"
	"time"
)

// func main() {
// 	fmt.Println("Hello, World!")
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(15 * time.Second)
	fmt.Fprintf(w, "<h1>Hello from Go! m luthfi jasir</h1>")
}
