package model

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	RoleID    string `json:"role_id" bson:"role_id"` // pengirim, kurir, penerima
	Divisi    string `json:"divisi" bson:"divisi"`
	Phone     string `json:"phone" bson:"phone"`
	Password  string `json:"password" bson:"password"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
}

type Roles struct {
	RoleID   string `json:"role_id" bson:"role_id"`
	RoleName string `json:"role_name" bson:"role_name"`
}

type Naskah struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	NoSurat     string `json:"noSurat" bson:"noSurat"`
	Judul       string `json:"judul" bson:"judul"`
	JenisNaskah string `json:"jenisNaskah" bson:"jenisNaskah"`
	Kategori    string `json:"kategori" bson:"kategori"`
	Deskripsi   string `json:"deskripsi" bson:"deskripsi"`
	FilePath    string `json:"filePath" bson:"filePath"`
	Status      string `json:"status" bson:"status"` // draft, dikirim, diterima, ditolak
	PengirimID  string `json:"pengirimID" bson:"pengirimID"`
	PenerimaID  string `json:"penerimaID" bson:"penerimaID"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}

type FormatSurat struct {
}

type Office struct {
	OfficeID   string `json:"office_id" bson:"office_id,omitempty"`
	NamaOffice string `json:"nama_office" bson:"nama_office"`
	Alamat     string `json:"alamat" bson:"alamat"`
	Kota       string `json:"kota" bson:"kota"`
	KodePos    string `json:"kode_pos" bson:"kode_pos"`
	NoTelp     string `json:"no_telp" bson:"no_telp"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
	UpdatedAt  string `json:"updatedAt" bson:"updatedAt"`
}

type Departemen struct {
	DepartemenID   string `json:"departemen_id" bson:"departemen_id,omitempty"`
	NamaDepartemen string `json:"nama_departemen" bson:"nama_departemen"`
	KodeDepartemen string `json:"kode_departemen" bson:"kode_departemen"`
	NoTelp         string `json:"no_telp" bson:"no_telp"`
	CreatedAt      string `json:"createdAt" bson:"createdAt"`
	UpdatedAt      string `json:"updatedAt" bson:"updatedAt"`
}

type Pengiriman struct {
	PengirimanID string `json:"pengiriman_id" bson:"pengiriman_id,omitempty"`
	NaskahID     string `json:"naskah_id" bson:"naskah_id"`
	KurirID      string `json:"kurir_id" bson:"kurir_id"`
	TglKirim     string `json:"tgl_kirim" bson:"tgl_kirim"`
	TglTerima    string `json:"tgl_terima" bson:"tgl_terima"`
	LokasiAwal   string `json:"lokasi_awal" bson:"lokasi_awal"`
	LokasiTujuan string `json:"lokasi_tujuan" bson:"lokasi_tujuan"`
	Status       string `json:"status" bson:"status"` // dikirim, diterima, ditolak
	Catatan      string `json:"catatan" bson:"catatan"`
	UpdatedAt    string `json:"updatedAt" bson:"updatedAt"`
}

type JwtClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type Tracking struct {
}
