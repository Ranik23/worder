package corrector_service

import (
	"Worder/internal/config"
	"Worder/internal/storage"
	"Worder/internal/utils/backtracking"
	"log/slog"
)



type Corrector interface {
	CorrectWord(word string) (string, bool);
}

type CorrectorService struct {
	storage 	storage.Storage
	logger 		*slog.Logger
	config		*config.Config
}


func NewCorrectorService(strg storage.Storage, logger *slog.Logger, config *config.Config) *CorrectorService {
	return &CorrectorService{
		storage: strg,
		logger: logger,
		config:  config,
	}
}

func (cs *CorrectorService) CorrectWord(word string) (string , bool) {
	
	dictionary := cs.config.Dictionary.Words

	alphabet := cs.config.Alphabet.Syllables

	var answer string

	backtracking.BackTracking(word, dictionary, alphabet ,&answer)

	if answer == "" {
		return "", false
	}

	return answer, true
}




