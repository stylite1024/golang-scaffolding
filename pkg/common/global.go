package common

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// database
	Source string
	Driver string
	DBName string
	DB     *gorm.DB
	// redis
	RDS *redis.Client
	// validator
	Trans ut.Translator
)
