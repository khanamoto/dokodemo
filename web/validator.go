package web

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateAll(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err)
		}
		return errors.New("varidation error")
	}
	return nil
}

func validateLoginUser(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	errs := err.(validator.ValidationErrors)
	cap := len(errs)
	msg := make([]validator.FieldError, 0, cap)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := "UserName Password"
		if strings.Contains(fieldName, err.Field()) {
			msg = append(msg, err)
		}
	}
	if len(msg) != 0 {
		log.Println(msg)
		return errors.New("varidation error")
	}
	return nil
}

func validateBaseUser(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	errs := err.(validator.ValidationErrors)
	cap := len(errs)
	msg := make([]validator.FieldError, 0, cap)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := "Name UserName Email Password"
		if strings.Contains(fieldName, err.Field()) {
			msg = append(msg, err)
		}
	}
	if len(msg) != 0 {
		log.Println(msg)
		return errors.New("varidation error")
	}
	return nil
}

func validateUserName(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	errs := err.(validator.ValidationErrors)
	cap := len(errs)
	msg := make([]validator.FieldError, 0, cap)

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Field() {
		case "UserName":
			msg = append(msg, err)
		}
	}
	if len(msg) != 0 {
		log.Println(msg)
		return errors.New("varidation error")
	}
	return nil
}
