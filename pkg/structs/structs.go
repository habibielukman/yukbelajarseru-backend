package structs

type Pelatihandata struct {
	Nilai           string
	Nama            string
	LinkPendaftaran string
	LinkStreaming   string
	LinkPoster      string
	LinkLatihan     string
	Pengantar       string
	Authenticated   bool
	Terdaftar       bool
}

// Struct
type Soal struct {
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
	Id         int `json:"id"`
}

// Struct for looping in /pelatihan
type LoopType struct {
	Number        int
	PelatihanData []Pelatihandata
}

// Struct for rendering
type A struct {
	UserID        int
	Nilai         int
	Add           int
	Loop          []LoopType
	SoalSoal      []Soal
	PelatihanId   int
	Nama          string
	Authenticated bool
	Username      string
	PelatihanData []Pelatihandata
}