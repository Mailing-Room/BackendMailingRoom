package model

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	UserID       string `json:"user_id" bson:"_id,omitempty"`
	RoleID       string `json:"role_id" bson:"role_id"` // pengirim, kurir, penerima
	OfficeID     string `json:"office_id" bson:"office_id"`
	DepartemenID string `json:"departemen_id" bson:"departemen_id"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	DivisiID     string `json:"divisi_id" bson:"divisi_id"`
	Jabatan      string `json:"jabatan" bson:"jabatan"`
	Phone        string `json:"phone" bson:"phone"`
	Password     string `json:"password" bson:"password"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
	UpdatedAt    string `json:"updatedAt" bson:"updatedAt"`
}

type Roles struct {
	RoleID   string `json:"role_id" bson:"_id,omitempty"`
	RoleName string `json:"role_name" bson:"role_name"`
}

type Naskah struct {
	NaskahID    string `json:"naskah_id" bson:"_id,omitempty"`
	NoSurat     string `json:"noSurat" bson:"noSurat"`
	Judul       string `json:"judul" bson:"judul"`
	JenisNaskah string `json:"jenisNaskah" bson:"jenisNaskah"`
	SifatNaskah string `json:"sifatNaskah" bson:"sifatNaskah"`
	KategoriID  string `json:"kategori_id" bson:"kategori_id"`
	Deskripsi   string `json:"deskripsi" bson:"deskripsi"`
	FilePath    string `json:"filePath" bson:"filePath"`
	Status      string `json:"status" bson:"status"` // draft, dikirim, diterima, ditolak
	PengirimID  string `json:"pengirim_id" bson:"pengirim_id"`
	PenerimaID  string `json:"penerima_id" bson:"penerima_id"`
	Kurir       string `json:"kurir_id" bson:"kurir_id"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
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
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
	UpdatedAt  string `json:"updatedAt" bson:"updatedAt"`
}

type SubDirektorat struct {
	SubDirektoratID   string `json:"sub_direktorat_id" bson:"_id,omitempty"`
	OfficeID          string `json:"office_id" bson:"office_id"`
	NamaSubDirektorat string `json:"nama_sub_direktorat" bson:"nama_sub_direktorat"`
	KodeSubDirektorat string `json:"kode_sub_direktorat" bson:"kode_sub_direktorat"`
	NoTelp            string `json:"no_telp" bson:"no_telp"`
	CreatedAt         string `json:"createdAt" bson:"createdAt"`
	UpdatedAt         string `json:"updatedAt" bson:"updatedAt"`
}

type Divisi struct {
	DivisiID        string `json:"divisi_id" bson:"_id,omitempty"`
	SubDirektoratID string `json:"sub_direktorat_id" bson:"sub_direktorat_id"`
	NamaDivisi      string `json:"nama_divisi" bson:"nama_divisi"`
	KodeDivisi      string `json:"kode_divisi" bson:"kode_divisi"`
	CreatedAt       string `json:"createdAt" bson:"createdAt"`
	UpdatedAt       string `json:"updatedAt" bson:"updatedAt"`
}

type JwtClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type Tracking struct {
}

type BalasanSurat struct {
}

// type Pengiriman struct {
// 	PengirimanID string `json:"pengiriman_id" bson:"_id,omitempty"`
// 	NaskahID     string `json:"naskah_id" bson:"naskah_id"`
// 	KurirID      string `json:"kurir_id" bson:"kurir_id"`
// 	TglKirim     string `json:"tgl_kirim" bson:"tgl_kirim"`
// 	TglTerima    string `json:"tgl_terima" bson:"tgl_terima"`
// 	LokasiAwal   string `json:"lokasi_awal" bson:"lokasi_awal"`
// 	LokasiTujuan string `json:"lokasi_tujuan" bson:"lokasi_tujuan"`
// 	Status       string `json:"status" bson:"status"` // dikirim, diterima, ditolak
// 	Catatan      string `json:"catatan" bson:"catatan"`
// 	UpdatedAt    string `json:"updatedAt" bson:"updatedAt"`
// }
