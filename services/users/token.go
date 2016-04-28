package users

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

func (t *Token) Generate(claims map[string]interface{}) (tokenString string, err error) {
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

	t.token = jwt.New(jwt.SigningMethodHS256)
	for k, v := range claims {
		t.token.Claims[k] = v
	}
	if _, ok := t.token.Claims["exp"]; !ok {
		t.token.Claims["exp"] = defaultExpires
	}

	tokenString, err = t.token.SignedString(t.secretKey)
	if err != nil {
		actualErr = err
		err = ErrGenerate
	}

	return
}

func (t *Token) Verify(tokenString string) (valid bool, err error) {
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
	t.token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return t.secretKey, nil
	})
	if err != nil {
		actualErr = err
		err = ErrVerify
	}
	valid = t.token.Valid

	return
}

func (t *Token) Claims() map[string]interface{} {
	return t.token.Claims
}
