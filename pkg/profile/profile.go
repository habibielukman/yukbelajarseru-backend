package profile

import (
	"encoding/json"
	"net/http"

	"github.com/habibielukman/pkg/db"
	"github.com/habibielukman/pkg/errhandler"
)

type Response struct {
	Message string `json:"message"`
}

func EditProfilePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := Response{Message: "Edit Profile Success"}
	js, _ := json.MarshalIndent(res, "", "    ")
	w.Write(js)
	nama := r.PostFormValue("nama")
	email := r.PostFormValue("email")
	userid := r.PostFormValue("UserID")
	nama_ := r.PostFormValue("Nama_")
	conn := db.GetConnection()
	
	if nama != nama_ && userid != email {
		query := `
	UPDATE users 
SET "name" = $3,
"email" = $2 
WHERE
	"email" = $1
	`
		query2 := `
		UPDATE works
		SET "user_id" = $1
		WHERE
			"user_id" = $2
		`
		_, err2 := conn.Exec(query2, email, userid)
		_, err2 = conn.Exec(query, userid, email, nama)
		errhandler.HandErr(err2)
	} else if nama != nama_ {
		query := `
	UPDATE users 
SET "name" = $2
WHERE
	"email" = $1
	`
		_, err2 := conn.Exec(query, userid, nama)
		errhandler.HandErr(err2)
	} else if userid != email {
		query := `
		UPDATE users 
SET "email" = $2 
WHERE
	"email" = $1
	`
		query2 := `
		UPDATE works
		SET "user_id" = $1
		WHERE
			"user_id" = $2
		`
		_, err2 := conn.Exec(query2, email, userid)
		_, err2 = conn.Exec(query, userid, email)
		errhandler.HandErr(err2)
	}
}