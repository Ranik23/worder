package storage



type Storage interface {
	SaveWord(word, corrected string) error
	DeleteWord(word string) error
	GetStats(word string) (int, error)
}
