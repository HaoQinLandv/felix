package internal

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

const iss = "dejavuzhou.felix.ssh2ws"

//var appSecret = randomString(32)

func jwtGenerateToken(jwtIdValue string, expire time.Duration, secretBytes []byte) (*jwtObj, error) {
	expireTime := time.Now().Add(expire)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        jwtIdValue,
		Issuer:    iss,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretBytes)
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	data := &jwtObj{Token: tokenString, Expire: expireTime, ExpireTs: expireTime.Unix()}
	return data, err
}

type jwtObj struct {
	Token    string    `json:"token"`
	Expire   time.Time `json:"expire"`
	ExpireTs int64     `json:"expire_ts"`
}

//jwtParseUser parse a jwt token and return an authorized identity
func jwtParseUser(tokenString string, secretBytes []byte) (string, error) {
	if tokenString == "" {
		return "", errors.New("token is not found in Authorization Bearer")
	}
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretBytes, nil
	})
	if err != nil {
		return "", err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return "", errors.New("token is expired")
	}
	if !claims.VerifyIssuer(iss, true) {
		return "", errors.New("token's issuer is wrong,greetings Hacker")
	}
	return claims.Id, err
}
