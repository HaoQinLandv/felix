package models

import (
	"errors"
	"fmt"
	"github.com/dejavuzhou/felix/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

var AppSecret = utils.RandomString(32)
var AppIss = "github.com/dejavuzhou/felix"
var ExpireTime = time.Hour * 24

func jwtGenerateToken(m *User) (*jwtObj, error) {
	m.Password = ""
	expireAfterTime := time.Hour * 24
	expireTime := time.Now().Add(expireAfterTime)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.ID),
		Issuer:    AppIss,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	data := &jwtObj{User: *m, Token: tokenString, Expire: expireTime, ExpireTs: expireTime.Unix()}
	return data, err
}

type jwtObj struct {
	User
	Token    string    `json:"token"`
	Expire   time.Time `json:"expire"`
	ExpireTs int64     `json:"expire_ts"`
}

//JwtParseUser
func JwtParseUser(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return nil, errors.New("token is expired")
	}
	if !claims.VerifyIssuer(AppIss, true) {
		return nil, errors.New("token's issuer is wrong,greetings Hacker")
	}
	key := fmt.Sprintf("login:%s", claims.Id)
	jwtObj, err := mem.GetJwtObj(key)
	if err != nil {
		return nil, err
	}
	return &jwtObj.User, err
}

//GetJwtObj
func (s *memoryStore) GetJwtObj(id string) (value *jwtObj, err error) {
	vv, err := s.Get(id, false)
	if err != nil {
		return nil, err
	}
	value, ok := vv.(*jwtObj)
	if ok {
		return value, nil
	}
	return nil, errors.New("mem:has value of this id, but is not type of *jwtObj")
}
