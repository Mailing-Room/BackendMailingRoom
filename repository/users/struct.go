package users

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUser(db *mongo.Client) *MUser {
	return &MUser{
		db: db.Database("MailingRoom"),
	}
}

type MUser struct {
	db *mongo.Database
}
