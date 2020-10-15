package utils

import "reflect"

// MergeStruct merge values of two given structures into a simple one
func MergeStruct(old interface{}, new interface{}) interface{} {
	oldElement := reflect.ValueOf(old).Elem()
	newElement := reflect.ValueOf(new).Elem()

	for i := 0; i < newElement.NumField(); i++ {
		currentField := newElement.Field(i)

		if !currentField.IsZero() {
			oldElement.FieldByName(newElement.Type().Field(i).Name).Set(currentField)
		}
	}

	return oldElement
}
