package models

type BookCsvLine struct {
	Title     string
	Author    string
	Genre     string
	Height    string
	Publisher string
}

type BookCsvLines []BookCsvLine
