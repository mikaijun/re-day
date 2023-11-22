package service

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, id string) (string, error) {
	log.Printf("%v\n", 'a')
	user := &model.User{}
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return "", errors.New("ユーザーが存在しません")
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
		return "", errors.New("既にログイン済みです")
	}

	return tokenString, nil
}

func Logout(ctx context.Context, db *gorm.DB) (bool, error) {
	userId := ctx.Value(model.AuthKey).(string)
	authExpirie := &model.AuthExpirie{}
	if err := db.Where("user_id = ?", userId).Delete(&authExpirie).Error; err != nil {
		return false, errors.New("ログアウトできませんでした")
	}
	return true, nil
}

func CreateUser(name string, db *gorm.DB) (*model.User, error) {
	user := model.User{
		ID:   uuid.New().String(),
		Name: name,
	}
	db.Create(&user)
	return &user, nil
}

func FindUsers(db *gorm.DB) ([]*model.User, error) {
	user := []*model.User{}
	db.Find(&user)
	return user, nil
}

func FindUserByTask(ctx context.Context, task *model.Task) (*model.User, error) {
	user, err := loader.LoadUser(ctx, task.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
