package web

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func validateInput(dataSet interface{}) {
	validate := validator.New()
	err := validate.Struct(dataSet)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 組み込み型のゼロ値は無視する
			switch err.Value().(type) {
			case string:
				if err.Value() != "" {
					fmt.Println(err)
				}
			case int:
				if err.Value() != 0 {
					fmt.Println(err)
				}
			}
		}
		return
	}
}
