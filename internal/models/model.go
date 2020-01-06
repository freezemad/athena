package models

import (
  "time"
  "github.com/jmoiron/sqlx"
)

type Model interface {
  GetId() int
}

type modelImpl struct {
  id int
  Created  time.Time
  Updated  time.Time
}

func (m *modelImpl) SetUpdated(updated time.Time) {
  m.Updated = updated
}
