package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AA struct {
	Status string `validate:"required"`
	Value1 *float64
	Value2 *float64
	Value3 *float64
}

func main() {
	aa := AA{
		Status: "ok",
		Value1: nil,
		Value2: nil,
		Value3: nil,
	}

	v := validator.New()
	v.RegisterStructValidation(validateStruct, AA{})

	err := v.Struct(aa)
	fmt.Println(err)

	aa = AA{}

	err = v.Struct(aa)
	fmt.Println(err)

	aa = AA{
		Status: "no",
	}

	err = v.Struct(aa)
	fmt.Println(err)
}

func validateStruct(sl validator.StructLevel) {
	aa := sl.Current().Interface().(AA)

	if aa.Status == "ok" {
		if aa.Value1 != nil && aa.Value2 != nil && aa.Value3 != nil {
			if *aa.Value1 < 0 || *aa.Value1 > 100 {
				sl.ReportError(aa.Value1, "Value1", "Value1", "required", "")
			}
			if *aa.Value2 < 0 || *aa.Value2 > 100 {
				sl.ReportError(aa.Value2, "Value2", "Value2", "required", "")
			}
			if *aa.Value3 < 0 || *aa.Value3 > 100 {
				sl.ReportError(aa.Value3, "Value3", "Value3", "required", "")
			}
		} else {
			sl.ReportError(aa, "Value1", "Value1", "required", "")
			sl.ReportError(aa, "Value2", "Value2", "required", "")
			sl.ReportError(aa, "Value3", "Value3", "required", "")
		}
	}
}
