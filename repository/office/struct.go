package office

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewOffice(db *mongo.Client) *MOffice {
	return &MOffice{
		db: db.Database("MailingRoom"),
	}
}

type MOffice struct {
	db *mongo.Database
}
