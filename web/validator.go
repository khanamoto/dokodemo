package web

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

// TODO: fieldName以外DRYにする

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

func validateBaseUser(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	errs := err.(validator.ValidationErrors)
	cap := len(errs)
	msg := make([]validator.FieldError, 0, cap)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		if fieldName == "Name" || fieldName == "UserName" || fieldName == "Email" || fieldName == "Password" {
			msg = append(msg, err)
		}
	}
	if len(msg) != 0 {
		log.Println(msg)
		return errors.New("varidation error")
	}
	return nil
}

func validateUserNameAndPassword(dataSet interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataSet)
	errs := err.(validator.ValidationErrors)
	cap := len(errs)
	msg := make([]validator.FieldError, 0, cap)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		if fieldName == "UserName" || fieldName == "Password" {
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
		fieldName := err.Field()
		if fieldName == "UserName" {
			msg = append(msg, err)
		}
	}
	if len(msg) != 0 {
		log.Println(msg)
		return errors.New("varidation error")
	}
	return nil
}
