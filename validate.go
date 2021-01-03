package main

import (
	"errors"
	"reflect"
	"strings"
)

type Validator interface {
	validate(interface{}) (bool, error)
}

type stringValidator struct {
	min int
	max int
}

func (s *stringValidator) validate(prop interface{}) (bool, error) {
	a := prop.(string)
	if len(a) < s.min {
		return false, errors.New("too short string")
	} else if len(a) > s.max {
		return false, errors.New("too long string")
	}
	return true, nil
}

func setValidator(cons string) Validator {
	a := strings.Split(cons, ",")

	if a[0] != "required" {

	}

}

func validateStruct(user User) []error {

	errs := []error{}

	props := reflect.ValueOf(user)
	for i := 0; i < props.NumField(); i++ {
		validator := setValidator(props.Type().Field(i).Tag.Get("validate"))

	}

	return errs

}
