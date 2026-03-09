package main

import (
	"fmt"
	"net/http"
)

func restaurantNew(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	desc := r.FormValue("description")
	cuisines, err := validateCuisines(r.FormValue("cuisines"))
	locations, err := validateLocations(r.FormValue("locations"))
	if err != nil {
		return http.Error(w, err.Error(), 422)
	}
	if name == "" {
		return http.Error(w, "name can't be empty", 422)
	}
	if desc == "" {
		return http.Error(w, "description can't be empty", 422)
	}
}

func writeOp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	wo := r.FormValue("writeOp")
	writeOp, err := strconv.Atoi(wo)
	if err != nil {
		http.Error(w, "Bad writeOP", http.UnprocessableEntity)
		return
	}
	// This switch statement is our router since it is a highly-optimized basic language feature using
	// integer-comparison. There's no need for a bunch of complicated URL-parsing like Rails does.
	switch writeOp {
	case 1: // Restaurant new
		restaurantNew(w, r)
	default:
		http.Error(w, "Unknown writeOP", http.UnprocessableEntity)
		return
	}
	// TODO
	// Unix style. If errmsg, then return errmsg. Else return 204 silent no-content.
	// Return HTTP-204. Yes this is synchronous. But switching this to async is fairly simple and
	// all in one place. This can be done by setting up SSE (don't even need websockets for notifications
	// to the client).
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
