package models


type Request struct {
	ID 				int		`gorm:"primaryKey"`
	Word 			string  `gorm:"unqiue"`
	CorrectedWord	string
}