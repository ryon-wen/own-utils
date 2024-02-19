package model

import (
	"time"

	"github.com/ryon-wen/own-utils/own"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterModel(sql *own.MySQL) {
	db = own.InitModel(sql)
	db.Table("orders")
}

type Orders struct {
	Id          int64
	UserId      int64
	GoodsList   string
	TotalAmount float64
	TradeNo     string
	TradeStatus int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func CreateOrder(order *Orders) error {
	return db.Omit("deleted_at").Create(&order).Error
}

func DeleteOrder(tradeNo string) error {
	return db.Where("trade_no = ?", tradeNo).Delete(&Orders{}).Error
}

func UpdateOrder(tradeNo string, tradeStatus int8) error {
	return db.Where("trade_no = ?", tradeNo).Update("trade_status", tradeStatus).Error
}
