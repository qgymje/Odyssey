package services

import (
	"errors"
	"fmt"
	"time"

	"Odyssey/utils"

	"github.com/dgrijalva/jwt-go"
)

var ErrGenerate = errors.New("生成token错误")
var ErrVerify = errors.New("验证token失败")

type Token struct {
	token     *jwt.Token
	secretKey []byte
	claims    map[string]interface{}
}

var defaultSecretKey string
var defaultExpires int64

func NewToken() *Token {
	defaultSecretKey = utils.GetConf().GetString("auth.secret_key")
	expires := utils.GetConf().GetInt("auth.expires") //second
	defaultExpires = time.Now().Add(time.Second * time.Duration(expires)).Unix()

	return &Token{
		secretKey: []byte(defaultSecretKey),
		claims:    make(map[string]interface{}),
	}
}

func (a *Token) Generate(claims map[string]interface{}) (tokenString string, err error) {
	var actualErr error

	defer func() {
		place := "services.Token.Generate"
		if err != nil {
			utils.GetLog().Error("%s : %s", place, actualErr.Error())
		} else {
			// can not leave the tokne string in logger file
			utils.GetLog().Debug("%s : generate token success", place)
		}
	}()

	a.token = jwt.New(jwt.SigningMethodHS256)
	for k, v := range claims {
		a.token.Claims[k] = v
	}
	if _, ok := a.token.Claims["exp"]; !ok {
		a.token.Claims["exp"] = defaultExpires
	}

	tokenString, err = a.token.SignedString(a.secretKey)
	if err != nil {
		actualErr = err
		err = ErrGenerate
	}

	return
}

func (a *Token) Verify(tokenString string) (valid bool, err error) {
	var actualErr error

	defer func() {
		place := "services.Token.Verify"
		if err != nil {
			utils.GetLog().Error("%s : %s", place, actualErr.Error())
		} else {
			utils.GetLog().Debug("%s : token verify success", place)
		}
	}()

	utils.GetLog().Debug("services.Token.Verify : tokenString : %s", tokenString)
	//自带过期处理
	a.token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.secretKey, nil
	})
	if err != nil {
		actualErr = err
		err = ErrVerify
	}
	valid = a.token.Valid

	return
}

func (a *Token) Claims() map[string]interface{} {
	return a.token.Claims
}
