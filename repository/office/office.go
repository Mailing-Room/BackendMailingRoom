package office

import (
	"backendmailingroom/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (o *MOffice) InputOffice(ctx context.Context, office model.Office) (model.Office, error) {
	collection := o.db.Collection("office")

	officeData := bson.M{
		"nama_office": office.NamaOffice,
		"alamat":      office.Alamat,
		"kota":        office.Kota,
		"kode_pos":    office.KodePos,
		"no_telp":     office.NoTelp,
		"createdAt":   time.Now().Format(time.RFC3339),
		"updatedAt":   time.Now().Format(time.RFC3339),
	}

	result, err := collection.InsertOne(ctx, officeData)
	if err != nil {
		return model.Office{}, fmt.Errorf("gagal menyimpan data office: %w", err)
	}

	insertedID := fmt.Sprintf("%v", result.InsertedID)

	office.OfficeID = insertedID
	office.CreatedAt = officeData["createdAt"].(string)
	office.UpdatedAt = officeData["updatedAt"].(string)

	log.Println("[INFO] Data office berhasil disimpam dengan ID:", insertedID)
	return office, nil
}
