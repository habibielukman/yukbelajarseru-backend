package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/habibielukman/pkg/errhandler"
	"github.com/habibielukman/pkg/structs"
	_ "github.com/jackc/pgx/stdlib"
)

var (
    outfile, _ = os.Create("logs2.log")
    l      = log.New(outfile, "", 0)
)

var connection *sql.DB

func GetConnection() *sql.DB {
	if connection == nil {
		connection, _ = sql.Open("pgx", "host=database-1.cwd3pnzyxmeg.ap-southeast-3.rds.amazonaws.com port=5432 dbname=connect_db user=ironman password=Habibie123!")
	}
	return connection
}

func InsertRow(name string, email string, password string, access_level int, conn *sql.DB) {
	query := `insert into users (name, email, password, created_at, updated_at, access_level) values ($1, $2, $3, $4, $5, $6)`
	_, err2 := conn.Exec(query, name, email, password, time.Now(), time.Now(), access_level)

	errhandler.HandErr(err2)
}

func InsertDoneQuiz(user_id string, nilai int, pelatihan_id int, conn *sql.DB) {
	query := `insert into works (pelatihan_id, user_id, nilai) values ($1, $2, $3)`
	_, err2 := conn.Exec(query, pelatihan_id, user_id, nilai)

	errhandler.HandErr(err2)
}

func InsertJawaban(user_id string, soalid int , jawaban string, pelatihan_id int, conn *sql.DB) {
	query := `insert into jawabanuser (pelatihan_id, user_id, idsoal, jawabanUser) values ($1, $2, $3, $4)`
	_, err2 := conn.Exec(query, pelatihan_id,user_id, soalid, jawaban)

	errhandler.HandErr(err2)
}

func GetJawaban(conn *sql.DB, pelatihanid int,userid string) []string {
	rows, err := conn.Query("select idsoal,jawabanUser from jawabanuser where pelatihan_id = " + strconv.Itoa(pelatihanid) + ", user_id = '" + userid+"'")
	errhandler.HandErr(err)
	defer rows.Close()
	var idsoal int
	var jawabanUser string
	var jawabans []string
	for rows.Next() {
		err := rows.Scan(&idsoal, &jawabanUser)
		errhandler.HandErr(err)
		jawabans = append(jawabans, jawabanUser)
	}
	return jawabans
}

func GetAllQuiz(conn *sql.DB, id string) [][]string {
	rows, err := conn.Query("select user_id,nilai from works where pelatihan_id = " + id)
	errhandler.HandErr(err)
	defer rows.Close()
	var nilai int
	var user_id string
	var user_ids []string
	var nilais []string
	for rows.Next() {
		err := rows.Scan(&user_id, &nilai)
		errhandler.HandErr(err)
		user_ids = append(user_ids, user_id)
		nilais = append(nilais, strconv.Itoa(nilai))
	}
	return [][]string{user_ids, nilais}
}

func GetAllRows(conn *sql.DB) [][]string {
	rows, err := conn.Query("select id, name, email, password from users")
	errhandler.HandErr(err)
	defer rows.Close()

	var name, email, password string
	var id string
	var names []string
	var emails []string
	var passwords []string
	var ids []string
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &password)
		errhandler.HandErr(err)
		names = append(names, name)
		emails = append(emails, email)
		passwords = append(passwords, password)
		ids = append(ids, id)
	}
	return [][]string{names, emails, passwords, ids}
}

func GetSoal(conn *sql.DB, Id int) []structs.Soal {
	rows, err := conn.Query("SELECT * FROM soal WHERE pelatihan_id = " + strconv.Itoa(Id))
	l.Println(strconv.Itoa(Id))
	errhandler.HandErr(err)
	defer rows.Close()

	var pelatihan_id, soal, jawaban_betul string
	var id int
	var soalSoal []structs.Soal
	loop_index := 0
	for rows.Next() {
		err := rows.Scan(&id, &pelatihan_id, &soal, &jawaban_betul)
		errhandler.HandErr(err)
		soal2 := structs.Soal{
			Id:         loop_index,
			Pertanyaan: soal,
			Jawaban:    jawaban_betul,
		}
		soalSoal = append(soalSoal, soal2)
		l.Println(soal)
		loop_index++
	}
	return soalSoal
}

func GetAllRowsPelatihan(conn *sql.DB) [][]string {
	rows, err := conn.Query("select id, nama, link_pendaftaran, link_streaming, link_poster, link_latihan, pengantar from pelatihan")
	errhandler.HandErr(err)
	defer rows.Close()

	var nama, pengantar, linkPendaftaran, linkStreaming, linkPoster, linkLatihan string
	var id int
	var namas []string
	var linkPendaftarans []string
	var linkStreamings []string
	var linkPosters []string
	var linkLatihans []string
	var pengantars []string
	for rows.Next() {
		err := rows.Scan(&id, &nama, &linkPendaftaran, &linkStreaming, &linkPoster, &linkLatihan, &pengantar)
		errhandler.HandErr(err)
		namas = append(namas, nama)
		pengantars = append(pengantars, pengantar)
		linkPendaftarans = append(linkPendaftarans, linkPendaftaran)
		linkStreamings = append(linkStreamings, linkStreaming)
		linkPosters = append(linkPosters, linkPoster)
		linkLatihans = append(linkLatihans, linkLatihan)
	}
	return [][]string{namas, linkPendaftarans, linkStreamings, linkPosters, linkLatihans, pengantars}
}
