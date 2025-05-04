package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var players = []player{}

func hellowhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hellow from go")
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func mathod_check(w http.ResponseWriter, r *http.Request, mathod string) bool {
	if r.Method != mathod {
		http.Error(w, "Invalid mathod", http.StatusBadRequest)
		return false
	}
	return true
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post allow", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusBadRequest)
		return
	}
	fmt.Printf("recived user: %+v\n", user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created", "name": user.Name,
	})
}

type player struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getPlayers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(players)
}
func createRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Mathod not allow", http.StatusBadRequest)
		return
	}
	var p player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	fmt.Print(p)
	if p.Password == "" || p.Username == "" {
		http.Error(w, "all fields are neccasary", http.StatusBadRequest)
		return
	}
	players = append(players, player{
		p.Username, p.Password,
	})
	fmt.Println("player saved succesfulyy!!! ", p)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "player saved succesfully",
		"username": p.Username,
	})

}
func login(w http.ResponseWriter, r *http.Request) {
	if !mathod_check(w, r, http.MethodPost) {
		return
	}
	var p player
	error := json.NewDecoder(r.Body).Decode(&p)
	if error != nil {
		http.Error(w, "invalid cradantial", http.StatusBadRequest)
		return
	}
	itr := -1
	for val := range players {
		if players[val].Username == p.Username {
			itr = val
			break
		}

	}
	if itr != -1 && players[itr].Password == p.Password {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message":  "login succesfull",
			"username": p.Username,
		})
	} else {
		http.Error(w, "invalid cradantial", http.StatusUnauthorized)
	}

}
func main() {
	http.HandleFunc("/", createUserHandler)
	http.HandleFunc("/reg", createRegister)
	http.HandleFunc("/users", getPlayers)
	http.HandleFunc("/login", login)
	fmt.Println("server started at port 8000")
	http.ListenAndServe(":8000", nil)
}
