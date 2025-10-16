package departemen

import (
	"backendmailingroom/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *MDepartemen) InputDepartemen(ctx context.Context, departemen model.Departemen) (model.Departemen, error) {
	collection := d.db.Collection("departemen")

	departemenData := bson.M{
		"nama_departemen": departemen.NamaDepartemen,
		"kode_departemen": departemen.KodeDepartemen,
		"alamat":          departemen.Alamat,
		"kode_pos":        departemen.KodePos,
		"no_telp":         departemen.NoTelp,
		"createdAt":       time.Now().Format(time.RFC3339),
		"updatedAt":       time.Now().Format(time.RFC3339),
	}

	result, err := collection.InsertOne(ctx, departemenData)
	if err != nil {
		return model.Departemen{}, fmt.Errorf("gagal menyimpan data departemen: %w", err)
	}

	insertedID := fmt.Sprintf("%v", result.InsertedID)
	departemen.DepartemenID = insertedID
	departemen.CreatedAt = departemenData["createdAt"].(string)
	departemen.UpdatedAt = departemenData["updatedAt"].(string)

	log.Println("[INFO] Data departemen berhasil disimpan dengan ID:", insertedID)
	return departemen, nil
}
