package model

import (
	"fmt"
	"time"

	"github.com/ryon-wen/own-utils/own"
	"gorm.io/gorm"

	"2106A-zg6/srv/goods/pb"
)

var db *gorm.DB
var gdb *gorm.DB
var sdb *gorm.DB

func RegisterModel(sql *own.MySQL) {
	db = own.InitModel(sql)
	gdb = db
	gdb.Table("goods")
	sdb = db
	sdb.Table("saga_logs")
}

type Goods struct {
	Id        int64
	Name      string
	Price     float64
	Stock     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func AutoTx(txF func(tx *gorm.DB) error) error {
	tx := gdb.Begin()
	err := txF(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AddGoods(goods *Goods) error {
	return gdb.Omit("deleted_at").Create(&goods).Error
}

func UpdateGoods(goods *Goods) error {
	return gdb.Omit("created_at", "deleted_at").Updates(&goods).Error
}

func GetGoods(id int64) (*Goods, error) {
	var g Goods
	err := gdb.Where("id = ?", id).First(&g).Error
	return &g, err
}

func DeleteGoods(id int64) error {
	return gdb.Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func FindGoods(offset, limit int) ([]*Goods, error) {
	var goods []*Goods
	err := gdb.Offset(offset).Limit(limit).Find(&goods).Error
	return goods, err
}

func UpdateGoodsStock(gs []*pb.GoodsStock, gid string) error {
	err := AutoTx(func(tx *gorm.DB) error {
		for _, v := range gs {
			var g Goods
			err := tx.Where("id = ?", v.ID).First(&g).Set("gorm:query_option", "FOR UPDATE").Error
			if err != nil {
				db.Table("saga_logs").Omit("deleted_at").Create(&SagaLog{
					Gid:       gid,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				return err
			}
			if g.Stock < v.Stock {
				db.Table("saga_logs").Omit("deleted_at").Create(&SagaLog{
					Gid:       gid,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				return fmt.Errorf("%v 库存不足", g.Name)
			}
			g.Stock -= v.Stock
			res := tx.Where("id = ?", v.ID).Update("stock", g.Stock)
			if res.Error != nil || res.RowsAffected == 0 {
				db.Table("saga_logs").Omit("deleted_at").Create(&SagaLog{
					Gid:       gid,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				return res.Error
			}
		}
		return nil
	})
	return err
}

func UpdateGoodsStockBack(gs []*pb.GoodsStock) error {
	err := AutoTx(func(tx *gorm.DB) error {
		for _, v := range gs {
			res := tx.Where("id = ?", v.ID).Update("stock = stock + ?", v.Stock)
			if res.Error != nil || res.RowsAffected == 0 {
				return res.Error
			}
		}
		return nil
	})
	return err
}

func FindByIds(ids []int64) ([]*Goods, error) {
	var goods []*Goods
	err := gdb.Where("id IN (?)", ids).Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func GoodToPb(goods *Goods) *pb.Good {
	return &pb.Good{
		ID:    goods.Id,
		Name:  goods.Name,
		Price: goods.Price,
		Stock: goods.Stock,
	}
}
