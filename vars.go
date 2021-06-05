package main

import "encoding/json"

var (
	nodeId, requests int
	target, delay    string
)

const examplePayload = `
{
	"id": 12345,
	"isbn_13": "978-1416572282",
	"isbn_10": "9781416572282",
	"title": "How to Teach Quantum Physics to Your Dog",
	"cover_image": "https://images-eu.ssl-images-amazon.com/images/I/41BS9KIfrFL.jpg",
	"n_sites": 304,
	"release_date": "2009-12-22T00:00:00Z00:00",
	"author": {
		"id": 67890,
		"name": "Chad Orzel",
		"country": "usa"
	}
}
`

var exapmePayloadB = []byte(examplePayload)

var examplePayloadInstance = func() *Book {
	v := new(Book)
	if err := json.Unmarshal(exapmePayloadB, v); err != nil {
		panic(err)
	}
	return v
}()
