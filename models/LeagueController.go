package models

import (
	"errors"
	"time"
)

type League struct {
	ID           uint
	DeletedAt    *time.Time `xml:"-" json:"-"`
	CreatedAt    *time.Time `xml:"-" json:"-"`
	UpdatedAt    *time.Time `xml:"-" json:"-"`
	Name         string
	LeagueConfig []LeagueConfig `json:",omitempty" xml:",omitempty"`
}

type LeagueController struct {
	db *ExportDB
}

func NewLeagueController(db *ExportDB) *LeagueController {
	return &LeagueController{db}
}

func (ct *LeagueController) GetAllLeagues() *[]League {
	leagues := new([]League)
	ct.db.Preload("LeagueConfig").Find(leagues)
	return leagues
}

// GetById will get a league by a specified league id
func (ct *LeagueController) GetById(id uint64) (*League, error) {
	league := League{}
	ct.db.First(&league, id)

	if league.ID == 0 {
		return nil, errors.New("No league has been found with that ID")
	}

	configs := []LeagueConfig{}
	ct.db.Model(&league).Related(&configs)
	league.LeagueConfig = configs
	return &league, nil
}

func (ct *LeagueController) Create(league *League) error {
	ct.db.Create(league)
	return nil
}