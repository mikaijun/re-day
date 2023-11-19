package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

func LoginFunc(db *gorm.DB, id string) (string, error) {
	user := &model.User{}
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return "", errors.New("IDが存在しません")
	}

	expirie := time.Now().AddDate(0, 0, 1)
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirie.Unix(),
	}
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("SIGNED_KEY")))
	if err != nil {
		return "", errors.New("トークン生成できませんでした")
	}
	authExpirie := &model.AuthExpirie{
		ID:        uuid.New().String(),
		UserId:    user.ID,
		ExpiresAt: expirie,
	}

	if err := db.Create(authExpirie).Error; err != nil {
		return "", errors.New("authExpirie fall error")
	}

	return tokenString, nil
}
