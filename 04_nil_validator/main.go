package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/go-playground/validator/v10"
)

type PersonParams struct {
	FirstName string   `validate:"required"`
	Value     *float64 `validate:"omitempty,value-limit"`
}

func valueLimit(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(float64)
	if !ok {
		return true
	}
	if v < 0 || v > 100 {
		return false
	}
	return true
}

func main() {
	data := PersonParams{
		FirstName: "Taro",
		Value:     aws.Float64(88.88),
	}

	fmt.Println(data.Value)

	v := validator.New()
	v.RegisterValidation("value-limit", valueLimit)

	err := v.Struct(data)
	fmt.Println(err)

	// data = PersonParams{
	// 	FirstName: "Taro",
	// 	Value:     nil,
	// }

	// err = v.Struct(data)
	// fmt.Println(err)

	// data = PersonParams{
	// 	FirstName: "Taro",
	// }

	// err = v.Struct(data)
	// fmt.Println(err)
}
