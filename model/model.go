package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 这块需要重构的，要用S3存储
type KeyStore struct {
	ID            uint64 `gorm:"primaryKey"`
	CreatetimeUtc int64  `gorm:"autoCreateTime"`
	Salt          string `gorm:"size:64"`
	Address       string `gorm:"size:80;uniqueIndex"`
	EncryptedKey  string `gorm:"size:255;comment:'v0 raw'"`
}

type Secret2FA struct {
	ID        uint64 `gorm:"primaryKey"`
	UUID      string `gorm:"size:64"`
	Address   string `gorm:"size:80;uniqueIndex"`
	Text      string `gorm:"size:32"`
	CreatedAt uint64 `gorm:"autoCreateTime"`
	CheckedAt uint64 `gorm:"autoUpdateTime"`
	DeletedAt uint64 `gorm:"default:null"`
}
