package main

import (
	"fmt"
	"net/http"
)

func writeOp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO
	// Unix style. If errmsg, then return errmsg. Else return 204 silent no-content.
	// Get the writeOp opcode (int)
	// Pass the r to the function to handle the write operation
	// Return HTTP-204. Yes this is synchronous. But switching this to async should be fairly simple and
	// all in one place.
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	fmt.Println("Started program")

	http.HandleFunc("/writeOp", writeOp) // Handle all write operations
	fmt.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
