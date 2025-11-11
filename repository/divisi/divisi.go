package divisi

import (
	"backendmailingroom/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// INPUT DIVISI FUNCTION HERE!
func (d *MDivisi) InputDivisi(ctx context.Context, divisi model.Divisi) (model.Divisi, error) {
	collection := d.db.Collection("divisi")

	divisiData := bson.M{
		"sub_direktorat_id": divisi.SubDirektoratID,
		"nama_divisi":       divisi.NamaDivisi,
		"kode_divisi":       divisi.KodeDivisi,
		"created_at":        time.Now().Format(time.RFC3339),
		"updated_at":        time.Now().Format(time.RFC3339),
	}

	result, err := collection.InsertOne(ctx, divisiData)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("gagal menyimpan data Divisi: %w", err)
	}

	// PERBAIKAN DI SINI - convert ObjectID ke Hex string
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("[ERROR] Gagal mengkonversi InsertedID ke ObjectID")
		return model.Divisi{}, fmt.Errorf("gagal mengkonversi InsertedID")
	}

	divisi.DivisiID = insertedID.Hex() // Gunakan .Hex() bukan fmt.Sprintf
	divisi.CreatedAt = divisiData["created_at"].(string)
	divisi.UpdatedAt = divisiData["updated_at"].(string)

	log.Println("[INFO] Data Divisi berhasil disimpan dengan ID:", divisi.DivisiID)
	return divisi, nil
}

func (d *MDivisi) GetDivisiByID(ctx context.Context, id string) (model.Divisi, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.Divisi{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{
		"_id": objectID,
	}

	var divisi model.Divisi
	collection := d.db.Collection("divisi")
	err = collection.FindOne(ctx, filter).Decode(&divisi)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("gagal menemukan divisi: %w", err)
	}
	return divisi, nil
}

func (d *MDivisi) GetAllDivisi(ctx context.Context) ([]model.Divisi, error) {
	var divisions []model.Divisi
	collection := d.db.Collection("divisi")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data divisi: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var divisi model.Divisi
		if err := cursor.Decode(&divisi); err != nil {
			log.Println("[ERROR] Gagal mendecode data divisi:", err)
			continue
		}
		divisions = append(divisions, divisi)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data divisi: %w", err)
	}
	return divisions, nil
}

func (d *MDivisi) GetDivisiBySubDirektoratID(ctx context.Context, subDirektoratID string) ([]model.Divisi, error) {
	var divisions []model.Divisi
	collection := d.db.Collection("divisi")

	filter := bson.M{
		"sub_direktorat_id": subDirektoratID,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data divisi: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var divisi model.Divisi
		if err := cursor.Decode(&divisi); err != nil {
			log.Println("[ERROR] Gagal mendecode data divisi:", err)
			continue
		}
		divisions = append(divisions, divisi)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data divisi: %w", err)
	}

	return divisions, nil
}

// GetDivisiBySubDirektoratName mengambil semua divisi berdasarkan nama SubDirektorat
func (d *MDivisi) GetDivisiBySubDirektoratName(ctx context.Context, namaSubDirektorat string) ([]model.Divisi, error) {
	var divisions []model.Divisi

	// Pertama, cari SubDirektorat berdasarkan nama
	subDirektoratCollection := d.db.Collection("sub_direktorat")
	var subDir model.SubDirektorat

	err := subDirektoratCollection.FindOne(ctx, bson.M{"nama_sub_direktorat": namaSubDirektorat}).Decode(&subDir)
	if err != nil {
		return nil, fmt.Errorf("subdirektorat '%s' tidak ditemukan: %w", namaSubDirektorat, err)
	}

	// Kemudian, cari semua divisi yang terkait dengan SubDirektorat tersebut
	collection := d.db.Collection("divisi")
	filter := bson.M{
		"sub_direktorat_id": subDir.SubDirektoratID,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data divisi: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var divisi model.Divisi
		if err := cursor.Decode(&divisi); err != nil {
			log.Println("[ERROR] Gagal mendecode data divisi:", err)
			continue
		}
		divisions = append(divisions, divisi)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengiterasi cursor divisi: %w", err)
	}

	log.Printf("[INFO] Ditemukan %d divisi untuk subdirektorat '%s'", len(divisions), namaSubDirektorat)
	return divisions, nil
}

func (d *MDivisi) DeleteDivisiByID(ctx context.Context, id string) (model.Divisi, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{"_id": objectID}
	log.Println("[INFO] Mencari divisi dengan filter:", filter)

	var deletedDivisi model.Divisi
	collection := d.db.Collection("divisi")
	err = collection.FindOneAndDelete(ctx, filter).Decode(&deletedDivisi)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("divisi dengan id %s tidak ditemukan: %w", id, err)
	}

	log.Println("[INFO] Data divisi berhasil dihapus dengan ID:", id)
	return deletedDivisi, nil
}

func (d *MDivisi) UpdateDivisi(ctx context.Context, id string, updatedData model.Divisi) (model.Divisi, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"sub_direktorat_id": updatedData.SubDirektoratID,
			"nama_divisi":       updatedData.NamaDivisi,
			"kode_divisi":       updatedData.KodeDivisi,
			"updated_at":        time.Now().Format(time.RFC3339),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	collection := d.db.Collection("divisi")

	var updatedDivisi model.Divisi
	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDivisi)
	if err != nil {
		return model.Divisi{}, fmt.Errorf("gagal memperbarui divisi dengan id %s: %w", id, err)
	}

	return updatedDivisi, nil
}
