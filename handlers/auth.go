package handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/daystram/go-gin-gorm-boilerplate/config"
	"github.com/daystram/go-gin-gorm-boilerplate/constants"
	"github.com/daystram/go-gin-gorm-boilerplate/datatransfers"
	"github.com/daystram/go-gin-gorm-boilerplate/models"
)

func (m *module) AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error) {
	var user models.User
	if user, err = m.db.userOrmer.GetOneByUsername(credentials.Username); err != nil {
		return "", errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("incorrect credentials")
	}

	return generateToken(user)
}

func generateToken(user models.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(constants.AuthenticationTimeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datatransfers.JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func (m *module) RegisterUser(credentials datatransfers.UserSignup) (err error) {
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}
	if _, err = m.db.userOrmer.InsertUser(models.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: string(hashedPassword),
		Bio:      credentials.Bio,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting user. %v", err)
	}
	return
}
