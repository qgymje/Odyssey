package services

import (
	"Odyssey/models"
	"fmt"
	"math/rand"
)

type SMS struct {
	phone         string
	model_smscode *models.SMSCode
}

func (s *SMS) Generate() string {
	rand.Seed(42)
	code := rand.Intn(10000)
	return fmt.Sprintf("%d", code)
}

func (s *SMS) Valid(code string) error {
	return nil
}

// 判断这个phone号码是否已经请求过验证码了
func (s *SMS) IsRequestedCode() bool {
	return true
}
