package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func EncodeToken(userID string, roleID string) (string, error) {
	// Buat token baru dengan metode signing HMAC SHA256
	token := jwt.New(jwt.SigningMethodHS256)

	// Tambahkan klaim (data yang disimpan dalam token)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role_id"] = roleID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix() // kadaluarsa 24 jam

	// Kunci rahasia (disarankan simpan di .env)
	secretKey := []byte("secret")

	// Tanda tangani token dengan secret key
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func DecodeToken(tokenString string) (string, string, error) {
	// Ambil secret key (gunakan environment variable di production)
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		secretKey = []byte("secret") // fallback (hindari di production)
	}

	// Parse dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing-nya sesuai (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	// Ambil klaim dari token jika valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Ambil user_id dari klaim
		userID, ok := claims["user_id"].(string)
		roleID, ok2 := claims["role_id"].(string)
		if !ok {
			return "", "", errors.New("user_id not found or invalid type")
		}
		if !ok2 {
			return "", "", errors.New("role_id not found or invalid type")
		}
		return userID, roleID, nil
	}

	return "", "", errors.New("invalid token")
}
