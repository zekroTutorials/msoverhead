package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

var mapInstance = func() map[string]interface{} {
	v := make(map[string]interface{})
	if err := json.Unmarshal(exapmePayloadB, &v); err != nil {
		panic(err)
	}
	return v
}()

func BenchmarkDeserializeStruct(b *testing.B) {
	reader := bytes.NewBuffer(exapmePayloadB)
	for i := 0; i < b.N; i++ {
		v := new(Book)
		json.NewDecoder(reader).Decode(v)
	}
}

func BenchmarkSerializeStruct(b *testing.B) {
	writer := bytes.NewBuffer([]byte{})
	for i := 0; i < b.N; i++ {
		json.NewEncoder(writer).Encode(examplePayloadInstance)
	}
}

func BenchmarkDeserializeMap(b *testing.B) {
	reader := bytes.NewBuffer(exapmePayloadB)
	for i := 0; i < b.N; i++ {
		v := make(map[string]interface{})
		json.NewDecoder(reader).Decode(&v)
	}
}

func BenchmarkSerializeMap(b *testing.B) {
	writer := bytes.NewBuffer([]byte{})
	for i := 0; i < b.N; i++ {
		json.NewEncoder(writer).Encode(mapInstance)
	}
}
