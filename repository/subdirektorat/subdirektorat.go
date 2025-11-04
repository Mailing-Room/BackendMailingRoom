package subdirektorat

import (
	"backendmailingroom/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *MSubDirektorat) InputSubDirektorat(ctx context.Context, subdirektorat model.SubDirektorat) (model.SubDirektorat, error) {
	collection := d.db.Collection("sub_direktorat")

	// Data yang akan disimpan ke MongoDB
	subdirektoratData := bson.M{
		"nama_sub_direktorat": subdirektorat.NamaSubDirektorat,
		"kode_sub_direktorat": subdirektorat.KodeSubDirektorat,
		"no_telp":             subdirektorat.NoTelp,
		"createdAt":           time.Now().Format(time.RFC3339),
		"updatedAt":           time.Now().Format(time.RFC3339),
	}

	// Insert ke MongoDB
	result, err := collection.InsertOne(ctx, subdirektoratData)
	if err != nil {
		return model.SubDirektorat{}, fmt.Errorf("gagal menyimpan data Sub Direktorat: %w", err)
	}

	// Ambil ID hasil insert
	insertedID := fmt.Sprintf("%v", result.InsertedID)

	// Set nilai ke struct untuk dikembalikan
	subdirektorat.SubDirektoratID = insertedID
	subdirektorat.CreatedAt = subdirektoratData["createdAt"].(string)
	subdirektorat.UpdatedAt = subdirektoratData["updatedAt"].(string)

	log.Println("[INFO] Data Sub Direktorat berhasil disimpan dengan ID:", insertedID)
	return subdirektorat, nil
}
