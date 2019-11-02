// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"reflect"
	"fmt"
	"regexp"
)

type Validate struct {}

func NewValidate() *Validate {
	return new(Validate)
}

func (validate *Validate) Do(ptr interface{}) error {

	t := reflect.TypeOf(ptr)

	if (t.Kind() != reflect.Ptr) || (t.Elem().Kind() != reflect.Struct) {
		panic("struct pointer is required")
	}

	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		reg := tag.Get("regexp")
		tips := tag.Get("tips")

		if reg == "" {
			continue
		}

		v := v.Field(i)

		strValue := ""

		switch v.Kind() {
		case reflect.String:
			strValue = v.String()
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = fmt.Sprintf("%d", v.Int())
		case reflect.Float32, reflect.Float64:
			strValue = fmt.Sprintf("%v", v.Float())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			strValue = string(v.Uint())
			break
		}

		if strValue != "" {
			matched, err := regexp.MatchString(reg, strValue)
			if err != nil {
				panic(err)
			}
			if !matched {
				if tips == "" {
					tips = fmt.Sprintf("%s is inlegal", field.Name)
				}
				return fmt.Errorf(tips)
			}
		}
	}

	return nil
}
