package db

import (
	"Worder/internal/config"
	"Worder/internal/models"
	"fmt"
	"log/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	ErrNotFound = gorm.ErrRecordNotFound
)

type PostgresStorage struct {
	DB *gorm.DB
	Logger *slog.Logger
	Config *config.Config
}

func NewPostgresStorage(logger *slog.Logger, config *config.Config) *PostgresStorage {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database")
		return nil
	}

	logger.Info("succesfully connected to database")

	db.AutoMigrate(models.Request{})

	return &PostgresStorage{
		DB: db,
		Logger: logger,
		Config: config,
	}
}

func (s *PostgresStorage) SaveWord(word, corrected string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		s.Logger.Error("failed to start transaction", "error", tx.Error)
		return tx.Error
	}
	
	result := s.DB.Create(&models.Request{
		Word:        word,
		CorrectedWord: corrected,
	})

	if result.Error != nil {
		s.Logger.Error("failed to insert request", "word", word, "corrected", corrected, "error", result.Error)
		return result.Error
	}
	tx.Commit()

	s.Logger.Info("successfully saved the word", "word", word, "corrected", corrected)
	return nil
}

func (s *PostgresStorage) DeleteWord(word string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		s.Logger.Error("failed to start transaction", "error", tx.Error)
		return tx.Error
	}

	var req models.Request

	result := s.DB.Where("word = ?", word).First(&req)
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			s.Logger.Error("word not found", "word", word)
			return ErrNotFound
		}

		s.Logger.Error("failed to find the request", "error", result.Error)
		return result.Error
	}

	result = s.DB.Delete(&req)

	if result.Error != nil {
		s.Logger.Error("failed to delete the request", "error", result.Error)
		return result.Error
	}

	tx.Commit()

	s.Logger.Info("successfully deleted the word", "word", word)

	return nil
}

func (s *PostgresStorage) GetStats(word string) (int, error)  {
	tx := s.DB.Begin()
	if tx.Error != nil {
		s.Logger.Error("failed to start transaction", "error", tx.Error)
		return -1, tx.Error
	}
	var count int64

	result := s.DB.Model(&models.Request{}).Where("word = ?", word).Count(&count)
	if result.Error != nil {
		s.Logger.Error("failed to get stats", "word", word, "error", result.Error)
		return 0, result.Error
	}

	return int(count), nil
}



