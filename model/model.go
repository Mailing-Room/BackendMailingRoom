package model

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	UserID          string `json:"user_id" bson:"_id,omitempty"`
	RoleID          string `json:"role_id" bson:"role_id"` // pengirim, kurir, penerima
	OfficeID        string `json:"office_id" bson:"office_id"`
	SubDirektoratID string `json:"sub_direktorat_id" bson:"sub_direktorat_id"`
	Name            string `json:"name" bson:"name"`
	Email           string `json:"email" bson:"email"`
	DivisiID        string `json:"divisi_id" bson:"divisi_id"`
	Phone           string `json:"phone" bson:"phone"`
	Password        string `json:"password" bson:"password"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	UpdatedAt       string `json:"updated_at" bson:"updated_at"`
}

type Roles struct {
	RoleID   string `json:"role_id" bson:"_id,omitempty"`
	RoleName string `json:"role_name" bson:"role_name"`
}

type Naskah struct {
	NaskahID         string           `json:"naskah_id" bson:"_id,omitempty"`
	NoNaskah         string           `json:"no_naskah" bson:"no_naskah"`
	Judul            string           `json:"judul" bson:"judul"`
	JenisNaskah      string           `json:"jenis_naskah" bson:"jenis_naskah"`
	SifatNaskah      string           `json:"sifat_naskah" bson:"sifat_naskah"`
	KategoriID       string           `json:"kategori_id" bson:"kategori_id"`
	Perihal          string           `json:"perihal" bson:"perihal"`
	MetodePengiriman MetodePengiriman `json:"metode_pengiriman" bson:"metode_pengiriman"` // digital, fisik
	FileDigital      string           `json:"file_digital" bson:"file_digital"`
	Status           string           `json:"status" bson:"status"` // draft, dikirim, diterima, ditolak
	PengirimID       string           `json:"pengirim_id" bson:"pengirim_id"`
	PenerimaID       string           `json:"penerima_id" bson:"penerima_id"`
	Petugas          string           `json:"petugas_id" bson:"petugas_id"`
	CreatedAt        string           `json:"created_at" bson:"created_at"`
	UpdatedAt        string           `json:"updated_at" bson:"updated_at"`
}

type Kategori struct {
	KategoriID   string `json:"kategori_id" bson:"_id,omitempty"`
	NamaKategori string `json:"nama_kategori" bson:"nama_kategori"`
}

type Office struct {
	OfficeID   string `json:"office_id" bson:"_id,omitempty"`
	NamaOffice string `json:"nama_office" bson:"nama_office"`
	Alamat     string `json:"alamat" bson:"alamat"`
	Kota       string `json:"kota" bson:"kota"`
	KodePos    string `json:"kode_pos" bson:"kode_pos"`
	NoTelp     string `json:"no_telp" bson:"no_telp"`
	CreatedAt  string `json:"created_at" bson:"created_at"`
	UpdatedAt  string `json:"updated_at" bson:"updated_at"`
}

type SubDirektorat struct {
	SubDirektoratID   string `json:"sub_direktorat_id" bson:"_id,omitempty"`
	OfficeID          string `json:"office_id" bson:"office_id"`
	NamaSubDirektorat string `json:"nama_sub_direktorat" bson:"nama_sub_direktorat"`
	KodeSubDirektorat string `json:"kode_sub_direktorat" bson:"kode_sub_direktorat"`
	NoTelp            string `json:"no_telp" bson:"no_telp"`
	CreatedAt         string `json:"created_at" bson:"created_at"`
	UpdatedAt         string `json:"updated_at" bson:"updated_at"`
}

type Divisi struct {
	DivisiID        string `json:"divisi_id" bson:"_id,omitempty"`
	SubDirektoratID string `json:"sub_direktorat_id" bson:"sub_direktorat_id"`
	NamaDivisi      string `json:"nama_divisi" bson:"nama_divisi"`
	KodeDivisi      string `json:"kode_divisi" bson:"kode_divisi"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	UpdatedAt       string `json:"updated_at" bson:"updated_at"`
}

type JwtClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type Response struct {
	Response string `json:"response"`
	Info     string `json:"info,omitempty"`
	Status   string `json:"status,omitempty"`
	Location string `json:"location,omitempty"`
}

// Constants for MetodePengiriman
type MetodePengiriman string

const (
	MetodeDigital MetodePengiriman = "Digital"
	MetodeFisik   MetodePengiriman = "Fisik"
)

type SifatNaskah string

const (
	SifatBiasa   SifatNaskah = "Biasa"
	SifatSegera  SifatNaskah = "Segera"
	SifatPenting SifatNaskah = "Penting"
	SifatRahasia SifatNaskah = "Rahasia"
)
