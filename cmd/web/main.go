package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/habibielukman/pkg/db"
	"github.com/habibielukman/pkg/profile"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

type Pelatihan struct {
	Name            string `json:"name"`
	LinkPendaftaran string `json:"link_pendaftaran"`
	LinkStreaming   string `json:"link_streaming"`
	LinkPoster      string `json:"link_poster"`
	LinkLatihan     string `json:"link_latihan"`
	Pengantar       string `json:"pengantar"`
}

type Hasil struct {
	Name  string `json:"name"`
	Nilai string `json:"nilai"`
}

type Soal struct {
	Soal    string `json:"soal"`
	Jawaban string `json:"jawaban"`
}

type SoalSoal []Soal
type AllUsers []User
type AllPelatihan []Pelatihan
type Hasils []Hasil

var (
	outfile, _ = os.Create("logs.log")
	l          = log.New(outfile, "", 0)
)

func main() {
	http.HandleFunc("/getdb", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var users AllUsers
		for idx := range db.GetAllRows(db.GetConnection())[0] {
			users = append(users, User{Id:db.GetAllRows(db.GetConnection())[3][idx],Name: db.GetAllRows(db.GetConnection())[0][idx], Email: db.GetAllRows(db.GetConnection())[1][idx], Password: db.GetAllRows(db.GetConnection())[2][idx]})
		}
		js, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			l.Println(err)
		}
		w.Write(js)
	})

	http.HandleFunc("/getpelatihan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var users AllPelatihan
		for idx := range db.GetAllRowsPelatihan(db.GetConnection())[0] {
			users = append(users, Pelatihan{Name: db.GetAllRowsPelatihan(db.GetConnection())[0][idx], LinkPendaftaran: db.GetAllRowsPelatihan(db.GetConnection())[1][idx], LinkStreaming: db.GetAllRowsPelatihan(db.GetConnection())[2][idx], LinkPoster: db.GetAllRowsPelatihan(db.GetConnection())[3][idx], LinkLatihan: db.GetAllRowsPelatihan(db.GetConnection())[4][idx], Pengantar: db.GetAllRowsPelatihan(db.GetConnection())[5][idx]})
		}
		js, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			l.Println(err)
		}
		w.Write(js)
	})

	http.HandleFunc("/addUser", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		// mydb := db.GetAllRowsPelatihan(db.GetConnection())
		db.InsertRow(name, email, password, 2, db.GetConnection())
	})

	http.HandleFunc("/addWorks", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		nilai, _ := strconv.Atoi(r.FormValue("nilai"))
		pelatihanId, _ := strconv.Atoi(r.FormValue("pelatihanID"))
		l.Println(username, nilai, (pelatihanId + 1))
		db.InsertDoneQuiz(username, nilai, pelatihanId+1, db.GetConnection())
	})

	http.HandleFunc("/editprofil", profile.EditProfilePost)

	for i := range db.GetAllRowsPelatihan(db.GetConnection())[0] {
		http.HandleFunc(fmt.Sprintf("/getworks%d", i), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			var users Hasils
			for idx := range db.GetAllQuiz(db.GetConnection(), strconv.Itoa(i+1))[0] {
				users = append(users, Hasil{Name: db.GetAllQuiz(db.GetConnection(), strconv.Itoa(i+1))[0][idx], Nilai: db.GetAllQuiz(db.GetConnection(), strconv.Itoa(i+1))[1][idx]})
			}
			js, err := json.MarshalIndent(users, "", "    ")
			if err != nil {
				l.Println(err)
			}
			w.Write(js)
		})

		http.HandleFunc(fmt.Sprintf("/getquiz%d", i), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			js, err := json.MarshalIndent(db.GetSoal(db.GetConnection(), i+1), "", "    ")
			if err != nil {
				l.Println(err)
			}
			w.Write(js)
		})
	}

	fmt.Println("Running on port 4000")
	http.ListenAndServe(":4000", nil)
}
