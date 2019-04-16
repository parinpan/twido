package twitter

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	. "twido/dataprovider"
)

var InstanceTypes = map[string]reflect.Type{
	"search-tweets":   reflect.TypeOf(TwitterSearch{}),
	"show-status":     reflect.TypeOf(Status{}),
	"lookup-statuses": reflect.TypeOf([]Status{}),
}

type RequestConverter struct {
	ConvertName string
	Data        []byte
}

func NewRequestConverter(convertName string, data []byte) *RequestConverter {
	return &RequestConverter{
		ConvertName: convertName,
		Data:        data,
	}
}

func (rc *RequestConverter) Convert() (interface{}, error) {
	instanceType, exist := InstanceTypes[rc.ConvertName]

	if !exist {
		errorMsg := "Can't find object converter for " + rc.ConvertName + " key"
		log.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	objectPtr := reflect.New(instanceType)

	if err := json.Unmarshal(rc.Data, objectPtr.Interface()); nil != err {
		log.Println("Couldn't unmarshal data to object in converter")
		return nil, err
	}

	return objectPtr.Interface(), nil
}
