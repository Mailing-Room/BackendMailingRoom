package users

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

// Fungsi untuk input data user
func (u *MUser) InputUser(ctx context.Context, user model.User) (model.User, error) {
	// Koleksi tempat data user disimpan
	collection := u.db.Collection("users")

	// Membuat dokumen BSON untuk disimpan ke MongoDB
	userData := bson.M{
		"role_id":           user.RoleID,
		"office_id":         user.OfficeID,
		"sub_direktorat_id": user.SubDirektoratID,
		"name":              user.Name,
		"email":             user.Email,
		"divisi_id":         user.DivisiID,
		"phone":             user.Phone,
		"password":          user.Password,
		"created_at":        time.Now().Format(time.RFC3339),
		"updated_at":        time.Now().Format(time.RFC3339),
	}

	// Menyimpan data ke MongoDB
	result, err := collection.InsertOne(ctx, userData)
	if err != nil {
		log.Println("[ERROR] Gagal menyimpan data user:", err)
		return model.User{}, fmt.Errorf("gagal menyimpan data user: %w", err)
	}

	// Ambil ID hasil insert (MongoDB ObjectID → string) - PERBAIKAN DI SINI
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("[ERROR] Gagal mengkonversi InsertedID ke ObjectID")
		return model.User{}, fmt.Errorf("gagal mengkonversi InsertedID")
	}

	user.UserID = insertedID.Hex() // Convert ke hex string
	user.CreatedAt = userData["created_at"].(string)
	user.UpdatedAt = userData["updated_at"].(string)

	log.Println("[INFO] Data user berhasil disimpan dengan ID:", user.UserID)
	return user, nil
}

func (u *MUser) GetUserForLogin(ctx context.Context, email string) (model.User, error) {
	filter := bson.M{
		"email": email,
	}
	log.Println("[INFO] Mencari user dengan filter:", filter)
	var user model.User
	collection := u.db.Collection("users")
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("gagal menemukan user: %w", err)
	}
	return user, nil
}

func (u *MUser) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	collection := u.db.Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data user: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("[ERROR] Gagal mendecode data user:", err)
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data user: %w", err)
	}

	return users, nil
}

func (u *MUser) GetUserByID(ctx context.Context, id string) (model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("id tidak valid: %w", err)
	}
	filter := bson.M{
		"_id": objectID,
	}
	var user model.User
	collection := u.db.Collection("users")
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("gagal menemukan user: %w", err)
	}
	return user, nil
}

func (u *MUser) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	filter := bson.M{
		"email": email,
	}
	log.Println("[INFO] Mencari user dengan filter:", filter)
	var user model.User
	collection := u.db.Collection("users")
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("gagal menemukan user: %w", err)
	}
	return user, nil
}

func (u *MUser) DeleteUserByID(ctx context.Context, id string) (model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{
		"_id": objectID,
	}
	log.Println("[INFO] Menghapus user dengan filter:", filter)

	var deletedUser model.User
	collection := u.db.Collection("users")
	err = collection.FindOneAndDelete(ctx, filter).Decode(&deletedUser)
	if err != nil {
		return model.User{}, fmt.Errorf("user dengan id %s tidak ditemukan", id)
	}

	log.Println("[INFO] User berhasil dihapus dengan ID:", id)
	return deletedUser, nil
}

func (u *MUser) UpdateUser(ctx context.Context, id string, updatedData model.User) (model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"role_id":           updatedData.RoleID,
			"office_id":         updatedData.OfficeID,
			"sub_direktorat_id": updatedData.SubDirektoratID,
			"name":              updatedData.Name,
			"email":             updatedData.Email,
			"divisi_id":         updatedData.DivisiID,
			"phone":             updatedData.Phone,
			"password":          updatedData.Password,
			"updated_at":        time.Now().Format(time.RFC3339),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	collection := u.db.Collection("users")

	var updatedUser model.User
	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		return model.User{}, fmt.Errorf("gagal memperbarui user: %w", err)
	}

	return updatedUser, nil
}
