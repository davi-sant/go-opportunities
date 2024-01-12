package config

import "gorm.io/gorm"

var (
	_db_ *gorm.DB
	logger *Logger
)

func InitDataBase() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix);
	return logger
}