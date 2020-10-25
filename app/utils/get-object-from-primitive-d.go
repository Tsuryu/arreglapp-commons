package utils

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetObjectFromPrimitiveD : marshaller from mongo primitive d to an map interface
func GetObjectFromPrimitiveD(d interface{}) map[string]interface{} {
	primitiveMetadata := d.(primitive.D)
	var objectMap = make(map[string]interface{})
	for _, pair := range primitiveMetadata {
		if reflect.TypeOf(pair.Value).Name() == "D" {
			objectMap[pair.Key] = GetObjectFromPrimitiveD(pair.Value)
		} else {
			objectMap[pair.Key] = pair.Value
		}
	}
	return objectMap
}
