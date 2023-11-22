package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	ctx context.Context
}

func (s *UserService) Login(id string) (string, error) {
	user := &model.User{}
	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
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

	if err := s.db.Create(authExpirie).Error; err != nil {
		return "", errors.New("既にログイン済みです")
	}

	return tokenString, nil
}

func (s *UserService) Logout() (bool, error) {
	userId := s.ctx.Value(model.AuthKey).(string)
	authExpirie := &model.AuthExpirie{}
	if err := s.db.Where("user_id = ?", userId).Delete(&authExpirie).Error; err != nil {
		return false, errors.New("ログアウトできませんでした")
	}
	return true, nil
}

func (s *UserService) CreateUser(name string) (*model.User, error) {
	user := model.User{
		ID:   uuid.New().String(),
		Name: name,
	}
	s.db.Create(&user)
	return &user, nil
}

func (s *UserService) FindUsers() ([]*model.User, error) {
	user := []*model.User{}
	s.db.Find(&user)
	return user, nil
}

func (s *UserService) FindUserByTask(task *model.Task) (*model.User, error) {
	user, err := loader.LoadUser(s.ctx, task.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
