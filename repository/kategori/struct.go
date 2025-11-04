package kategori

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewKategori(db *mongo.Client) *MKategori {
	return &MKategori{
		db: db.Database("MailingRoom"),
	}
}

type MKategori struct {
	db *mongo.Database
}
