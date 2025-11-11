package office

import (
	"backendmailingroom/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (o *MOffice) InputOffice(ctx context.Context, office model.Office) (model.Office, error) {
	collection := o.db.Collection("office")

	officeData := bson.M{
		"nama_office": office.NamaOffice,
		"alamat":      office.Alamat,
		"kota":        office.Kota,
		"kode_pos":    office.KodePos,
		"no_telp":     office.NoTelp,
		"created_at":  time.Now().Format(time.RFC3339),
		"updated_at":  time.Now().Format(time.RFC3339),
	}

	result, err := collection.InsertOne(ctx, officeData)
	if err != nil {
		return model.Office{}, fmt.Errorf("gagal menyimpan data office: %w", err)
	}

	insertedID := fmt.Sprintf("%v", result.InsertedID)

	office.OfficeID = insertedID
	office.CreatedAt = officeData["created_at"].(string)
	office.UpdatedAt = officeData["updated_at"].(string)

	log.Println("[INFO] Data office berhasil disimpam dengan ID:", insertedID)
	return office, nil
}

func (o *MOffice) GetOfficeByID(ctx context.Context, id string) (model.Office, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Office{}, fmt.Errorf("id tidak valid: %w", err)
	}
	filter := bson.M{
		"_id": objectID,
	}
	var office model.Office
	collection := o.db.Collection("office")
	err = collection.FindOne(ctx, filter).Decode(&office)
	if err != nil {
		return model.Office{}, fmt.Errorf("gagal menemukan user: %w", err)
	}
	return office, nil
}

func (o *MOffice) GetAllOffice(ctx context.Context) ([]model.Office, error) {
	var offices []model.Office
	collection := o.db.Collection("office")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data office: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var office model.Office
		if err := cursor.Decode(&office); err != nil {
			log.Println("[ERROR] Gagal mendecode data office:", err)
			continue
		}
		offices = append(offices, office)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data office: %w", err)
	}
	return offices, nil
}

func (o *MOffice) GetOfficeByKota(ctx context.Context, kota string) ([]model.Office, error) {
	var offices []model.Office
	collection := o.db.Collection("office")

	filter := bson.M{
		"kota": bson.M{
			"$regex":   kota,
			"$options": "i", // case-insensitive
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Gagal mencari office berdasarkan kota: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var office model.Office
		if err := cursor.Decode(&office); err != nil {
			return nil, fmt.Errorf("Gagal decode data office: %w", err)
		}
		offices = append(offices, office)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("Terjadi kesalahan saat mengambil data office: %w", err)
	}

	return offices, nil
}

func (o *MOffice) DeleteOfficeByID(ctx context.Context, id string) (model.Office, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Office{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{
		"_id": objectID,
	}

	var deletedOffice model.Office
	collection := o.db.Collection("office")
	err = collection.FindOneAndDelete(ctx, filter).Decode(&deletedOffice)
	if err != nil {
		return model.Office{}, fmt.Errorf("gagal menghapus office: %w", err)
	}

	return deletedOffice, nil
}

func (o *MOffice) UpdateOffice(ctx context.Context, id string, updatedData model.Office) (model.Office, error) {
	collection := o.db.Collection("office")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Office{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{
		"_id": objectID,
	}

	update := bson.M{
		"$set": bson.M{
			"nama_office": updatedData.NamaOffice,
			"alamat":      updatedData.Alamat,
			"kota":        updatedData.Kota,
			"kode_pos":    updatedData.KodePos,
			"no_telp":     updatedData.NoTelp,
		},
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return model.Office{}, fmt.Errorf("office dengan id %s tidak ditemukan", id)
		}
		return model.Office{}, fmt.Errorf("gagal memperbarui office: %w", result.Err())
	}

	var updatedOffice model.Office
	if err := result.Decode(&updatedOffice); err != nil {
		return model.Office{}, fmt.Errorf("gagal mendecode data office yang diperbarui: %w", err)
	}

	return updatedOffice, nil
}
