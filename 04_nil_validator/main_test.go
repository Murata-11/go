package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/go-playground/validator/v10"
)

func TestXxx(t *testing.T) {
	data := PersonParams{
		FirstName: "Taro",
		Value:     aws.Float64(88.88),
	}

	fmt.Println(data.Value)

	v := validator.New()
	v.RegisterValidation("value-limit", valueLimit)

	err := v.Struct(data)
	fmt.Println(err)

	data = PersonParams{
		FirstName: "Taro",
		Value:     nil,
	}

	err = v.Struct(data)
	fmt.Println(err)
}
