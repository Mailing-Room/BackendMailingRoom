package naskah

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewNaskah(db *mongo.Client) *MNaskah {
	return &MNaskah{
		db: db.Database("MailingRoom"),
	}
}

type MNaskah struct {
	db *mongo.Database
}
