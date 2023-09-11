package models

import "gorm.io/gorm"

type Thoughts struct {
  gorm.Model

  Say string `gorm:"not null"`
}
