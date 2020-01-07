package models

import (
  "time"
  "github.com/jmoiron/sqlx"
)

type Model interface {
  GetId() int
}

type modelImpl struct {
  id       int        `db:"id"`
  Created  time.Time  `db:"Created"`
  Updated  time.Time  `db:"Updated"`
}

func (m *modelImpl) SetUpdated(updated time.Time) {
  m.Updated = updated
}
