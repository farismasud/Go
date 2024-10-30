import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/users", handleUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Get all users from the database
	users, err := getAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode the users into JSON and write them to the response
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Decode the user from the request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the user in the database
	if err := createUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the created user to the response
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the user from the request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the user in the database
	if err := updateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the updated user to the response
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request path
	userID := r.URL.Query().Get("id")

	// Delete the user from the database
	if err := deleteUser(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write a success message to the response
	fmt.Fprintf(w, "User deleted successfully")
}
