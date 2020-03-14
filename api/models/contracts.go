package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Exclusivities struct {
	Countries []string `json:"countries"`
}

type Contract struct {
	ID             uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Fiscal_data_id uint32         `json:"fiscal_data_id" validate:"required" `
	Hotel_id       uint32         `json:"hotel_id" validate:"required" `
	Agent_id       uint32         `json:"agent_id" validate:"required"`
	Type           string         `json:"type" validate:"required" `
	Name           string         `json:"name" validate:"required" `
	PPU            string         `json:"ppu" validate:"required"`
	Subtype        string         `json:"subtype" validate:"required"`
	Exclusivities  postgres.Jsonb `gorm:"type:jsonb" json:"exclusivities"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Contract) ExclusivitiesUnmarshall() (*Exclusivities, error) {
	exec := &Exclusivities{}
	err := json.Unmarshal(c.Exclusivities.RawMessage, &Exclusivities{})
	if err != nil {
		return exec, err
	}
	return exec, err
}

func (c *Contract) Contracts(db *gorm.DB) (*[]Contract, error) {
	var err error
	contracts := []Contract{}
	err = db.Table("contracts").Find(&contracts).Error
	if err != nil {
		return &contracts, err
	}
	return &contracts, nil
}

func (c *Contract) SaveContract(db *gorm.DB) (*Contract, error) {
	var err error
	err = db.Table("contracts").Create(&c).Error
	if err != nil {
		return &Contract{}, err
	}
	return c, nil
}
