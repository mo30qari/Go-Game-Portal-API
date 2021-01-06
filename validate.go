package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var mailFormat = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Validator interface {
	validate(interface{}) (error)
}

type stringValidator struct {
	min int
	max int
}

func (s stringValidator) validate(prop interface{}) (error) {
	a := prop.(string)
	if len(a) < s.min {
		return errors.New("too short string")
	} else if len(a) > s.max {
		return errors.New("too long string")
	}
	return nil
}

type emailValidator struct{}

func (e emailValidator) validate(prop interface{}) (error) {
	a := prop.(string)
	if !mailFormat.MatchString(a) {
		return errors.New("wrong email")
	}
	return nil

}

type passwordValidator struct {
	min int
	max int
}

func (p passwordValidator) validate(prop interface{}) (error) {
	a := prop.(string)
	if len(a) < p.min {
		return errors.New("too short")
	} else if len(a) > p.max {
		return errors.New("too long")
	}
	return nil
}

type defaultValidator struct{}

func (d defaultValidator) validate(prop interface{}) (error) {
	return nil
}

func getValidator(tags string) Validator {
	splittedTags := strings.Split(tags, ",")

	switch splittedTags[0] {

	case "string":
		validator := stringValidator{}
		fmt.Sscanf(strings.Join(splittedTags[1:], ","), "Min=%d,Max=%d", &validator.min, &validator.max)
		return validator

	case "email":
		validator := emailValidator{}
		return validator

	case "password":
		validator := passwordValidator{}
		fmt.Sscanf(strings.Join(splittedTags[1:],","),"Min=%d,Max=%d", &validator.min, &validator.max)
		return validator
	}

	return defaultValidator{}
}

func validateStruct(str interface{}) []error {
	errs := []error{}

	props := reflect.ValueOf(str)

	for i := 0; i < props.NumField(); i++ {
		validator := getValidator(props.Type().Field(i).Tag.Get("validate"))
		err := validator.validate(props.Field(i).Interface())

		if err != nil {
			errs = append(errs, fmt.Errorf("%s %s", props.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}
