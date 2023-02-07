package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type user struct {
	Id              string     `json:"id"`
	SecretCode      string     `json:"secretcode"`
	Name            string     `json:"name"`
	EmailID         string     `json:"emailID"`
	ListofPlaylists []playlist `json:"listofplaylists"`
}
type playlist struct {
	PId         string `json:"pid"`
	Name        string `json:"name"`
	ListofSongs []song `json:"listofsongs"`
}
type song struct {
	SId      string `json:"sid"`
	Name     string `json:"name"`
	Composer string `json:"composer"`
}

var playlists1 = []playlist{
	{
		PId:         "3456",
		Name:        "goodmorning",
		ListofSongs: songs1,
	},
	{
		PId:         "3457",
		Name:        "heelson",
		ListofSongs: songs1,
	},
}

var songs1 = []song{
	{
		SId:      "3456",
		Name:     "believer",
		Composer: "Imagine Dragons",
	},
}

var users = []user{
	{
		Id:              "123",
		SecretCode:      "34576",
		Name:            "Ishika",
		EmailID:         "34576",
		ListofPlaylists: playlists1,
	},
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println(r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var usr user
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}

		t, err := getUserByAttributes(usr.SecretCode)
		if err != nil {
			errors.New("user not found")
			return
		}

		err2 := json.NewEncoder(w).Encode(t)
		if err2 != nil {
			http.Error(w, err2.Error(), 500)
			return
		}
	}
}

func getUserByAttributes(secretCode string) (*user, error) {

	for i, t := range users {
		if t.SecretCode == secretCode {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func registerUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println(r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var newuser user

	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&newuser)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		newuser.SecretCode = string(rand.Intn(20000))
		newuser.ListofPlaylists = []playlist{}
		newuser.Id = string(rand.Intn(20000))

		users = append(users, newuser)

		if err = json.NewEncoder(w).Encode(&newuser); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//print(newuser)
	}
}

func viewprofile(w http.ResponseWriter, r *http.Request) {
	var usr user
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println(r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}

		t, err := getUserByAttributes(usr.SecretCode)
		if err != nil {
			errors.New("user not found")
			return
		}

		if err2 := json.NewEncoder(w).Encode(t); err2 != nil {
			http.Error(w, err2.Error(), 500)
			return
		}
	}
}
func getsongsofplaylist(w http.ResponseWriter, r *http.Request) {

	var usr user
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}

		t, err := getUserByAttributes(usr.SecretCode)
		if err != nil {
			errors.New("user not found")
			return
		}
		var p playlist
		for _, p = range t.ListofPlaylists {

			if p.Name == usr.ListofPlaylists[0].Name {
				break
			}
		}
		if er := json.NewEncoder(w).Encode(p.ListofSongs); er != nil {
			http.Error(w, er.Error(), 400)
		}
	}
}
func createplaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println(r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var usr user
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}

		t, err := getUserByAttributes(usr.SecretCode)
		if err != nil {
			errors.New("user not found")
			return
		}
		t.ListofPlaylists = append(t.ListofPlaylists, usr.ListofPlaylists...)

		if err2 := json.NewEncoder(w).Encode(t.ListofPlaylists); err2 != nil {
			http.Error(w, err2.Error(), 500)
			return
		}
	}
}

func addsongstoplaylist(w http.ResponseWriter, r *http.Request) {

	var usr user

	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
			http.Error(w, err.Error(), 400)
		}
		flag := false

		for userIndex, tempUser := range users {
			if usr.SecretCode == tempUser.SecretCode {

				for playlistIndex, tempPlaylist := range users[userIndex].ListofPlaylists {

					if usr.ListofPlaylists[0].Name == tempPlaylist.Name {
						{
							p := users[userIndex].ListofPlaylists[playlistIndex].ListofSongs

							p = append(p, usr.ListofPlaylists[0].ListofSongs...)

							users[userIndex].ListofPlaylists[playlistIndex].ListofSongs = p

							if err := json.NewEncoder(w).Encode(p); err != nil {
								http.Error(w, err.Error(), 400)
							}

							flag = true
							break
						}
					}
				}
				if flag {
					break
				}
			}
		}
	}
}

//handle concurrency
//file structuring

func deleteplaylist(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println(r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var usr user
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}

		fmt.Println(r.Body)
		fmt.Println(usr)
		var userIndex int
		var tempUser user
		for userIndex, tempUser = range users {
			fmt.Println(usr.SecretCode == tempUser.SecretCode, usr.SecretCode, tempUser.SecretCode)
			if usr.SecretCode == tempUser.SecretCode {
				for playlistIndex, tempPlaylist := range users[userIndex].ListofPlaylists {
					fmt.Println(usr.ListofPlaylists[0].PId == tempPlaylist.PId, usr.ListofPlaylists[0].PId, tempPlaylist.PId)
					if usr.ListofPlaylists[0].PId == tempPlaylist.PId {
						users[userIndex].ListofPlaylists =
							append(users[userIndex].ListofPlaylists[:playlistIndex], users[userIndex].ListofPlaylists[playlistIndex+1:]...)
						fmt.Println(users[userIndex].ListofPlaylists)
						break
					}
				}
				break
			}
		}

		json.NewEncoder(w).Encode(users[userIndex])

	}
}

func deletesong(w http.ResponseWriter, r *http.Request) {
	var usr user
	if r.Method == "POST" {

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil {
			http.Error(w, er.Error(), 400)
		}
		var userIndex, playlistIndex int
		var tempUser user
		var tempPlaylist playlist
		for userIndex, tempUser = range users {
			if usr.SecretCode == tempUser.SecretCode {

				for playlistIndex, tempPlaylist = range users[userIndex].ListofPlaylists {
					if usr.ListofPlaylists[0].PId == tempPlaylist.PId {

						for songIndex, tempSong := range tempPlaylist.ListofSongs {
							if usr.ListofPlaylists[0].ListofSongs[0].SId == tempSong.SId {

								users[userIndex].ListofPlaylists[playlistIndex].ListofSongs =
									append(tempPlaylist.ListofSongs[:songIndex],
										tempPlaylist.ListofSongs[songIndex+1:]...)
							}
						}
					}
				}
			}
		}
		if er := json.NewEncoder(w).Encode(users[userIndex].ListofPlaylists[playlistIndex].ListofSongs); er != nil {
			http.Error(w, er.Error(), 400)
		}
	}
}
func main() {

	http.HandleFunc("/login", loginUser)
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/viewprofile", viewprofile)
	http.HandleFunc("/getsongsofplaylist", getsongsofplaylist)
	http.HandleFunc("/createplaylist", createplaylist)
	http.HandleFunc("/deletesong", deletesong)
	http.HandleFunc("/addsongstoplaylist", addsongstoplaylist)
	http.HandleFunc("/deleteplaylist", deleteplaylist)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("listen and serve:", err)
	}

}
