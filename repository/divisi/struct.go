package divisi

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewDivisi(db *mongo.Client) *MDivisi {
	return &MDivisi{
		db: db.Database("MailingRoom"),
	}
}

type MDivisi struct {
	db *mongo.Database
}
