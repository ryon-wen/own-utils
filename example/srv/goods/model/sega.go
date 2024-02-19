package model

import (
	"time"
)

type SagaLog struct {
	Id        int64
	Gid       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func GidExist(gid string) bool {
	if db.Where("gid = ?", gid).First(&SagaLog{}).RowsAffected == 1 {
		return true
	}
	return false
}
