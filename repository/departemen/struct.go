package departemen

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewDepartemen(db *mongo.Client) *MDepartemen {
	return &MDepartemen{
		db: db.Database("MailingRoom"),
	}
}

type MDepartemen struct {
	db *mongo.Database
}
