package main

import (
	"math/rand"
	"time"
)

type Author struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Book struct {
	Id          int     `json:"id"`
	Isbn13      string  `json:"isbn_13"`
	Isbn10      string  `json:"isbn_10"`
	Title       string  `json:"title"`
	CoverImage  string  `json:"cover_image"`
	NSites      int     `json:"n_sites"`
	ReleaseDate string  `json:"release_date"`
	Author      *Author `json:"author"`
}

func RandomBook() *Book {
	return &Book{
		Id:          rand.Int(),
		Isbn13:      RandomString(14),
		Isbn10:      RandomString(13),
		Title:       RandomString(30),
		CoverImage:  RandomString(30),
		NSites:      rand.Int(),
		ReleaseDate: RandomString(20),
		Author: &Author{
			Id:      rand.Int(),
			Name:    RandomString(20),
			Country: RandomString(10),
		},
	}
}

type Stage struct {
	Node int    `json:"node"`
	Name string `json:"name"`
	T    int64  `json:"t"`
}

type TimeWrapper struct {
	Payload *Book    `json:"payload"`
	Stages  []*Stage `json:"stages"`
}

func NewTimeWrapper(payload *Book) (w *TimeWrapper) {
	w = &TimeWrapper{
		Payload: payload,
	}

	w.AddStageNow("init")

	return
}

func (w *TimeWrapper) AddStage(name string, t int64) {
	w.Stages = append(w.Stages, &Stage{nodeId, name, t})
}

func (w *TimeWrapper) AddStageNow(name string) {
	w.AddStage(name, time.Now().UnixNano())
}
