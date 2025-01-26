package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type PersonParams struct {
	FirstName string `json:"firstName" validate:"required"`     //デフォルトで使用可能な必須チェックを行うバリデーション。
	Age       int    `json:"age" validate:"required,age-limit"` //カスタムバリデーション。独自の条件でバリデーションを作成し使用できるようにする。
}

type Validate struct {
	validator *validator.Validate
}

func NewValidator() Validate {
	return Validate{validator: validator.New()}
}

func (p *PersonParams) Validate() error {
	v := validator.New()
	v.RegisterValidation("age-limit", ageLimit)

	return v.Struct(p)
}

func ageLimit(fl validator.FieldLevel) bool {
	a, ok := fl.Field().Interface().(int)
	if !ok {
		return false
	}
	if a < 20 {
		return false
	}
	return true
}

func main() {
	person := []PersonParams{
		{FirstName: "Taro", Age: 20},
		{FirstName: "Taro", Age: 19},
	}
	for _, p := range person {
		err := p.Validate()
		fmt.Println(err)
	}
}
