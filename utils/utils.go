package utils

import "go.mongodb.org/mongo-driver/bson"

func ToDocument(v interface{}) (bson.D, error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	var doc bson.D
	if err = bson.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return doc, nil
}
