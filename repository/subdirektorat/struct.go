package subdirektorat

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSubdirektorat(db *mongo.Client) *MSubDirektorat {
	return &MSubDirektorat{
		db: db.Database("MailingRoom"),
	}
}

type MSubDirektorat struct {
	db *mongo.Database
}
