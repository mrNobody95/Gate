package models

import (
	"github.com/mrNobody95/Gate/brokerages"
	"gorm.io/gorm"
)

type Candle struct {
	gorm.Model
	Low        float64
	Vol        float64
	Time       int64
	Open       float64
	High       float64
	Close      float64
	Symbol     brokerages.Symbol
	Brokerage  brokerages.BrokerageName
	Resolution Resolution
	Indicators
}

func (c *Candle) LoadLast() error {
	return db.Model(&Candle{}).
		Where("brokerage LINE ?", c.Brokerage).
		Where("symbol LIKE ?", c.Symbol).
		Where("resolution_label LIKE ?", c.Resolution.Label).
		Last(&c).Error
}
