package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var mailFormat = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Validator interface {
	validate(interface{}) (bool, error)
}

type stringValidator struct {
	min int
	max int
}

func (s stringValidator) validate(prop interface{}) (bool, error) {
	a := prop.(string)
	if len(a) < s.min {
		return false, errors.New("too short string")
	} else if len(a) > s.max {
		return false, errors.New("too long string")
	}
	return true, nil
}

type emailValidator struct{}

func (e emailValidator) validate(prop interface{}) (bool, error) {
	a := prop.(string)
	if !mailFormat.MatchString(a) {
		return false, errors.New("wrong email")
	}
	return true, nil

}

type passwordValidator struct {
	min int
	max int
}

func (p passwordValidator) validate(prop interface{}) (bool, error) {
	a := prop.(string)
	if len(a) < p.min {
		return false, errors.New("too short")
	} else if len(a) > p.max {
		return false, errors.New("too long")
	}
	return true, nil
}

type defaultValidator struct{}

func (d defaultValidator) validate(prop interface{}) (bool, error) {
	return true, nil
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

func validateStruct(str interface{}) /*[]error*/ {
	//errs := []error{}

	props := reflect.ValueOf(str)
	for i := 0; i < props.NumField(); i++ {
		validator := getValidator(props.Type().Field(i).Tag.Get("validate"))
		result, err := validator.validate(props.Field(i).Interface())

		if err != nil {
			errorMessage := props.Type().Field(i).Name + ": " + strconv.FormatBool(result) + " >> " + err.Error()
			fmt.Println(errorMessage)
		} else {
			message := props.Type().Field(i).Name + ": " + "Everything is Good!"
			fmt.Println(message)
		}

	}

}
