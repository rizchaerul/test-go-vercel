package handler

import (
	"fmt"
	"net/http"
	"os/exec"
)

// func main() {
// 	fmt.Println("Hello, World!")
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	body := io.ReadAll(r.Body)
	cmd := exec.Command("./project1")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(w, "<h1>error"+err.Error()+"</h1>")
	}

	// time.Sleep(15 * time.Second)
	fmt.Fprintf(w, "<h1>Hello from Go!"+string(output)+body+"</h1>")
}
