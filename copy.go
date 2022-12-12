package goutils

import (
	"reflect"
	"unsafe"
)

func Copy(origin, toCopy any) any {

	rs := reflect.ValueOf(origin)
	rsOrigin := reflect.New(rs.Type()).Elem()
	rsOrigin.Set(rs)

	rs = reflect.ValueOf(toCopy)
	rstoCopy := reflect.New(rs.Type()).Elem()
	rstoCopy.Set(rs)

	var fieldNameOrigin []string
	for i := 0; i < rsOrigin.NumField(); i++ {
		fieldNameOrigin = append(fieldNameOrigin, rsOrigin.Type().Field(i).Name)
	}

	var fieldNameToCopy []string
	for i := 0; i < rstoCopy.NumField(); i++ {
		fieldNameToCopy = append(fieldNameToCopy, rstoCopy.Type().Field(i).Name)
	}

	sameName := removeNotEqual(fieldNameOrigin, fieldNameToCopy)
	for i := 0; i < len(sameName); i++ {
		rf := reflect.NewAt(rstoCopy.FieldByName(sameName[i]).Type(), unsafe.Pointer(rstoCopy.FieldByName(sameName[i]).UnsafeAddr())).Elem()
		reflect.NewAt(rsOrigin.FieldByName(sameName[i]).Type(), unsafe.Pointer(rsOrigin.FieldByName(sameName[i]).UnsafeAddr())).Elem().Set(rf)
	}

	return rsOrigin.Interface()
}

func removeNotEqual(strSlice1, strSlice2 []string) []string {
	var item []string
	for _, value := range strSlice1 {
		if contains(strSlice2, value) {
			item = append(item, value)
		}
	}
	return item
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
